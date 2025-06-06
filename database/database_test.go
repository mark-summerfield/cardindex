// Copyright © 2025 Mark Summerfield. All rights reserved.
// License: GPL-3

package database

import (
	"os"
	"runtime"
	"slices"
	"strings"
	"testing"
	"time"
)

func Test_Database_Empty(t *testing.T) {
	filename := os.TempDir() + "/empty.cix"
	os.Remove(filename)
	database, err := NewDatabase(filename)
	checkErr(t, err)
	defer database.Close()
	version, err := SqliteVersion()
	checkErr(t, err)
	if !strings.HasPrefix(version, "3.") {
		t.Errorf("expected version 3.x.y; got : %s", version)
	}
	if database.Filename() != filename {
		t.Errorf("expected filename %q; got: %q", filename,
			database.Filename())
	}
	counts, err := database.CardCounts()
	checkErr(t, err)
	checkCardCounts(t, &CardCounts{0, 0, 0}, &counts)
	when, err := database.ConfigCreated()
	checkErr(t, err)
	year := time.Now().Year()
	if when.Year() != year {
		t.Errorf("invalid year expected %d; got %d", year, when.Year())
	}
	when, err = database.ConfigUpdated()
	checkErr(t, err)
	if when.Year() != year {
		t.Errorf("invalid year expected %d; got %d", year, when.Year())
	}
}

func Test_Database_Cix(t *testing.T) {
	database, err := NewDatabase("../eg/pcw.cix")
	checkErr(t, err)
	defer database.Close()
	counts, err := database.CardCounts()
	checkErr(t, err)
	checkCardCounts(t, &CardCounts{29, 2, 0}, &counts)
	for i, expected := range []string{"1978", "1979", "1980"} {
		box, err := database.Box(i + 1)
		checkErr(t, err)
		if box.name != expected {
			t.Errorf("expected box %q; got: %q", expected, box.name)
		}
	}
}

func Test_Database_New1(t *testing.T) {
	filename := os.TempDir() + "/new1.cix"
	os.Remove(filename)
	database, err := NewDatabase(filename)
	checkErr(t, err)
	defer database.Close()
	if database.Filename() != filename {
		t.Errorf("expected filename %q; got: %q", filename,
			database.Filename())
	}
	counts, err := database.CardCounts()
	checkErr(t, err)
	checkCardCounts(t, &CardCounts{0, 0, 0}, &counts)
	for i, body := range []string{
		"A Title\nThe first line.",
		"Another Title\nAnother first line.",
		"Yet another title\nAnd another first line.",
		"A title with no first line. Instead two sentences.",
	} {
		cid, err := database.CardAdd(body)
		checkErr(t, err)
		if cid != i+1 {
			t.Errorf("expected cid %d; got: %d", i+1, cid)
		}
		counts, err = database.CardCounts()
		checkErr(t, err)
		checkCardCounts(t, &CardCounts{i + 1, i + 1, 0}, &counts)
	}
	err = database.CardEdit(3, "YET Another Title\nWith another first line.")
	checkErr(t, err)
	cid := 2
	card2a, err := database.Card(cid)
	checkErr(t, err)
	if err = database.CardDelete(cid); err == nil {
		t.Error("expected error (can't delete unless hidden)")
	}
	card2b, err := database.Card(cid)
	checkErr(t, err)
	if card2a.body == "" || (card2a.body != card2b.body) {
		t.Errorf("expected card; got: %s", card2b)
	}
	counts, err = database.CardCounts()
	checkErr(t, err)
	for cid := range counts.Visible + counts.Hidden {
		if cid == 0 {
			continue
		}
		hidden, err := database.CardHidden(cid)
		checkErr(t, err)
		if hidden {
			t.Errorf("expected card cid %d to be visible", cid)
		}
	}
	err = database.CardHide(cid)
	checkErr(t, err)
	hidden, err := database.CardHidden(cid)
	checkErr(t, err)
	if !hidden {
		t.Errorf("expected card cid %d to be hidden", cid)
	}
	err = database.CardUnhide(cid)
	checkErr(t, err)
	hidden, err = database.CardHidden(cid)
	checkErr(t, err)
	if hidden {
		t.Errorf("expected card cid %d to be visible", cid)
	}
	err = database.CardHide(cid)
	checkErr(t, err)
	hidden, err = database.CardHidden(cid)
	checkErr(t, err)
	if !hidden {
		t.Errorf("expected card cid %d to be hidden", cid)
	}
	err = database.CardDelete(cid)
	checkErr(t, err)
	if _, err = database.Card(cid); err == nil {
		t.Errorf("expected error for deleted card cid %d", cid)
	}
}

