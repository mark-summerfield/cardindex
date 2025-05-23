// Copyright © 2025 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

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

func (me Search) Sql(orderby bool) (string, []any) {
	args := []any{}
	sql := "SELECT cid, Name FROM Cards WHERE hidden = "
	if me.hidden {
		sql += "TRUE"
	} else {
		sql += "FALSE"
	}
	if me.searchText != "" {
		sql += " AND cid IN (SELECT ROWID FROM vt_fts_cards(?))"
		args = append(args, me.searchText)
	}
	if orderby && me.oid != OID_IGNORE {
		sql += " " + me.oid.Sql()
	}
	sql += ";"
	if len(args) > 0 {
		return sql, args
	}
	return sql, nil
}
