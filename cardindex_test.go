// Copyright © 2025 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"os"
	"runtime"
	"strings"
	"testing"
	"time"
)

func Test_Empty(t *testing.T) {
	filename := os.TempDir() + "/empty.cix"
	os.Remove(filename)
	model, err := NewModel(filename)
	checkErr(t, err)
	defer model.Close()
	version, err := SqliteVersion()
	checkErr(t, err)
	if !strings.HasPrefix(version, "3.") {
		t.Errorf("expected version 3.x.y; got : %s", version)
	}
	if model.Filename() != filename {
		t.Errorf("expected filename %q; got: %q", filename,
			model.Filename())
	}
	counts, err := model.CardCounts()
	checkErr(t, err)
	checkCardCounts(t, &CardCounts{0, 0, 0}, &counts)
	when, err := model.ConfigCreated()
	checkErr(t, err)
	year := time.Now().Year()
	if when.Year() != year {
		t.Errorf("invalid year expected %d; got %d", year, when.Year())
	}
	when, err = model.ConfigUpdated()
	checkErr(t, err)
	if when.Year() != year {
		t.Errorf("invalid year expected %d; got %d", year, when.Year())
	}
}

func Test_Cix(t *testing.T) {
	model, err := NewModel("eg/pcw.cix")
	checkErr(t, err)
	defer model.Close()
	counts, err := model.CardCounts()
	checkErr(t, err)
	checkCardCounts(t, &CardCounts{28, 2, 0}, &counts)
	for i, expected := range []string{"1978", "1979", "1980"} {
		box, err := model.Box(i + 1)
		checkErr(t, err)
		if box.name != expected {
			t.Errorf("expected box %q; got: %q", expected, box.name)
		}
	}
}

func Test_New1(t *testing.T) {
	filename := os.TempDir() + "/new1.cix"
	os.Remove(filename)
	model, err := NewModel(filename)
	checkErr(t, err)
	defer model.Close()
	if model.Filename() != filename {
		t.Errorf("expected filename %q; got: %q", filename,
			model.Filename())
	}
	counts, err := model.CardCounts()
	checkErr(t, err)
	checkCardCounts(t, &CardCounts{0, 0, 0}, &counts)
	for i, body := range []string{
		"A Title\nThe first line.",
		"Another Title\nAnother first line.",
		"Yet another title\nAnd another first line.",
		"A title with no first line. Instead two sentences.",
	} {
		cid, err := model.CardAdd(body)
		checkErr(t, err)
		if cid != i+1 {
			t.Errorf("expected cid %d; got: %d", i+1, cid)
		}
		counts, err = model.CardCounts()
		checkErr(t, err)
		checkCardCounts(t, &CardCounts{i + 1, i + 1, 0}, &counts)
	}
	err = model.CardEdit(3, "YET Another Title\nWith another first line.")
	checkErr(t, err)
	cid := 2
	card2a, err := model.Card(cid)
	checkErr(t, err)
	if err = model.CardDelete(cid); err == nil {
		t.Error("expected error (can't delete unless hidden)")
	}
	card2b, err := model.Card(cid)
	checkErr(t, err)
	if card2a.body == "" || (card2a.body != card2b.body) {
		t.Errorf("expected card; got: %s", card2b)
	}
	counts, err = model.CardCounts()
	checkErr(t, err)
	for cid := range counts.Visible + counts.Hidden {
		if cid == 0 {
			continue
		}
		hidden, err := model.CardHidden(cid)
		checkErr(t, err)
		if hidden {
			t.Errorf("expected card cid %d to be visible", cid)
		}
	}
	err = model.CardHide(cid)
	checkErr(t, err)
	hidden, err := model.CardHidden(cid)
	checkErr(t, err)
	if !hidden {
		t.Errorf("expected card cid %d to be hidden", cid)
	}
	err = model.CardUnhide(cid)
	checkErr(t, err)
	hidden, err = model.CardHidden(cid)
	checkErr(t, err)
	if hidden {
		t.Errorf("expected card cid %d to be visible", cid)
	}
	err = model.CardHide(cid)
	checkErr(t, err)
	hidden, err = model.CardHidden(cid)
	checkErr(t, err)
	if !hidden {
		t.Errorf("expected card cid %d to be hidden", cid)
	}
	err = model.CardDelete(cid)
	checkErr(t, err)
	if _, err = model.Card(cid); err == nil {
		t.Errorf("expected error for deleted card cid %d", cid)
	}
}