func Test_Database_New2(t *testing.T) {
	filename := os.TempDir() + "/new2.cix"
	os.Remove(filename)
	database, err := NewDatabase(filename)
	checkErr(t, err)
	defer database.Close()
	counts, err := database.CardCounts()
	checkErr(t, err)
	checkCardCounts(t, &CardCounts{0, 0, 0}, &counts)
	cardnames, err := database.CardNamesVisible(OID_NAME)
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
		cid, err := database.CardAdd(body)
		checkErr(t, err)
		if cid != i+1 {
			t.Errorf("expected cid %d; got: %d", i+1, cid)
		}
		counts, err = database.CardCounts()
		checkErr(t, err)
		checkCardCounts(t, &CardCounts{i + 1, i + 1, 0}, &counts)
	}
	cardnames, err = database.CardNamesVisible(OID_UPDATED)
	checkErr(t, err)
	if len(cardnames) != counts.Visible {
		t.Errorf("expected %d cardnames; got: %d", counts.Visible,
			len(cardnames))
	}
	cardnames, err = database.CardNamesUnboxed(OID_CREATED)
	checkErr(t, err)
	if len(cardnames) != counts.Unboxed {
		t.Errorf("expected %d cardnames; got: %d", counts.Unboxed,
			len(cardnames))
	}
	cardnames, err = database.CardNamesHidden(OID_NAME)
	checkErr(t, err)
	if len(cardnames) != counts.Hidden {
		t.Errorf("expected %d cardnames; got: %d", counts.Hidden,
			len(cardnames))
	}
	boxes, err := database.Boxes()
	checkErr(t, err)
	if len(boxes) > 0 {
		t.Errorf("expected no boxes; got: %d", len(boxes))
	}
	bid1, err := database.BoxAdd("Special Box")
	checkErr(t, err)
	boxes, err = database.Boxes()
	checkErr(t, err)
	if len(boxes) != 1 {
		t.Errorf("expected 1 box; got: %d", len(boxes))
	}
	bid2, err := database.BoxAdd("Ordinary Box")
	checkErr(t, err)
	boxes, err = database.Boxes()
	checkErr(t, err)
	if len(boxes) != 2 {
		t.Errorf("expected 2 boxes; got: %d", len(boxes))
	}
	box, err := database.Box(bid1)
	checkErr(t, err)
	name := "Special Box"
	if box.name != name {
		t.Errorf("expected box %q; got: %q", name, box.name)
	}
	box, err = database.Box(bid2)
	checkErr(t, err)
	name = "Ordinary Box"
	if box.name != name {
		t.Errorf("expected box %q; got: %q", name, box.name)
	}
	for _, bid := range []int{bid1, bid2} {
		in_use, err := database.BoxInUse(bid)
		checkErr(t, err)
		if in_use {
			t.Errorf("expected box %d to be not in use", bid)
		}
	}
	err = database.BoxDelete(bid1)
	checkErr(t, err)
	boxes, err = database.Boxes()
	checkErr(t, err)
	if len(boxes) != 1 {
		t.Errorf("expected 1 boxes; got: %d", len(boxes))
	}
	cid := 3
	err = database.BoxAddCards(bid2, cid)
	checkErr(t, err)
	if err = database.BoxDelete(bid2); err == nil {
		t.Errorf("expected error deleting box %d", bid2)
	}
	in_use, err := database.BoxInUse(bid2)
	checkErr(t, err)
	if !in_use {
		t.Errorf("expected box %d to be in use", bid2)
	}
	err = database.BoxRemoveCard(bid2, cid)
	checkErr(t, err)
	in_use, err = database.BoxInUse(bid2)
	checkErr(t, err)
	if in_use {
		t.Errorf("expected box %d to not be in use", bid2)
	}
	err = database.BoxDelete(bid2)
	checkErr(t, err)
	boxes, err = database.Boxes()
	checkErr(t, err)
	if len(boxes) > 0 {
		t.Errorf("expected no boxes; got: %d", len(boxes))
	}
}

