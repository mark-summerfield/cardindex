-- Copyright © 2025 Mark Summerfield. All Rights Reserved.

PRAGMA USER_VERSION = 1;

-- ==================== TABLES ====================

-- Commonmark markdown Body, e.g., for **bold**, _italic_, lists,
-- urls [Website](http://www.eg.com), dates (e.g., YYYY-MM-DD) and
-- images ![Cover Image](file:///home/mark/mags/image.png).
CREATE TABLE Cards (
    cid INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    Body TEXT NOT NULL, -- Simple Markdown; first "line" is Card's Name
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

CREATE TABLE Queries ( -- See default queries INSERTed below
    qid INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    Name TEXT DEFAULT '' NOT NULL, -- Use "Query #qid" if empty
    MatchText TEXT,
    LinkedToCid INTEGER, -- if not NULL matches all cards linked to this one
    Unboxed BOOL, -- if TRUE match: Card.bid NOT IN CardsInBox.bid
    InBoxes TEXT, -- Space-separated list of bids
    NotInBoxes TEXT,
    UpdatedAfter TEXT, -- For all these NULL means don't care
    UpdatedBefore TEXT,
    CreatedAfter TEXT,
    CreatedBefore TEXT,
    Hidden BOOL DEFAULT FALSE, -- by default not Hidden
    OrderBy TEXT DEFAULT 'updated DESC', -- default most to least recent

    FOREIGN KEY(LinkedToCid) REFERENCES Cards(cid),
    CHECK(Hidden IS NULL OR Hidden IN (0, 1)),
    CHECK(Unboxed IS NULL OR Unboxed IN (0, 1))
);

-- e.g., for MDI window sizes and positions
CREATE TABLE Config (
    Key TEXT PRIMARY KEY NOT NULL,
    Value TEXT
) WITHOUT ROWID;

-- ==================== VIEWS and VIRTUALS ====================

CREATE VIEW CardNames AS SELECT cid, TRIM(LTRIM(Name, '#')) AS Name
    FROM _CardNames;

-- Truncates at first newline or after . ! ? or at 50 chars.
CREATE VIEW _CardNames AS
    SELECT cid, TRIM((SUBSTR(Body, 1,
                        MIN(50,
                            INSTR(Body || CHAR(10), CHAR(10)) - 1,
                            INSTR(Body || '.', '.'),
                            INSTR(Body || '!', '!'),
                            INSTR(Body || '?', '?')
                        ))))
        AS Name FROM Cards ORDER BY LOWER(Name);

CREATE VIEW CardsView AS
    SELECT cid, Body, hidden, DATETIME(created) AS created,
                              DATETIME(updated) AS updated
    FROM Cards ORDER BY updated DESC;

CREATE VIRTUAL TABLE v_fts_cards USING FTS5(Body, tokenize=porter);

-- ==================== TRIGGERS ====================

CREATE TRIGGER insert_queries_trigger AFTER INSERT ON Queries
    FOR EACH ROW
        WHEN EXISTS (SELECT 1 FROM Queries WHERE Queries.Name = '' AND
                                                 Queries.qid = NEW.qid)
BEGIN
    -- UPDATE Queries SET Name = FORMAT('Query #%d', NEW.qid)
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
    INSERT OR REPLACE INTO v_fts_cards (rowid, Body) VALUES
        (NEW.cid, NEW.Body);
END;

CREATE TRIGGER update_card_body_trigger AFTER UPDATE OF Body ON Cards
    FOR EACH ROW -- update FTS
BEGIN
    INSERT OR REPLACE INTO v_fts_cards (rowid, Body) VALUES
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
    DELETE FROM v_fts_cards WHERE rowid = OLD.cid; -- remove from FTS
    DELETE FROM CardsInBox WHERE cid = OLD.cid; -- leave any boxes
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
    (0, 'All Cards'); -- Excludes hidden (hidden is FALSE by default)
INSERT INTO Queries (qid, Name, Unboxed) VALUES
    (1, 'Unboxed Cards', TRUE); -- Excludes hidden
INSERT INTO Queries (qid, Name, Hidden) VALUES
    (2, 'Hidden Cards', TRUE);
