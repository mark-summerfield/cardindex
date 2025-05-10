-- Copyright © 2025 Mark Summerfield. All Rights Reserved.

PRAGMA user_version = 1;

-- HTML body (e.g., for bold, italic, color) and for links
-- (e.g., http://... card://123) and for dates (e.g., YYYY-MM-DD).
CREATE TABLE Cards (
    cid INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    Body TEXT NOT NULL, -- HTML; first "line" is Card's Name
    Image BLOB, -- SVG or PNG etc.
    created TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE VIRTUAL TABLE v_card_words USING FTS5(Body, tokenize=porter);

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

CREATE TABLE Queries (
    qid INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    Query TEXT NOT NULL -- NOTE JSON-format details of saved query
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
    INSERT OR REPLACE INTO v_card_words (rowid, Body) VALUES
        (NEW.cid, text_for_html(NEW.Body));
END;

CREATE TRIGGER update_card_body_trigger AFTER UPDATE OF Body ON Cards
    FOR EACH ROW -- update FTS
BEGIN
    INSERT OR REPLACE INTO v_card_words (rowid, Body) VALUES
        (NEW.cid, text_for_html(NEW.Body));
END;

CREATE TRIGGER delete_card_trigger_before BEFORE DELETE ON Cards
    FOR EACH ROW
        WHEN EXISTS (SELECT 1 FROM Card_x_Group
                     WHERE Card_x_Group.cid = OLD.cid AND
                           Card_x_Group.gid = 0) -- 0 is Hidden group's gid
BEGIN
    SELECT RAISE(ABORT, 'can only delete hidden cards');
END;

CREATE TRIGGER delete_card_trigger_after AFTER DELETE ON Cards
    FOR EACH ROW
BEGIN
    DELETE FROM v_card_words WHERE rowid = OLD.cid; -- remove from FTS
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

INSERT INTO Config (Key, Value) VALUES ('Created', DATETIME('NOW'));
INSERT INTO Config (Key, Value) VALUES ('Updated', DATETIME('NOW'));
INSERT INTO Config (Key, Value) VALUES ('N', 1); -- for optimizing

INSERT INTO Groups (gid, Name) VALUES (0, '«Hidden»');

INSERT INTO Queries (qid, Query) VALUES
    (0, '{"name": "All Cards (excl. Hidden)", "predefined": "all"}');
INSERT INTO Queries (qid, Query) VALUES
    (1, '{"name": "Ungrouped Cards", "predefined": "ungrouped"}');
INSERT INTO Queries (qid, Query) VALUES
    (2, '{"name": "Hidden Cards", "predefined": "ungrouped"}');