func Test_Database_Search1(t *testing.T) {
	filename := os.TempDir() + "/search1.cix"
	os.Remove(filename)
	database, err := NewDatabase(filename)
	checkErr(t, err)
	defer database.Close()
	counts, err := database.CardCounts()
	checkErr(t, err)
	checkCardCounts(t, &CardCounts{0, 0, 0}, &counts)
	cardnames, err := database.CardNamesVisible(OID_NAME)
	checkErr(t, err)
	if len(cardnames) > 0 {
		t.Errorf("expected 0 cardnames; got: %d", len(cardnames))
	}
	for i, body := range []string{
		"A Title\nThe first line. Red",
		"Another Title\nAnother first line. Green",
		"Yet another title\nAnd another first line. Blue",
		"A title with no first line. Instead two sentences. red",
	} {
		cid, err := database.CardAdd(body)
		checkErr(t, err)
		if cid != i+1 {
			t.Errorf("expected cid %d; got: %d", i+1, cid)
		}
		counts, err = database.CardCounts()
		checkErr(t, err)
		checkCardCounts(t, &CardCounts{i + 1, i + 1, 0}, &counts)
	}
	boxes, err := database.Boxes()
	checkErr(t, err)
	if len(boxes) > 0 {
		t.Errorf("expected no boxes; got: %d", len(boxes))
	}
	_, err = database.BoxAdd("Special Box")
	checkErr(t, err)
	boxes, err = database.Boxes()
	checkErr(t, err)
	if len(boxes) != 1 {
		t.Errorf("expected 1 box; got: %d", len(boxes))
	}
	_, err = database.BoxAdd("Ordinary Box")
	checkErr(t, err)
	boxes, err = database.Boxes()
	checkErr(t, err)
	if len(boxes) != 2 {
		t.Errorf("expected 2 boxes; got: %d", len(boxes))
	}

	estrings := []string{}
	search := NewSearch("", false, OID_NAME)
	sid1, err := database.SearchAdd(search)
	checkErr(t, err)
	search, err = database.Search(sid1)
	checkErr(t, err)
	expected := `Search#1 searchText="" hidden=false oid=1(Name)`
	estrings = append(estrings, expected)
	if search.String() != expected {
		t.Errorf("expected search %q; got: %q", expected, search)
	}

	search = NewSearch("Red", false, OID_NAME)
	sid2, err := database.SearchAdd(search)
	checkErr(t, err)
	search, err = database.Search(sid2)
	checkErr(t, err)
	expected = `Search#2 searchText="Red" hidden=false oid=1(Name)`
	estrings = append(estrings, expected)
	if search.String() != expected {
		t.Errorf("expected search %q; got: %q", expected, search)
	}

	search = NewSearch("", true, OID_UPDATED)
	sid3, err := database.SearchAdd(search)
	checkErr(t, err)
	search, err = database.Search(sid3)
	checkErr(t, err)
	expected = `Search#3 searchText="" hidden=true oid=2(Updated)`
	estrings = append(estrings, expected)
	if search.String() != expected {
		t.Errorf("expected search %q; got: %q", expected, search)
	}

	searches, err := database.Searches()
	checkErr(t, err)
	for _, search := range searches {
		s := search.String()
		n := 0
		for _, estring := range estrings {
			if s == estring {
				n++
			}
		}
		if n != 1 {
			t.Errorf("expected search %q", s)
		}
	}

	expectedCardNames := [][]CardName{
		{
			{1, "A Title"},
			{4, "A title with no first line."},
			{2, "Another Title"},
			{3, "Yet another title"},
		},
		{
			{1, "A Title"},
			{4, "A title with no first line."},
		},
		{},
	}
	for i, sid := range []int{sid1, sid2, sid3} {
		if cardnames, err := database.CardNamesForSid(sid); err != nil {
			t.Errorf("unexpected error %s", err)
		} else {
			expectednames := expectedCardNames[i]
			if !slices.Equal(cardnames, expectednames) {
				t.Errorf("expected search results #%d of\n%v; got:\n%v",
					i, expectednames, cardnames)
			}
		}
	}
}

