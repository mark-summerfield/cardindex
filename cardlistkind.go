// Copyright © 2025 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

type CardListKind int

const (
	CARD_LIST_VISIBLE_KIND CardListKind = -10
	CARD_LIST_UNBOXED_KIND CardListKind = -11
	CARD_LIST_HIDDEN_KIND  CardListKind = -12
)

func (me CardListKind) String() string {
	switch me {
	case CARD_LIST_UNBOXED_KIND:
		return "Unboxed"
	case CARD_LIST_HIDDEN_KIND:
		return "Hidden"
	}
	return "Visible"
}

func (me CardListKind) IconName() string {
	switch me {
	case CARD_LIST_UNBOXED_KIND:
		return SVG_VIEW_UNBOXED
	case CARD_LIST_HIDDEN_KIND:
		return SVG_VIEW_HIDDEN
	}
	return SVG_VIEW_VISIBLE
}
