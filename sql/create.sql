-- Copyright © 2025 Mark Summerfield. All Rights Reserved.

PRAGMA USER_VERSION = 1;

-- ==================== TABLES ====================

-- Commonmark markdown Body, e.g., for **bold**, _italic_, lists,
-- urls [Website](http://www.eg.com), dates (e.g., YYYY-MM-DD) and
-- images ![Cover Image](file:///home/mark/mags/image.png).
CREATE TABLE Cards (
    cid INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    Body TEXT NOT NULL, -- First "line" is Card's Name (see CardNames view)
    hidden BOOL DEFAULT FALSE NOT NULL,
    created REAL DEFAULT (JULIANDAY('NOW')) NOT NULL,
    updated REAL DEFAULT (JULIANDAY('NOW')) NOT NULL,

    CHECK(hidden IN (0, 1))
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

-- if Name = '' use: COALESCE(MatchText, 'Query #' || qid)
CREATE TABLE Queries ( -- See default queries INSERTed below
    qid INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    Name TEXT DEFAULT '' NOT NULL,
    MatchText TEXT,
    Unboxed BOOL, -- If TRUE match cards that are not in CardsInBox
    InBoxes TEXT, -- Space-separated list of bids
    NotInBoxes TEXT,
    Hidden BOOL DEFAULT FALSE, -- By default not Hidden
    OrderBy TEXT DEFAULT 'updated DESC', -- Default most to least recent

    CHECK(Hidden IS NULL OR Hidden IN (0, 1)),
    CHECK(Unboxed IS NULL OR Unboxed IN (0, 1))
);

-- e.g., for MDI window sizes and positions
CREATE TABLE Config (
    Key TEXT PRIMARY KEY NOT NULL,
    Value TEXT
) WITHOUT ROWID;

-- ==================== VIEWS and VIRTUALS ====================

CREATE VIEW ViewCardsUnboxed AS
    SELECT cid, Name, Body, created, updated
        FROM viewCards
        WHERE hidden = FALSE AND cid NOT IN (SELECT cid FROM CardsInBox);

CREATE VIEW ViewCardsVisible AS
    SELECT cid, Name, Body, created, updated FROM viewCards
        WHERE hidden = FALSE;

CREATE VIEW ViewCardsHidden AS
    SELECT cid, Name, Body, created, updated FROM viewCards
        WHERE hidden = TRUE;

-- Truncates at first newline or after . ! ? or at 50 chars.
CREATE VIEW viewCards AS
    SELECT cid, LTRIM(LTRIM(TRIM(SUBSTR(Body, 1,
                      MIN(
                          50,
                          INSTR(Body || CHAR(10), CHAR(10)) - 1,
                          INSTR(Body || '.', '.'),
                          INSTR(Body || '!', '!'),
                          INSTR(Body || '?', '?')
                      ))), '#'))
        AS Name, Body, hidden, DATETIME(created) AS created,
                               DATETIME(updated) AS updated
        FROM Cards;

CREATE VIRTUAL TABLE vt_fts_cards USING FTS5(Body, tokenize=porter);

-- ==================== TRIGGERS ====================

CREATE TRIGGER insert_queries_trigger AFTER INSERT ON Queries
    FOR EACH ROW
        WHEN EXISTS (SELECT 1 FROM Queries WHERE Queries.Name = '' AND
                                                 Queries.qid = NEW.qid)
BEGIN
    -- UPDATE Queries SET Name = FORMAT('Query #%d', NEW.qid) -- new syntax
    UPDATE Queries SET Name = PRINTF('Query #%d', NEW.qid) -- old syntax
        WHERE qid = NEW.qid;
END;

CREATE TRIGGER update_cards_timestamp_trigger AFTER UPDATE ON Cards
    FOR EACH ROW
BEGIN
    UPDATE Cards SET updated = (JULIANDAY('NOW')) WHERE cid = OLD.cid;
END;

CREATE TRIGGER insert_card_trigger AFTER INSERT ON Cards
    FOR EACH ROW -- update FTS
BEGIN
    INSERT OR REPLACE INTO vt_fts_cards (rowid, Body) VALUES
        (NEW.cid, NEW.Body);
END;

CREATE TRIGGER update_card_body_trigger AFTER UPDATE OF Body ON Cards
    FOR EACH ROW -- update FTS
BEGIN
    INSERT OR REPLACE INTO vt_fts_cards (rowid, Body) VALUES
        (NEW.cid, NEW.Body);
END;

CREATE TRIGGER delete_card_trigger_before BEFORE DELETE ON Cards
    FOR EACH ROW
        WHEN EXISTS (SELECT 1 FROM Cards WHERE Cards.cid = OLD.cid AND
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
        WHEN EXISTS (SELECT 1 FROM Cards WHERE Cards.bid = OLD.bid)
BEGIN
    SELECT RAISE(ABORT, 'can only delete unused boxes');
END;

CREATE TRIGGER delete_query BEFORE DELETE ON Queries
    FOR EACH ROW
        WHEN Old.qid IN (0, 1, 2)
BEGIN
    SELECT RAISE(ABORT, 'can only delete user created queries');
END;

-- ==================== INSERTIONS ====================

INSERT INTO Config (Key, Value) VALUES ('Created', JULIANDAY('NOW'));
INSERT INTO Config (Key, Value) VALUES ('Updated', JULIANDAY('NOW'));
INSERT INTO Config (Key, Value) VALUES ('N', 1); -- for optimizing

INSERT INTO Queries (qid, Name) VALUES
    (0, 'All Cards'); -- ViewCardsVisible
INSERT INTO Queries (qid, Name, Unboxed) VALUES
    (1, 'Unboxed Cards', TRUE); -- ViewCardsUnboxed
INSERT INTO Queries (qid, Name, Hidden) VALUES
    (2, 'Hidden Cards', TRUE); -- ViewCardsHidden
