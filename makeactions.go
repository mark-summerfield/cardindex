// Copyright © 2025 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"github.com/mappu/miqt/qt"
)

func (me *App) MakeActions() {
	me.makeFileActions()
	me.makeEditActions()
	me.makeCardActions()
	me.makeBoxActions()
	me.makeSearchActions()
	me.makeWindowActions()
	me.makeHelpActions()
}

func (me *App) makeFileActions() {
	me.fileNewAction = qt.NewQAction3(getIcon(SVG_FILE_NEW), "&New…")
	me.fileNewAction.SetToolTip("Create a new card index")
	me.fileNewAction.SetShortcutsWithShortcuts(qt.QKeySequence__New)
	me.fileOpenAction = qt.NewQAction3(getIcon(SVG_FILE_OPEN), "&Open…")
	me.fileOpenAction.SetToolTip("Open an existing card index")
	me.fileOpenAction.SetShortcutsWithShortcuts(qt.QKeySequence__Open)
	me.fileSaveAction = qt.NewQAction3(getIcon(SVG_FILE_SAVE), "&Save")
	me.fileSaveAction.SetToolTip("Save unsaved changes")
	me.fileSaveAction.SetShortcutsWithShortcuts(qt.QKeySequence__Save)
	me.fileSaveAsAction = qt.NewQAction3(getIcon(SVG_FILE_SAVE_AS),
		"Save &As…")
	me.fileSaveAsAction.SetToolTip(
		"Save a copy of the card index and switch to using the copy")
	me.fileExportAction = qt.NewQAction3(getIcon(SVG_FILE_EXPORT),
		"&Export…")
	me.fileExportAction.SetToolTip("Export the card index")
	me.fileConfigureAction = qt.NewQAction3(getIcon(SVG_FILE_CONFIGURE),
		"&Configure…")
	me.fileConfigureAction.SetToolTip(
		"Configure the card index application")
	me.fileQuitAction = qt.NewQAction3(getIcon(SVG_FILE_QUIT), "&Quit")
	me.fileQuitAction.SetToolTip(
		"Save unsaved changed and quit the card index application")
	me.fileQuitAction.SetShortcutsWithShortcuts(qt.QKeySequence__Quit)
	for range MAX_RECENT_FILES {
		action := qt.NewQAction3(getIcon(SVG_FILE_OPEN), "")
		action.SetToolTip("Open an existing card index")
		action.SetVisible(false)
		me.fileOpenActions = append(me.fileOpenActions, action)
	}
}

