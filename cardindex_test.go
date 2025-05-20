// Copyright © 2025 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"os"
	"strings"
	"testing"
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