func Test_Database_Search2(t *testing.T) {
	search := NewSearch("", false, OID_NAME)
	expected := "SELECT cid, Name FROM Cards WHERE hidden = FALSE " +
		"ORDER BY LOWER(Name);"
	query, args := search.Query(true)
	if query != expected {
		t.Errorf("expected search %q; got: %q", expected, query)
	}
	if args != nil {
		t.Error("expected no args")
	}
	search = NewSearch("Red", false, OID_NAME)
	expected = "SELECT cid, Name FROM Cards WHERE hidden = FALSE AND " +
		"cid IN (SELECT ROWID FROM vt_fts_cards(?)) ORDER BY LOWER(Name);"
	query, args = search.Query(true)
	if query != expected {
		t.Errorf("expected search %q; got: %q", expected, query)
	}
	if len(args) != 1 {
		t.Error("expected 1 arg")
	} else {
		if searchText, ok := args[0].(string); !ok {
			t.Errorf("expected 1 string; got %T %v", args[0], args[0])
		} else if searchText != "Red" {
			t.Error("expected arg \"Red\"")
		}
	}
	search = NewSearch("", true, OID_UPDATED)
	expected = "SELECT cid, Name FROM Cards WHERE hidden = TRUE " +
		"ORDER BY updated DESC;"
	query, args = search.Query(true)
	if query != expected {
		t.Errorf("expected search %q; got: %q", expected, query)
	}
	if args != nil {
		t.Error("expected no args")
	}
}

func Test_Database_Search3(t *testing.T) {
	filename := os.TempDir() + "/search3.cix"
	os.Remove(filename)
	database, err := NewDatabase(filename)
	checkErr(t, err)
	defer database.Close()
	counts, err := database.CardCounts()
	checkErr(t, err)
	checkCardCounts(t, &CardCounts{0, 0, 0}, &counts)
	cardnames, err := database.CardNamesVisible(OID_NAME)
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
		cid, err := database.CardAdd(body)
		checkErr(t, err)
		if cid != i+1 {
			t.Errorf("expected cid %d; got: %d", i+1, cid)
		}
		counts, err = database.CardCounts()
		checkErr(t, err)
		checkCardCounts(t, &CardCounts{i + 1, i + 1, 0}, &counts)
	}
	search := NewSearch("cafe", false, OID_NAME)
	expected := "SELECT cid, Name FROM Cards WHERE hidden = FALSE AND " +
		"cid IN (SELECT ROWID FROM vt_fts_cards(?)) ORDER BY LOWER(Name);"
	query, args := search.Query(true)
	if query != expected {
		t.Errorf("expected search %q; got: %q", expected, query)
	}
	if len(args) != 1 {
		t.Error("expected 1 arg")
	} else {
		if searchText, ok := args[0].(string); !ok {
			t.Errorf("expected 1 string; got %T %v", args[0], args[0])
		} else if searchText != "cafe" {
			t.Error("expected arg \"cafe\"")
		}
	}
	if cardnames, err := database.CardNamesForSearch(search); err != nil {
		t.Errorf("unexpected error %T %s", err, err)
	} else if len(cardnames) != 0 {
		t.Errorf("unexpected search result: %v", cardnames)
	}
	cafes := []string{"The Blue café", "The lost cafe", "The red Cafe"}
	for _, body := range cafes {
		_, err := database.CardAdd(body)
		checkErr(t, err)
	}
	if cardnames, err := database.CardNamesForSearch(search); err != nil {
		t.Errorf("unexpected error %T %s", err, err)
	} else if len(cardnames) != 3 {
		t.Errorf("expected 3 search results; got: %d", len(cardnames))
	} else {
		for i, cardname := range cardnames {
			if !(i+5 == cardname.cid && cafes[i] == cardname.name) {
				t.Errorf("expected cafe %q; got %q", cafes[i],
					cardname.name)
			}
		}
	}
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
