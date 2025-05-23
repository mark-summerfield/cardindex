-- Copyright © 2025 Mark Summerfield. All Rights Reserved.

PRAGMA USER_VERSION = 1;

-- ==================== TABLES ====================

-- Name truncates at first newline or after . ! ? or at 50 chars.
-- Body is plain text using Commonmark markdown,
--      e.g., for **bold**, _italic_, lists,
--      urls [A Website](http://www.eg.com), and
--      images ![An Image](file:///home/mark/mags/image.png).
CREATE TABLE Cards (
    cid INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    Name TEXT AS (LTRIM(LTRIM(TRIM((SUBSTR(Body, 1,
                    MIN(50,
                        INSTR(Body || CHAR(10), CHAR(10)) - 1,
                        INSTR(Body || '.', '.'),
                        INSTR(Body || '!', '!'),
                        INSTR(Body || '?', '?')
                    )
                  ))), '#'))) VIRTUAL,
    Body TEXT NOT NULL,
    hidden BOOL DEFAULT FALSE NOT NULL,
    created REAL DEFAULT (JULIANDAY('NOW')) NOT NULL,
    updated REAL DEFAULT (JULIANDAY('NOW')) NOT NULL,

    CHECK(hidden IN (FALSE, TRUE))
);

-- Any box may contain any cards
-- To link two or more cards, create a box for them and add them to it
CREATE TABLE Boxes (
    bid INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    Name TEXT UNIQUE NOT NULL
);

CREATE TABLE CardsInBox (
    cid TEXT NOT NULL,
    bid TEXT NOT NULL,

    PRIMARY KEY (cid, bid),
    FOREIGN KEY(cid) REFERENCES Cards(cid),
    FOREIGN KEY(bid) REFERENCES Boxes(bid)
);

CREATE TABLE Searches (
    sid INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    SearchText TEXT NOT NULL,
    Hidden BOOL DEFAULT FALSE NOT NULL,
    Oid INTEGER DEFAULT 1, -- Default Name

    CHECK(Hidden IN (FALSE, TRUE)),
    CHECK(Oid IN (1, 2, 3)) -- 1→Name 2→Updated 3→Created
);

-- e.g., for MDI window sizes and positions
CREATE TABLE Config (
    Key TEXT PRIMARY KEY NOT NULL,
    Value TEXT
) WITHOUT ROWID;

-- ==================== VIEWS and VIRTUALS ====================

CREATE VIEW Counts AS
    SELECT Visible, Unboxed, Hidden
        FROM CountCardsVisible, CountCardsUnboxed, CountCardsHidden
        LIMIT 1;

CREATE VIEW CountCardsVisible AS
    SELECT COUNT(*) AS Visible FROM Cards WHERE hidden = FALSE LIMIT 1;

CREATE VIEW CountCardsHidden AS
    SELECT COUNT(*) AS Hidden FROM Cards WHERE hidden = TRUE LIMIT 1;

CREATE VIEW CountCardsUnboxed AS
    SELECT COUNT(*) AS Unboxed FROM Cards
        WHERE hidden = FALSE AND cid NOT IN (SELECT cid FROM CardsInBox)
        LIMIT 1;

-- we need created and updated for ORDER BY

CREATE VIEW ViewCardNamesVisible AS
    SELECT cid, Name, created, updated FROM Cards WHERE hidden = FALSE;

CREATE VIEW ViewCardNamesUnboxed AS
    SELECT cid, Name, created, updated FROM Cards
        WHERE hidden = FALSE AND cid NOT IN (SELECT cid FROM CardsInBox);

CREATE VIEW ViewCardNamesHidden AS
    SELECT cid, Name, created, updated FROM Cards WHERE hidden = TRUE;

CREATE VIEW ViewCardsVisible AS
    SELECT cid, Name, Body, DATETIME(created) AS created,
                            DATETIME(updated) AS updated
        FROM Cards WHERE hidden = FALSE;

CREATE VIEW ViewCardsHidden AS
    SELECT cid, Name, Body, DATETIME(created) AS created,
                            DATETIME(updated) AS updated
        FROM Cards WHERE hidden = TRUE;

CREATE VIEW ViewCardsUnboxed AS
    SELECT cid, Name, Body, DATETIME(created) AS created,
                            DATETIME(updated) AS updated
        FROM Cards
        WHERE hidden = FALSE AND cid NOT IN (SELECT cid FROM CardsInBox);

CREATE VIRTUAL TABLE vt_fts_cards
    USING FTS5(Body, tokenize='porter unicode61 remove_diacritics 2');

-- ==================== TRIGGERS ====================

CREATE TRIGGER insert_card_trigger AFTER INSERT ON Cards
    FOR EACH ROW -- update FTS
BEGIN
    INSERT OR REPLACE INTO vt_fts_cards (rowid, Body) VALUES
        (NEW.cid, NEW.Body);
END;

CREATE TRIGGER update_cards_timestamp_trigger AFTER UPDATE ON Cards
    FOR EACH ROW
BEGIN
    UPDATE Cards SET updated = (JULIANDAY('NOW')) WHERE cid = OLD.cid;
END;

CREATE TRIGGER update_card_body_trigger AFTER UPDATE OF Body ON Cards
    FOR EACH ROW -- update FTS
BEGIN
    INSERT OR REPLACE INTO vt_fts_cards (rowid, Body) VALUES
        (NEW.cid, NEW.Body);
END;

CREATE TRIGGER delete_card_trigger_before BEFORE DELETE ON Cards
    FOR EACH ROW
        WHEN EXISTS (SELECT TRUE FROM Cards WHERE Cards.cid = OLD.cid AND
                                                  OLD.hidden = FALSE)
BEGIN
    SELECT RAISE(ABORT, 'can only delete hidden cards');
END;

CREATE TRIGGER delete_card_trigger_after AFTER DELETE ON Cards
    FOR EACH ROW
BEGIN
    DELETE FROM vt_fts_cards WHERE rowid = OLD.cid; -- remove from FTS
    DELETE FROM CardsInBox WHERE cid = OLD.cid; -- remove from any boxes
END;

CREATE TRIGGER delete_box BEFORE DELETE ON Boxes
    FOR EACH ROW
        WHEN EXISTS (SELECT TRUE FROM CardsInBox
                     WHERE CardsInBox.bid = OLD.bid)
BEGIN
    SELECT RAISE(ABORT, 'can only delete unused boxes');
END;

-- ==================== INSERTIONS ====================

INSERT INTO Config (Key, Value) VALUES ('Created', JULIANDAY('NOW'));
INSERT INTO Config (Key, Value) VALUES ('Updated', JULIANDAY('NOW'));
INSERT INTO Config (Key, Value) VALUES ('N', 1); -- for optimizing
