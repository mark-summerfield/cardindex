// Copyright © 2025 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	_ "embed"
)

//go:embed Version.dat
var Version string

//go:embed sql/prepare.sql
var SQL_PREPARE string

//go:embed sql/create.sql
var SQL_CREATE string

const (
	APPNAME = "CardIndex"

	DRIVER    = "sqlite"
	MAX_OPENS = 11

	CREATED = "Created"
	UPDATED = "Updated"

	SQL_BEGIN    = "BEGIN;"
	SQL_COMMIT   = "COMMIT;"
	SQL_OPTIMIZE = `INSERT INTO vt_fts_cards(vt_fts_cards)
						VALUES ('optimize');
					VACUUM;`
	SQL_VERSION               = "SELECT SQLITE_VERSION();"
	SQL_UPDATE_CONFIG_UPDATED = `UPDATE Config SET Value = JULIANDAY('NOW')
								 WHERE Key = 'Updated';`
	SQL_GET_CONFIG_N        = "SELECT Value FROM Config WHERE Key = N;"
	SQL_GET_WHEN            = "SELECT DATETIME(Value) FROM Config WHERE Key = ?;"
	SQL_ZERO_CONFIG_N       = "UPDATE Config SET Value = 0 WHERE Key = N;"
	SQL_GET_COUNTS          = "SELECT Visible, Unboxed, Hidden FROM Counts"
	SQL_INSERT_CARD         = "INSERT INTO Cards (Body) VALUES (?);"
	SQL_UPDATE_CARD         = "UPDATE Cards SET Body = ? WHERE cid = ?;"
	SQL_CARD_BODY           = "SELECT Body FROM Cards WHERE cid = ?;"
	SQL_CARD_HIDDEN         = "SELECT hidden FROM Cards WHERE cid = ?;"
	SQL_SET_CARD_VISIBILITY = "UPDATE Cards SET hidden = ? WHERE cid = ?;"
	SQL_DELETE_CARD         = "DELETE FROM Cards WHERE cid = ?;"
)

// var E100 = errors.New("E100: failed to get counts")
