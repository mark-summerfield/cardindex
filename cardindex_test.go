// Copyright © 2025 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"fmt"
	"os"
	"testing"
)

func Test01(t *testing.T) {
	filename := os.TempDir() + "/t1.cix"
	os.Remove(filename)
	model, err := NewModel(filename)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	counts, err := model.Counts()
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	fmt.Println(counts)
	defer model.Close()
}

func Test02(t *testing.T) {
	model, err := NewModel("eg/pcw.cix")
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	counts, err := model.Counts()
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	fmt.Println(counts)
	defer model.Close()
}
