// Copyright © 2025 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import "embed"

//go:embed Version.dat
var Version string

//go:embed images/*.svg
var ICONS embed.FS

const (
	APPNAME = "CardIndex"
	DOMAIN  = "MNS"

	TIMEOUT_LONG = 10000

	CONFIG_WINDOW           = "Window"
	CONFIG_WINDOW_STATE     = "State"
	CONFIG_WINDOW_GEOMETRY  = "Geometry"
	CONFIG_CURSOR_BLINK     = "CursorBlink"
	DEFAULT_CURSOR_BLINK    = true
	CONFIG_MOST_RECENT_FILE = "MostRecentFile"

	SVG_ICON                 = "cardindex.svg"
	SVG_FILE_NEW             = "document-new.svg"
	SVG_FILE_OPEN            = "document-open.svg"
	SVG_FILE_SAVE            = "document-save.svg"
	SVG_FILE_SAVE_AS         = "document-save-as.svg"
	SVG_FILE_EXPORT          = "export.svg"
	SVG_FILE_CONFIGURE       = "document-properties.svg"
	SVG_FILE_QUIT            = "shutdown.svg"
	SVG_EDIT_COPY            = "edit-copy.svg"
	SVG_EDIT_CUT             = "edit-cut.svg"
	SVG_EDIT_PASTE           = "edit-paste.svg"
	SVG_CARD_NEW             = "card-new.svg"
	SVG_CARD_ADD_TO_BOX      = "card-add-to-box.svg"
	SVG_CARD_REMOVE_FROM_BOX = "card-remove-from-box.svg"
	SVG_CARD_EXPORT          = "card-export.svg"
	SVG_CARD_UNHIDE          = "card-unhide.svg"
	SVG_CARD_HIDE            = "card-hide.svg"
	SVG_CARD_DELETE          = "card-delete.svg"
	SVG_BOX_NEW              = "box-new.svg"
	SVG_BOX_ADD_FROM_SEARCH  = "box-add-from-search.svg"
	SVG_BOX_ADD_FROM_BOX     = "box-add-from-box.svg"
	SVG_BOX_DELETE           = "box-delete.svg"
	SVG_HELP_HELP            = "help.svg"
	SVG_HELP_ABOUT           = "about.svg"
)

// var E100 = errors.New("E100: failed to get counts")