func (me *App) makeEditActions() {
	me.editUndoAction = qt.NewQAction3(getIcon(SVG_EDIT_UNDO), "&Undo")
	me.editUndoAction.SetToolTip("Undo the last card text editing action")
	me.editUndoAction.SetShortcutsWithShortcuts(qt.QKeySequence__Undo)
	me.editRedoAction = qt.NewQAction3(getIcon(SVG_EDIT_REDO), "&Redo")
	me.editRedoAction.SetToolTip("Redo the last card text editing action")
	me.editRedoAction.SetShortcutsWithShortcuts(qt.QKeySequence__Redo)
	me.editCopyAction = qt.NewQAction3(getIcon(SVG_EDIT_COPY), "&Copy")
	me.editCopyAction.SetToolTip("Copy the selected text to the clipboard")
	me.editCopyAction.SetShortcutsWithShortcuts(qt.QKeySequence__Copy)
	me.editCutAction = qt.NewQAction3(getIcon(SVG_EDIT_CUT), "Cu&t")
	me.editCutAction.SetToolTip("Cut the selected text to the clipboard")
	me.editCutAction.SetShortcutsWithShortcuts(qt.QKeySequence__Cut)
	me.editPasteAction = qt.NewQAction3(getIcon(SVG_EDIT_PASTE), "&Paste")
	me.editPasteAction.SetToolTip(
		"Paste the clipboard text at the cursor position")
	me.editPasteAction.SetShortcutsWithShortcuts(qt.QKeySequence__Paste)
	me.editBoldAction = qt.NewQAction3(getIcon(SVG_EDIT_BOLD), "&Bold")
	me.editBoldAction.SetToolTip("Toggle bold text")
	me.editBoldAction.SetShortcutsWithShortcuts(qt.QKeySequence__Bold)
	me.editItalicAction = qt.NewQAction3(getIcon(SVG_EDIT_ITALIC),
		"&Italic")
	me.editItalicAction.SetToolTip("Toggle italic text")
	me.editItalicAction.SetShortcutsWithShortcuts(qt.QKeySequence__Italic)
	me.editMonospaceAction = qt.NewQAction3(
		getIcon(SVG_EDIT_MONOSPACE), "&Monospace")
	me.editMonospaceAction.SetToolTip("Toggle monospace text")
	me.editMonospaceAction.SetShortcut(qt.NewQKeySequence2("Ctrl+M"))
	me.editBulletListAction = qt.NewQAction3(getIcon(SVG_EDIT_BULLET_LIST),
		"Bulleted &List")
	me.editBulletListAction.SetToolTip("Start bullet list")
	me.editNumberedListAction = qt.NewQAction3(
		getIcon(SVG_EDIT_NUMBERED_LIST), "&Numbered List")
	me.editNumberedListAction.SetToolTip("Start numbered list")
	me.editEndListAction = qt.NewQAction3(getIcon(SVG_EDIT_END_LIST),
		"&End List")
	me.editEndListAction.SetToolTip("End list")
	me.editInsertWebLinkAction = qt.NewQAction3(
		getIcon(SVG_EDIT_INSERT_WEB_LINK), "Insert &Web link…")
	me.editInsertWebLinkAction.SetToolTip("Insert web link (URL)")
	me.editInsertFileLinkAction = qt.NewQAction3(
		getIcon(SVG_EDIT_INSERT_FILE_LINK), "Insert &File link…")
	me.editInsertFileLinkAction.SetToolTip("Insert file link (filename)")
	me.editInsertSymbolAction = qt.NewQAction3(
		getIcon(SVG_EDIT_INSERT_SYMBOL), "Insert &Symbol…")
	me.editInsertSymbolAction.SetToolTip("Insert symbol")
}

func (me *App) makeCardActions() {
	me.cardNewAction = qt.NewQAction3(getIcon(SVG_CARD_NEW), "&New")
	me.cardNewAction.SetToolTip("Create a new card")
	me.cardNewAction.SetShortcut(qt.NewQKeySequence3(int(qt.Key_F7)))
	me.cardViewVisibleAction = qt.NewQAction3(getIcon(SVG_VIEW_VISIBLE),
		"&View Visible Cards")
	me.cardViewVisibleAction.SetToolTip("View all visible cards")
	me.cardViewUnboxedAction = qt.NewQAction3(getIcon(SVG_VIEW_UNBOXED),
		"View Unboxed &Cards")
	me.cardViewUnboxedAction.SetToolTip("View all unboxed cards")
	me.cardViewHiddenAction = qt.NewQAction3(getIcon(SVG_VIEW_HIDDEN),
		"V&iew Hidden Cards")
	me.cardViewHiddenAction.SetToolTip("View all hidden cards")
	me.cardAddToBoxAction = qt.NewQAction3(getIcon(SVG_CARD_ADD_TO_BOX),
		"&Add to Box…")
	me.cardAddToBoxAction.SetToolTip("Add the current card to a box")
	me.cardRemoveFromBoxAction = qt.NewQAction3(
		getIcon(SVG_CARD_REMOVE_FROM_BOX), "&Remove from Box…")
	me.cardRemoveFromBoxAction.SetToolTip(
		"Remove the current card from a box")
	me.cardExportAction = qt.NewQAction3(getIcon(SVG_CARD_EXPORT),
		"&Export…")
	me.cardExportAction.SetToolTip("Export the current card")
	me.cardUnhideAction = qt.NewQAction3(getIcon(SVG_CARD_UNHIDE),
		"&Unhide")
	me.cardUnhideAction.SetToolTip("Unhide the current card")
	me.cardHideAction = qt.NewQAction3(getIcon(SVG_CARD_HIDE), "&Hide")
	me.cardHideAction.SetToolTip("Hide the current card")
	me.cardDeleteAction = qt.NewQAction3(getIcon(SVG_CARD_DELETE),
		"&Delete…")
	me.cardDeleteAction.SetToolTip(
		"Permanently delete the current card: it is safer to Hide cards")
}

