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
	by         string
}

func NewQuery(name, matchText string, inBoxes, notInBoxes []int,
	hidden bool, by string,
) Query {
	return Query{
		INVALID_ID, name, matchText, inBoxes, notInBoxes,
		hidden, by,
	}
}

func (me Query) String() string { // for debugging
	ins := strForBids(me.inBoxes)
	notins := strForBids(me.notInBoxes)
	return fmt.Sprintf(
		"Query#%d %q match=%q in=[%s] not-in=[%s] hidden=%t by=%s", me.qid,
		me.name, me.matchText, ins, notins, me.hidden, me.by)
}
