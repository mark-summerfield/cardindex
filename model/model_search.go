// Copyright © 2025 Mark Summerfield. All rights reserved.
// License: GPL-3

package model

import "fmt"

type Search struct {
	sid        int
	searchText string
	hidden     bool
	oid        Oid
}

func NewSearch(searchText string, hidden bool, oid Oid) Search {
	return Search{INVALID_ID, searchText, hidden, oid}
}

func (me Search) String() string { // for debugging
	return fmt.Sprintf("Search#%d searchText=%q hidden=%t oid=%d(%s)",
		me.sid, me.searchText, me.hidden, me.oid, me.oid)
}

func (me Search) Query(orderby bool) (string, []any) {
	args := []any{}
	query := "SELECT cid, Name FROM Cards WHERE hidden = "
	if me.hidden {
		query += "TRUE"
	} else {
		query += "FALSE"
	}
	if me.searchText != "" {
		query += " AND cid IN (SELECT ROWID FROM vt_fts_cards(?))"
		args = append(args, me.searchText)
	}
	if orderby && me.oid != OID_IGNORE {
		query += " " + me.oid.Query()
	}
	query += ";"
	if len(args) > 0 {
		return query, args
	}
	return query, nil
}
