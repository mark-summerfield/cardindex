-- Copyright © 2025 Mark Summerfield. All Rights Reserved.

PRAGMA user_version = 1;

-- Markdown Body (e.g., for bold, italic, and lists), and for links
-- [Apple II](card://123) and for dates (e.g., YYYY-MM-DD).
CREATE TABLE Cards (
    cid INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    Body TEXT NOT NULL, -- Simple Markdown; first "line" is Card's Name
    Image BLOB, -- SVG or PNG etc.
    hidden BOOL DEFAULT FALSE NOT NULL,
    created TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,

    CHECK(hidden IN (0, 1))
);

CREATE VIRTUAL TABLE v_fts_cards USING FTS5(Body, tokenize=porter);

CREATE TABLE Groups (
    gid INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    Name TEXT UNIQUE NOT NULL
);

CREATE TABLE Card_x_Group (
    cid TEXT NOT NULL,
    gid TEXT NOT NULL,

    PRIMARY KEY (cid, gid),
    FOREIGN KEY(cid) REFERENCES Cards(cid),
    FOREIGN KEY(gid) REFERENCES Groups(gid)
);

CREATE TABLE Queries ( -- See default queries INSERTed below
    qid INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    Name TEXT NOT NULL,
    MatchText TEXT,
    HasImage BOOL,
    Ungrouped BOOL, -- if TRUE match: Card.gid NOT IN Card_x_Group.gid
    InGroups TEXT, -- Space-separated list of gids
    NotInGroups TEXT,
    UpdatedAfter TEXT,
    UpdatedBefore TEXT,
    CreatedAfter TEXT, -- For all these NULL means don't care
    CreatedBefore TEXT,
    Hidden BOOL DEFAULT FALSE, -- by default not Hidden
    OrderBy TEXT DEFAULT 'updated DESC' -- default most to least recent

    CHECK(HasImage IS NULL OR HasImage IN (0, 1)),
    CHECK(Hidden IS NULL OR Hidden IN (0, 1)),
    CHECK(Ungrouped IS NULL OR Ungrouped IN (0, 1))
);

-- e.g., for MDI window sizes and positions
CREATE TABLE Config (
    Key TEXT PRIMARY KEY NOT NULL,
    Value TEXT
) WITHOUT ROWID;

CREATE TRIGGER update_cards_timestamp_trigger AFTER UPDATE ON Cards
    FOR EACH ROW
BEGIN
    UPDATE Cards SET updated = CURRENT_TIMESTAMP WHERE cid = OLD.cid;
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
    DELETE FROM Card_x_Group WHERE cid = OLD.cid; -- leave any groups
END;

CREATE TRIGGER delete_group BEFORE DELETE ON Groups
    FOR EACH ROW
        WHEN EXISTS (SELECT 1 FROM Cards WHERE Cards.gid = OLD.gid)
BEGIN
    SELECT RAISE(ABORT, 'can only delete unused groups');
END;

CREATE TRIGGER delete_query BEFORE DELETE ON Queries
    FOR EACH ROW
        WHEN Old.qid IN (0, 1, 2)
BEGIN
    SELECT RAISE(ABORT, 'can only delete user created queries');
END;

INSERT INTO Config (Key, Value) VALUES ('Created', CURRENT_TIMESTAMP);
INSERT INTO Config (Key, Value) VALUES ('Updated', CURRENT_TIMESTAMP);
INSERT INTO Config (Key, Value) VALUES ('N', 1); -- for optimizing

INSERT INTO Queries (qid, Name) VALUES
    (0, 'All Cards'); -- Excludes hidden (hidden is FALSE by default)
INSERT INTO Queries (qid, Name, Ungrouped) VALUES
    (1, 'Ungrouped Cards', TRUE); -- Excludes hidden
INSERT INTO Queries (qid, Name, Hidden) VALUES
    (2, 'Hidden Cards', TRUE);
