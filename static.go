// Copyright © 2025 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import "embed"

//go:embed Version.dat
var Version string

//go:embed images/*.svg
var ICONS embed.FS

type OnError func(string)

const (
	APPNAME = "CardIndex"
	DOMAIN  = "MNS"

	MAX_RECENT_FILES = 9
	FILE_FILTER      = "Card Indexes (*.cix)"
	TIMEOUT_LONG     = 10000

	CONFIG_WINDOW           = "Window"
	CONFIG_WINDOW_STATE     = "State"
	CONFIG_WINDOW_GEOMETRY  = "Geometry"
	CONFIG_CURSOR_BLINK     = "CursorBlink"
	DEFAULT_CURSOR_BLINK    = true
	CONFIG_MOST_RECENT_FILE = "MostRecentFile"
	CONFIG_RECENT_FILES     = "RecentFiles"
	CONFIG_RECENT_FILE      = "File"
	CONFIG_MAX_RECENT_FILES = "MaxRecentFiles"

	SVG_ICON                  = "cardindex"
	SVG_FILE_NEW              = "document-new"
	SVG_FILE_OPEN             = "document-open"
	SVG_FILE_SAVE             = "document-save"
	SVG_FILE_SAVE_AS          = "document-save-as"
	SVG_FILE_EXPORT           = "export"
	SVG_FILE_CONFIGURE        = "document-properties"
	SVG_FILE_QUIT             = "shutdown"
	SVG_EDIT_UNDO             = "edit-undo"
	SVG_EDIT_REDO             = "edit-redo"
	SVG_EDIT_COPY             = "edit-copy"
	SVG_EDIT_CUT              = "edit-cut"
	SVG_EDIT_PASTE            = "edit-paste"
	SVG_EDIT_BOLD             = "format-text-bold"
	SVG_EDIT_ITALIC           = "format-text-italic"
	SVG_EDIT_MONOSPACE        = "format-text-mono"
	SVG_EDIT_BULLET_LIST      = "format-bullet-list"
	SVG_EDIT_NUMBERED_LIST    = "format-number-list"
	SVG_EDIT_END_LIST         = "format-no-list"
	SVG_EDIT_INSERT_WEB_LINK  = "web-link"
	SVG_EDIT_INSERT_FILE_LINK = "file-link"
	SVG_EDIT_INSERT_SYMBOL    = "accessories-character-map"
	SVG_CARD_NEW              = "card-new"
	SVG_CARD_ADD_TO_BOX       = "card-add-to-box"
	SVG_CARD_REMOVE_FROM_BOX  = "card-remove-from-box"
	SVG_CARD_EXPORT           = "card-export"
	SVG_CARD_UNHIDE           = "card-unhide"
	SVG_CARD_HIDE             = "card-hide"
	SVG_CARD_DELETE           = "card-delete"
	SVG_BOX_NEW               = "box-new"
	SVG_BOX_ADD_FROM_SEARCH   = "box-add-from-search"
	SVG_BOX_ADD_FROM_BOX      = "box-add-from-box"
	SVG_BOX_DELETE            = "box-delete"
	SVG_VIEW_VISIBLE          = "visible-cards"
	SVG_VIEW_UNBOXED          = "unboxed-cards"
	SVG_VIEW_HIDDEN           = "hidden-cards"
	SVG_VIEW_BOXES            = "boxes"
	SVG_VIEW_SEARCHES         = "searches"
	SVG_HELP_HELP             = "help"
	SVG_HELP_ABOUT            = "about"
	SVG_SEARCH_NEW            = "search-new"
	SVG_SEARCH_DELETE         = "search-delete"
	SVG_WINDOW                = "window"
	SVG_WINDOW_NEXT           = "window-next"
	SVG_WINDOW_PREV           = "window-prev"
	SVG_WINDOW_CASCADE        = "window-cascade"
	SVG_WINDOW_TILE           = "window-tile"
	SVG_WINDOWS               = "windows"
	SVG_WINDOW_CLOSE          = "window-close"
)

// var E100 = errors.New("E100: failed to get counts")
