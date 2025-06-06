// Copyright © 2025 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

type ListKind int

const (
	VISIBLE_LIST_KIND ListKind = -10
	UNBOXED_LIST_KIND ListKind = -11
	HIDDEN_LIST_KIND  ListKind = -12
)

func (me ListKind) String() string {
	switch me {
	case UNBOXED_LIST_KIND:
		return "Unboxed"
	case HIDDEN_LIST_KIND:
		return "Hidden"
	}
	return "Visible"
}

func (me ListKind) IconName() string {
	switch me {
	case UNBOXED_LIST_KIND:
		return SVG_VIEW_UNBOXED
	case HIDDEN_LIST_KIND:
		return SVG_VIEW_HIDDEN
	}
	return SVG_VIEW_VISIBLE
}
