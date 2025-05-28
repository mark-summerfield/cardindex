// Copyright © 2025 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import "embed"

//go:embed Version.dat
var Version string

//go:embed images/*.svg
var ICONS embed.FS

//go:embed sql/prepare.sql
var SQL_PREPARE string

//go:embed sql/create.sql
var SQL_CREATE string

const (
	APPNAME = "CardIndex"
	DOMAIN  = "MNS"

	CONFIG_WINDOW           = "Window"
	CONFIG_WINDOW_STATE     = "State"
	CONFIG_WINDOW_GEOMETRY  = "Geometry"
	CONFIG_CURSOR_BLINK     = "CursorBlink"
	DEFAULT_CURSOR_BLINK    = true
	CONFIG_MOST_RECENT_FILE = "MostRecentFile"

	SVG_ICON                 = "cardindex.svg"
	SVG_FILE_NEW             = "document-new.svg"
	SVG_FILE_OPEN            = "document-open.svg"
	SVG_FILE_SAVE            = "document-save.svg"
	SVG_FILE_SAVE_AS         = "document-save-as.svg"
	SVG_FILE_EXPORT          = "export.svg"
	SVG_FILE_CONFIGURE       = "document-properties.svg"
	SVG_FILE_QUIT            = "shutdown.svg"
	SVG_EDIT_COPY            = "edit-copy.svg"
	SVG_EDIT_CUT             = "edit-cut.svg"
	SVG_EDIT_PASTE           = "edit-paste.svg"
	SVG_CARD_NEW             = "card-new.svg"
	SVG_CARD_ADD_TO_BOX      = "card-add-to-box.svg"
	SVG_CARD_REMOVE_FROM_BOX = "card-remove-from-box.svg"
	SVG_CARD_EXPORT          = "card-export.svg"
	SVG_CARD_UNHIDE          = "card-unhide.svg"
	SVG_CARD_HIDE            = "card-hide.svg"
	SVG_CARD_DELETE          = "card-delete.svg"

	TIMEOUT_LONG = 10000

	DRIVER    = "sqlite"
	MAX_OPENS = 11

	INVALID_ID = -1

	NAME    = "Name"
	UPDATED = "Updated"
	CREATED = "Created"

	OID_IGNORE  Oid = 0
	OID_NAME    Oid = 1
	OID_UPDATED Oid = 2
	OID_CREATED Oid = 3

	SQL_BEGIN    = "BEGIN;"
	SQL_COMMIT   = "COMMIT;"
	SQL_OPTIMIZE = `INSERT INTO vt_fts_cards(vt_fts_cards)
						VALUES ('optimize');
					VACUUM;`
	SQL_VERSION       = "SELECT SQLITE_VERSION();"
	SQL_CONFIG_UPDATE = `UPDATE Config SET Value = JULIANDAY('NOW')
							WHERE Key = 'Updated';`
	SQL_CONFIG_GET_N    = "SELECT Value FROM Config WHERE Key = N;"
	SQL_CONFIG_GET_WHEN = `SELECT DATETIME(Value) FROM Config WHERE Key = ?;`
	SQL_CONFIG_ZERO_N   = "UPDATE Config SET Value = 0 WHERE Key = N;"
	SQL_CARD_COUNTS     = "SELECT Visible, Unboxed, Hidden FROM Counts;"
	SQL_CARD_GET        = `SELECT Body, hidden, DATETIME(created),
								DATETIME(updated) FROM Cards WHERE cid = ?;`
	SQL_CARD_INSERT        = "INSERT INTO Cards (Body) VALUES (?);"
	SQL_CARD_UPDATE        = "UPDATE Cards SET Body = ? WHERE cid = ?;"
	SQL_CARD_GET_HIDDEN    = "SELECT hidden FROM Cards WHERE cid = ?;"
	SQL_CARD_VISIBILITY    = "UPDATE Cards SET hidden = ? WHERE cid = ?;"
	SQL_CARD_DELETE        = "DELETE FROM Cards WHERE cid = ?;"
	SQL_CARD_NAMES_VISIBLE = "SELECT cid, Name FROM ViewCardNamesVisible;"
	SQL_CARD_NAMES_UNBOXED = "SELECT cid, Name FROM ViewCardNamesUnboxed;"
	SQL_CARD_NAMES_HIDDEN  = "SELECT cid, Name FROM ViewCardNamesHidden;"
	SQL_BOX_GET            = "SELECT Name FROM Boxes WHERE bid = ?;"
	SQL_BOX_INSERT         = "INSERT INTO Boxes (Name) VALUES (?);"
	SQL_BOX_UPDATE         = "UPDATE Boxes Set Name = ? WHERE bid = ?;"
	SQL_BOX_DELETE         = "DELETE FROM Boxes WHERE bid = ?;"
	SQL_BOXES              = `SELECT bid, Name FROM Boxes
							  ORDER BY LOWER(Name);`
	SQL_BOX_IN_USE = `SELECT EXISTS (SELECT * FROM CardsInBox 
									 WHERE bid = ?) AS Found;`
	SQL_BOX_ADD_CARD    = "INSERT INTO CardsInBox (cid, bid) VALUES (?, ?);"
	SQL_BOX_REMOVE_CARD = `DELETE FROM CardsInBox
								WHERE cid = ? AND bid = ?;`
	SQL_SEARCH_GET = `SELECT SearchText, Hidden, Oid FROM Searches
							WHERE sid = ?;`
	SQL_SEARCH_INSERT = `INSERT INTO Searches (SearchText, Hidden, Oid)
							VALUES (?, ?, ?);`
	SQL_SEARCH_UPDATE = `UPDATE Searches SearchText = ?, Hidden = ?,
							Oid = ?;`
	SQL_SEARCH_DELETE = "DELETE FROM Searches WHERE sid = ?;"
	SQL_SEARCHES      = `SELECT sid, SearchText, Hidden, Oid FROM Searches
							ORDER BY LOWER(SearchText)`
)

// var E100 = errors.New("E100: failed to get counts")
