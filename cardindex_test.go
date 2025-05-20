// Copyright © 2025 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"
)

func Test01(t *testing.T) {
	filename := os.TempDir() + "/t1.cix"
	os.Remove(filename)
	model, err := NewModel(filename)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	defer model.Close()
	version, err := model.Version()
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if !strings.HasPrefix(version, "3.") {
		t.Errorf("expected version 3.x.y; got : %s", version)
	}
	if model.Filename() != filename {
		t.Errorf("expected filename %q; got: %q", filename,
			model.Filename())
	}
	counts, err := model.Counts()
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	checkCounts(t, &Counts{0, 0, 0}, counts)
	when, err := model.ConfigCreated()
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	year := time.Now().Year()
	if when.Year() != year {
		t.Errorf("invalid year expected %d; got %d", year, when.Year())
	}
	when, err = model.ConfigUpdated()
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if when.Year() != year {
		t.Errorf("invalid year expected %d; got %d", year, when.Year())
	}
}

func Test02(t *testing.T) {
	model, err := NewModel("eg/pcw.cix")
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	defer model.Close()
	counts, err := model.Counts()
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	checkCounts(t, &Counts{28, 2, 0}, counts)
}

func Test03(t *testing.T) {
	filename := os.TempDir() + "/t3.cix"
	os.Remove(filename)
	model, err := NewModel(filename)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	defer model.Close()
	if model.Filename() != filename {
		t.Errorf("expected filename %q; got: %q", filename,
			model.Filename())
	}
	counts, err := model.Counts()
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	checkCounts(t, &Counts{0, 0, 0}, counts)
	for i, body := range []string{
		"A Title\nThe first line.",
		"Another Title\nAnother first line.",
		"Yet another title\nAnd another first line.",
		"A title with no first line. Instead two sentences.",
	} {
		cid, err := model.CardAdd(body)
		if err != nil {
			t.Errorf("unexpected error: %s", err)
		}
		if cid != i+1 {
			t.Errorf("expected cid %d; got: %d", i+1, cid)
		}
		counts, err = model.Counts()
		if err != nil {
			t.Errorf("unexpected error: %s", err)
		}
		checkCounts(t, &Counts{i + 1, i + 1, 0}, counts)
	}
	if err = model.CardEdit(3,
		"YET Another Title\nWith another first line."); err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	cid := 2
	body2a, err := model.CardBody(cid)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if err = model.CardDelete(cid); err == nil {
		t.Error("expected error (can't delete unless hidden)")
	}
	body2b, err := model.CardBody(cid)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if body2a == "" || (body2a != body2b) {
		t.Errorf("expected body; got: %q", body2b)
	}
	counts, err = model.Counts()
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	for cid := range counts.Visible + counts.Hidden {
		if cid == 0 {
			continue
		}
		hidden, err := model.CardHidden(cid)
		if err != nil {
			t.Errorf("unexpected error: %s", err)
		}
		if hidden {
			t.Errorf("expected card cid %d to be visible", cid)
		}
	}
	err = model.CardHide(cid)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	hidden, err := model.CardHidden(cid)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if !hidden {
		t.Errorf("expected card cid %d to be hidden", cid)
	}
	err = model.CardUnhide(cid)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	hidden, err = model.CardHidden(cid)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if hidden {
		t.Errorf("expected card cid %d to be visible", cid)
	}
	err = model.CardHide(cid)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	hidden, err = model.CardHidden(cid)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if !hidden {
		t.Errorf("expected card cid %d to be hidden", cid)
	}
	err = model.CardDelete(cid)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	_, err = model.CardBody(cid)
	if err == nil {
		t.Error("expected error")
	}
}

func Test04(t *testing.T) {
	filename := os.TempDir() + "/t3.cix"
	os.Remove(filename)
	model, err := NewModel(filename)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	defer model.Close()
	counts, err := model.Counts()
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	checkCounts(t, &Counts{0, 0, 0}, counts)
	fmt.Println("expecting error")
	for cardname := range model.AllVisibleCardNames() {
		fmt.Println("#", cardname)
	}
	for i, body := range []string{
		"A Title\nThe first line.",
		"Another Title\nAnother first line.",
		"Yet another title\nAnd another first line.",
		"A title with no first line. Instead two sentences.",
	} {
		cid, err := model.CardAdd(body)
		if err != nil {
			t.Errorf("unexpected error: %s", err)
		}
		if cid != i+1 {
			t.Errorf("expected cid %d; got: %d", i+1, cid)
		}
		counts, err = model.Counts()
		if err != nil {
			t.Errorf("unexpected error: %s", err)
		}
		checkCounts(t, &Counts{i + 1, i + 1, 0}, counts)
	}
	fmt.Println("expecting 4")
	for cardname := range model.AllVisibleCardNames() {
		fmt.Println("#", cardname)
	}
}

func checkCounts(t *testing.T, expected, actual *Counts) {
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
