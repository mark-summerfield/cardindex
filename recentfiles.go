// Copyright © 2025 Mark Summerfield. All rights reserved.

package main

import (
	"slices"

	"github.com/mark-summerfield/ufile"
)

type RecentFiles struct {
	filenames []string
	maximum   int
}

func NewRecentFiles(maximum int) RecentFiles {
	return RecentFiles{maximum: maximum}
}

func (me *RecentFiles) Add(filename string) {
	filename = ufile.AbsPath(filename)
	if len(me.filenames) > 0 && me.filenames[0] == filename {
		return // found as first nothing to do
	}
	if i := slices.Index(me.filenames,
		filename); i != -1 { // remove if present
		me.filenames = append(me.filenames[:i], me.filenames[i+1:]...)
	}
	me.filenames = append([]string{filename}, me.filenames...) // put first
	if len(me.filenames) > me.maximum {
		me.filenames = me.filenames[:me.maximum]
	}
}

func (me *RecentFiles) Files() []string {
	filenames := make([]string, 0, len(me.filenames))
	for _, filename := range me.filenames {
		if ufile.FileExists(filename) {
			filenames = append(filenames, filename)
		}
	}
	return filenames
}