func Test_New2(t *testing.T) {
	filename := os.TempDir() + "/new2.cix"
	os.Remove(filename)
	model, err := NewModel(filename)
	checkErr(t, err)
	defer model.Close()
	counts, err := model.CardCounts()
	checkErr(t, err)
	checkCardCounts(t, &CardCounts{0, 0, 0}, &counts)
	cardnames, err := model.CardNamesVisible(NAME)
	checkErr(t, err)
	if len(cardnames) > 0 {
		t.Errorf("expected 0 cardnames; got: %d", len(cardnames))
	}
	for i, body := range []string{
		"A Title\nThe first line.",
		"Another Title\nAnother first line.",
		"Yet another title\nAnd another first line.",
		"A title with no first line. Instead two sentences.",
	} {
		cid, err := model.CardAdd(body)
		checkErr(t, err)
		if cid != i+1 {
			t.Errorf("expected cid %d; got: %d", i+1, cid)
		}
		counts, err = model.CardCounts()
		checkErr(t, err)
		checkCardCounts(t, &CardCounts{i + 1, i + 1, 0}, &counts)
	}
	cardnames, err = model.CardNamesVisible(UPDATED)
	checkErr(t, err)
	if len(cardnames) != counts.Visible {
		t.Errorf("expected %d cardnames; got: %d", counts.Visible,
			len(cardnames))
	}
	cardnames, err = model.CardNamesUnboxed(CREATED)
	checkErr(t, err)
	if len(cardnames) != counts.Unboxed {
		t.Errorf("expected %d cardnames; got: %d", counts.Unboxed,
			len(cardnames))
	}
	cardnames, err = model.CardNamesHidden(NAME)
	checkErr(t, err)
	if len(cardnames) != counts.Hidden {
		t.Errorf("expected %d cardnames; got: %d", counts.Hidden,
			len(cardnames))
	}
	boxes, err := model.Boxes()
	checkErr(t, err)
	if len(boxes) > 0 {
		t.Errorf("expected no boxes; got: %d", len(boxes))
	}
	bid1, err := model.BoxAdd("Special Box")
	checkErr(t, err)
	boxes, err = model.Boxes()
	checkErr(t, err)
	if len(boxes) != 1 {
		t.Errorf("expected 1 box; got: %d", len(boxes))
	}
	bid2, err := model.BoxAdd("Ordinary Box")
	checkErr(t, err)
	boxes, err = model.Boxes()
	checkErr(t, err)
	if len(boxes) != 2 {
		t.Errorf("expected 2 boxes; got: %d", len(boxes))
	}
	box, err := model.Box(bid1)
	checkErr(t, err)
	name := "Special Box"
	if box.name != name {
		t.Errorf("expected box %q; got: %q", name, box.name)
	}
	box, err = model.Box(bid2)
	checkErr(t, err)
	name = "Ordinary Box"
	if box.name != name {
		t.Errorf("expected box %q; got: %q", name, box.name)
	}
	for _, bid := range []int{bid1, bid2} {
		in_use, err := model.BoxInUse(bid)
		checkErr(t, err)
		if in_use {
			t.Errorf("expected box %d to be not in use", bid)
		}
	}
	err = model.BoxDelete(bid1)
	checkErr(t, err)
	boxes, err = model.Boxes()
	checkErr(t, err)
	if len(boxes) != 1 {
		t.Errorf("expected 1 boxes; got: %d", len(boxes))
	}
	cid := 3
	err = model.BoxAddCard(cid, bid2)
	checkErr(t, err)
	if err = model.BoxDelete(bid2); err == nil {
		t.Errorf("expected error deleting box %d", bid2)
	}
	in_use, err := model.BoxInUse(bid2)
	checkErr(t, err)
	if !in_use {
		t.Errorf("expected box %d to be in use", bid2)
	}
	err = model.BoxRemoveCard(cid, bid2)
	checkErr(t, err)
	in_use, err = model.BoxInUse(bid2)
	checkErr(t, err)
	if in_use {
		t.Errorf("expected box %d to not be in use", bid2)
	}
	err = model.BoxDelete(bid2)
	checkErr(t, err)
	boxes, err = model.Boxes()
	checkErr(t, err)
	if len(boxes) > 0 {
		t.Errorf("expected no boxes; got: %d", len(boxes))
	}
}

