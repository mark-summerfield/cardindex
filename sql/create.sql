-- Copyright © 2024-25 Mark Summerfield. All Rights Reserved.

PRAGMA user_version = 1;

-- body should be in markup (e.g., for bold, italic, color) and for links
-- (e.g., http://... card://123) and for dates (e.g., YYYY-MM-DD).
CREATE TABLE Cards (
    cid INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    Name TEXT NOT NULL,
    Body TEXT,
    created TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

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
    Query TEXT NOT NULL -- TODO details of saved query; may use more fields
);

-- e.g., for MDI window sizes and positions
CREATE TABLE Config (
    Key TEXT PRIMARY KEY NOT NULL,
    Value TEXT
) WITHOUT ROWID;

CREATE TRIGGER update_cards_timestamp_trigger
    AFTER UPDATE ON Cards FOR EACH ROW
BEGIN
    UPDATE Cards SET updated = CURRENT_TIMESTAMP WHERE cid = old.cid;
END;

-- Disallow deleting the Hidden group
CREATE TRIGGER delete_group_gid
    BEFORE DELETE ON Groups FOR EACH ROW WHEN OLD.gid = 1
BEGIN
    SELECT RAISE(ABORT, 'Cannot delete the Hidden Group');
END;

-- NOTE always check that a group is not in use before deleting it
CREATE TRIGGER delete_group
    BEFORE DELETE ON Groups FOR EACH ROW
        WHEN EXISTS (SELECT 1 FROM Cards WHERE Cards.gid = OLD.gid)
BEGIN
    SELECT RAISE(ABORT,
                 'Cannot delete a Group that at least one card uses');
END;

INSERT INTO Groups (gid, Name) VALUES (0, '«Hidden»');
