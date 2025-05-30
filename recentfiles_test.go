// Copyright © 2025 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"slices"
	"testing"

	"github.com/mark-summerfield/ufile"
)

func Test_RecentFiles(t *testing.T) {
	recentfiles := NewRecentFiles(5)
	files := recentfiles.Files()
	if len(files) != 0 {
		t.Errorf("expected 0 files; got %d", len(files))
	}
	expected := []string{ufile.AbsPath("README.md")}
	recentfiles.Add(expected[0])
	files = recentfiles.Files()
	if !slices.Equal(expected, files) {
		t.Errorf("expected\n%v; got\n%v", expected, files)
	}
	FILENAME := ufile.AbsPath("MANIFEST")
	expected = append([]string{ufile.AbsPath(FILENAME)}, expected...)
	recentfiles.Add(expected[0])
	files = recentfiles.Files()
	if !slices.Equal(expected, files) {
		t.Errorf("expected\n%v; got\n%v", expected, files)
	}
	expected = append([]string{ufile.AbsPath("go.mod")}, expected...)
	recentfiles.Add(expected[0])
	files = recentfiles.Files()
	if !slices.Equal(expected, files) {
		t.Errorf("expected\n%v; got\n%v", expected, files)
	}
	expected = append([]string{ufile.AbsPath("build.sh")}, expected...)
	recentfiles.Add(expected[0])
	files = recentfiles.Files()
	if !slices.Equal(expected, files) {
		t.Errorf("expected\n%v; got\n%v", expected, files)
	}
	i := slices.Index(expected, FILENAME)
	expected = append(expected[:i], expected[i+1:]...)
	expected = append([]string{ufile.AbsPath(FILENAME)}, expected...)
	recentfiles.Add(expected[0])
	files = recentfiles.Files()
	if !slices.Equal(expected, files) {
		t.Errorf("expected\n%v; got\n%v", expected, files)
	}
	recentfiles.Add(expected[0])
	files = recentfiles.Files()
	if !slices.Equal(expected, files) {
		t.Errorf("expected\n%v; got\n%v", expected, files)
	}
	expected = append([]string{ufile.AbsPath("st.sh")}, expected...)
	recentfiles.Add(expected[0])
	files = recentfiles.Files()
	if !slices.Equal(expected, files) {
		t.Errorf("expected\n%v; got\n%v", expected, files)
	}
	expected = append([]string{ufile.AbsPath("LICENSE")}, expected...)
	recentfiles.Add(expected[0])
	files = recentfiles.Files()
	expected = expected[:recentfiles.maximum]
	if !slices.Equal(expected, files) {
		t.Errorf("expected\n%v; got\n%v", expected, files)
	}
	expected = append([]string{ufile.AbsPath("no such file")}, expected...)
	recentfiles.Add(expected[0])
	files = recentfiles.Files()
	expected = expected[1:recentfiles.maximum]
	if !slices.Equal(expected, files) {
		t.Errorf("expected\n%v; got\n%v", expected, files)
	}
}