func (me *App) makeBoxActions() {
	me.boxNewAction = qt.NewQAction3(getIcon(SVG_BOX_NEW), "&New…")
	me.boxNewAction.SetToolTip("Create a new box")
	me.boxNewAction.SetShortcut(qt.NewQKeySequence3(int(qt.Key_F8)))
	me.boxViewAction = qt.NewQAction3(getIcon(SVG_VIEW_BOXES),
		"&View Boxes")
	me.boxViewAction.SetToolTip("View all boxes")
	me.boxAddFromSearchAction = qt.NewQAction3(
		getIcon(SVG_BOX_ADD_FROM_SEARCH), "Add Cards from &Search…")
	me.boxAddFromSearchAction.SetToolTip(
		"Add cards from the specified search to the current box")
	me.boxAddFromBoxAction = qt.NewQAction3(
		getIcon(SVG_BOX_ADD_FROM_BOX), "Add Cards from &Box…")
	me.boxAddFromBoxAction.SetToolTip(
		"Add cards from the specified box to the current box")
	me.boxDeleteAction = qt.NewQAction3(getIcon(SVG_BOX_DELETE), "&Delete…")
	me.boxDeleteAction.SetToolTip("Permanently delete the current box")
}

func (me *App) makeSearchActions() {
	me.searchNewAction = qt.NewQAction3(getIcon(SVG_SEARCH_NEW), "&New")
	me.searchNewAction.SetToolTip("Create a new search")
	me.searchNewAction.SetShortcut(qt.NewQKeySequence3(int(qt.Key_F9)))
	me.searchViewAction = qt.NewQAction3(getIcon(SVG_VIEW_SEARCHES),
		"&View Searches")
	me.searchViewAction.SetToolTip("View all searches")
	me.searchDeleteAction = qt.NewQAction3(getIcon(SVG_SEARCH_DELETE),
		"&Delete…")
	me.searchDeleteAction.SetToolTip(
		"Permanently delete the current search")
}

func (me *App) makeWindowActions() {
	me.windowNextAction = qt.NewQAction3(getIcon(SVG_WINDOW_NEXT), "&Next")
	me.windowNextAction.SetToolTip("Activate next window")
	me.windowNextAction.SetShortcutsWithShortcuts(
		qt.QKeySequence__NextChild)
	me.windowPrevAction = qt.NewQAction3(getIcon(SVG_WINDOW_PREV),
		"&Previous")
	me.windowPrevAction.SetToolTip("Activate previous window")
	me.windowPrevAction.SetShortcutsWithShortcuts(
		qt.QKeySequence__PreviousChild)
	me.windowCascadeAction = qt.NewQAction3(getIcon(SVG_WINDOW_CASCADE),
		"C&ascade")
	me.windowCascadeAction.SetToolTip("Cascade windows")
	me.windowTileAction = qt.NewQAction3(getIcon(SVG_WINDOW_TILE), "&Tile")
	me.windowTileAction.SetToolTip("Tile windows")
	me.windowCloseAction = qt.NewQAction3(getIcon(SVG_WINDOW_CLOSE),
		"&Close")
	me.windowCloseAction.SetToolTip("Close active window")
	me.windowCloseAction.SetShortcutsWithShortcuts(
		qt.QKeySequence__Close)
}

func (me *App) makeHelpActions() {
	me.helpHelpAction = qt.NewQAction3(getIcon(SVG_HELP_HELP), "&Help")
	me.helpHelpAction.SetToolTip("Show help")
	me.helpHelpAction.SetShortcutsWithShortcuts(
		qt.QKeySequence__HelpContents)
	me.helpAboutAction = qt.NewQAction3(getIcon(SVG_HELP_ABOUT), "&About")
	me.helpAboutAction.SetToolTip("Show about box")
}
