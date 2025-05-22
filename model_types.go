// Copyright © 2025 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"fmt"
	"time"
)

type CardCounts struct{ Visible, Unboxed, Hidden int }

type CardName struct {
	cid  int
	name string
}

func (me CardName) String() string { // for debugging
	return fmt.Sprintf("CardName#%d %q", me.cid, me.name)
}

type Card struct {
	cid     int
	body    string
	hidden  bool
	created time.Time
	updated time.Time
}

func (me Card) String() string { // for debugging
	cid := me.cid
	if me.hidden {
		cid = -cid
	}
	created := me.created.Format(time.DateTime)
	updated := me.updated.Format(time.DateTime)
	return fmt.Sprintf("Card#%d %s • %s\n\t%q", cid, created, updated,
		me.body)
}

type Box struct {
	bid  int
	name string
}

func (me Box) String() string { // for debugging
	return fmt.Sprintf("Box#%d %q", me.bid, me.name)
}

type Query struct {
	qid        int
	name       string
	matchText  string
	inBoxes    []int
	notInBoxes []int
	hidden     bool
	oid        Oid
}

func NewQuery(name, matchText string, inBoxes, notInBoxes []int,
	hidden bool, oid Oid,
) Query {
	return Query{
		INVALID_ID, name, matchText, inBoxes, notInBoxes, hidden, oid,
	}
}

func (me Query) String() string { // for debugging
	ins := strForBids(me.inBoxes)
	notins := strForBids(me.notInBoxes)
	return fmt.Sprintf(
		"Query#%d %q match=%q in=[%s] not-in=[%s] hidden=%t oid=%d(%s)",
		me.qid, me.name, me.matchText, ins, notins, me.hidden, me.oid,
		me.oid)
}

func (me Query) Sql() (string, []any) {
	args := []any{}
	sql := "SELECT cid, Name FROM Cards WHERE"
	if me.hidden {
		sql += " hidden = TRUE"
	} else {
		sql += " hidden = FALSE"
	}
	if me.matchText != "" {
		sql += " AND cid IN (SELECT ROWID FROM vt_fts_cards(?))"
		args = append(args, me.matchText)
	}
	if len(me.inBoxes) > 0 {
		boxes := strForBids(me.inBoxes)
		sql += " AND cid IN (SELECT cid FROM CardsInBox WHERE bid IN (" +
			boxes + "))"
	}
	if len(me.notInBoxes) > 0 {
		boxes := strForBids(me.notInBoxes)
		sql += " AND cid NOT IN (SELECT cid FROM CardsInBox WHERE bid " +
			"IN (" + boxes + "))"
	}
	if me.oid != OID_IGNORE {
		sql += " " + me.oid.Sql() + ";"
	}
	if len(args) > 0 {
		return sql, args
	}
	return sql, nil
}
