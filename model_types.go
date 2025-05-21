// Copyright © 2025 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/mark-summerfield/ufunc"
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

func (me Query) String() string { // for debugging
	ins := strings.Join(slices.Collect(ufunc.Map(me.inBoxes,
		func(i int) (string, bool) {
			return strconv.Itoa(i), true
		})), ",")
	notins := strings.Join(slices.Collect(ufunc.Map(me.notInBoxes,
		func(i int) (string, bool) {
			return strconv.Itoa(i), true
		})), ",")
	return fmt.Sprintf(
		"Query#%d %q match=%q in=%s not-in=%s hidden=%t by=%s", me.qid,
		me.name, me.matchText, ins, notins, me.hidden, me.by)
}