func Test_Query(t *testing.T) {
	filename := os.TempDir() + "/query.cix"
	os.Remove(filename)
	model, err := NewModel(filename)
	checkErr(t, err)
	defer model.Close()
	counts, err := model.CardCounts()
	checkErr(t, err)
	checkCardCounts(t, &CardCounts{0, 0, 0}, &counts)
	cardnames, err := model.CardNamesVisible(NAME)
	checkErr(t, err)
	if len(cardnames) > 0 {
		t.Errorf("expected 0 cardnames; got: %d", len(cardnames))
	}
	for i, body := range []string{
		"A Title\nThe first line. Red",
		"Another Title\nAnother first line. Green",
		"Yet another title\nAnd another first line. Blue",
		"A title with no first line. Instead two sentences. Red",
	} {
		cid, err := model.CardAdd(body)
		checkErr(t, err)
		if cid != i+1 {
			t.Errorf("expected cid %d; got: %d", i+1, cid)
		}
		counts, err = model.CardCounts()
		checkErr(t, err)
		checkCardCounts(t, &CardCounts{i + 1, i + 1, 0}, &counts)
	}
	boxes, err := model.Boxes()
	checkErr(t, err)
	if len(boxes) > 0 {
		t.Errorf("expected no boxes; got: %d", len(boxes))
	}
	bid1, err := model.BoxAdd("Special Box")
	checkErr(t, err)
	boxes, err = model.Boxes()
	checkErr(t, err)
	if len(boxes) != 1 {
		t.Errorf("expected 1 box; got: %d", len(boxes))
	}
	bid2, err := model.BoxAdd("Ordinary Box")
	checkErr(t, err)
	boxes, err = model.Boxes()
	checkErr(t, err)
	if len(boxes) != 2 {
		t.Errorf("expected 2 boxes; got: %d", len(boxes))
	}
	estrings := []string{}
	query := NewQuery("Q1", "", []int{bid2}, []int{}, false, NAME)
	qid1, err := model.QueryAdd(query)
	checkErr(t, err)
	query, err = model.Query(qid1)
	checkErr(t, err)
	expected := `Query#1 "Q1" match="" in=[2] not-in=[] hidden=false by=N`
	estrings = append(estrings, expected)
	if query.String() != expected {
		t.Errorf("expected query %q; got: %q", expected, query)
	}

	query = NewQuery("", "Red", []int{}, []int{bid1, bid2}, false, NAME)
	qid2, err := model.QueryAdd(query)
	checkErr(t, err)
	query, err = model.Query(qid2)
	checkErr(t, err)
	expected = `Query#2 "Red" match="Red" in=[] not-in=[1,2] hidden=false by=N`
	estrings = append(estrings, expected)
	if query.String() != expected {
		t.Errorf("expected query %q; got: %q", expected, query)
	}

	query = NewQuery("", "", []int{}, []int{}, true, UPDATED)
	qid3, err := model.QueryAdd(query)
	checkErr(t, err)
	query, err = model.Query(qid3)
	checkErr(t, err)
	expected = `Query#3 "Query #3" match="" in=[] not-in=[] hidden=true by=U`
	estrings = append(estrings, expected)
	if query.String() != expected {
		t.Errorf("expected query %q; got: %q", expected, query)
	}

	queries, err := model.Queries()
	checkErr(t, err)
	for _, query := range queries {
		s := query.String()
		for _, estring := range estrings {
			if s[:15] == estring[:15] {
				if s != estring {
					t.Errorf("expected query %q; got: %q", estring, s)
				}
			}
		}
	}

	// TODO perform & test queries
}

func checkErr(t *testing.T, err error) {
	if err != nil {
		_, _, lino, ok := runtime.Caller(1)
		if !ok {
			lino = 0
		}
		t.Errorf("unexpected error @%d: %s", lino, err)
	}
}

func checkCardCounts(t *testing.T, expected, actual *CardCounts) {
	if expected.Visible != actual.Visible {
		t.Errorf("expected Visible %d; got: %d", expected.Visible,
			actual.Visible)
	}
	if expected.Unboxed != actual.Unboxed {
		t.Errorf("expected Unboxed %d; got: %d", expected.Unboxed,
			actual.Unboxed)
	}
	if expected.Hidden != actual.Hidden {
		t.Errorf("expected Hidden %d; got: %d", expected.Hidden,
			actual.Hidden)
	}
}
