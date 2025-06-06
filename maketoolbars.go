// Copyright © 2025 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

func (me *App) MakeToolbars() {
	me.makeFileToolbar()
	me.makeEditToolbars()
	me.makeCardToolbar()
	me.makeBoxToolbar()
	me.makeSearchToolbar()
}

func (me *App) makeFileToolbar() {
	const FILE = "File"
	me.fileToolbar = me.window.AddToolBarWithTitle(FILE)
	me.fileToolbar.SetObjectName(FILE)
	me.fileToolbar.QWidget.AddAction(me.fileNewAction)
	me.fileToolbar.QWidget.AddAction(me.fileOpenAction)
	me.fileToolbar.QWidget.AddAction(me.fileSaveAction)
}

func (me *App) makeEditToolbars() {
	const EDIT1 = "EditUndoRedo"
	me.edit1Toolbar = me.window.AddToolBarWithTitle(EDIT1)
	me.edit1Toolbar.SetObjectName(EDIT1)
	me.edit1Toolbar.QWidget.AddAction(me.editUndoAction)
	me.edit1Toolbar.QWidget.AddAction(me.editRedoAction)
	const EDIT2 = "EditCopyPaste"
	me.edit2Toolbar = me.window.AddToolBarWithTitle(EDIT2)
	me.edit2Toolbar.SetObjectName(EDIT2)
	me.edit2Toolbar.QWidget.AddAction(me.editCopyAction)
	me.edit2Toolbar.QWidget.AddAction(me.editCutAction)
	me.edit2Toolbar.QWidget.AddAction(me.editPasteAction)
	const EDIT3 = "EditCharAttrib"
	me.edit3Toolbar = me.window.AddToolBarWithTitle(EDIT3)
	me.edit3Toolbar.SetObjectName(EDIT3)
	me.edit3Toolbar.QWidget.AddAction(me.editBoldAction)
	me.edit3Toolbar.QWidget.AddAction(me.editItalicAction)
	me.edit3Toolbar.QWidget.AddAction(me.editMonospaceAction)
	const EDIT4 = "EditLists"
	me.edit4Toolbar = me.window.AddToolBarWithTitle(EDIT4)
	me.edit4Toolbar.SetObjectName(EDIT4)
	me.edit4Toolbar.QWidget.AddAction(me.editBulletListAction)
	me.edit4Toolbar.QWidget.AddAction(me.editNumberedListAction)
	me.edit4Toolbar.QWidget.AddAction(me.editEndListAction)
	const EDIT5 = "EditInserts"
	me.edit5Toolbar = me.window.AddToolBarWithTitle(EDIT5)
	me.edit5Toolbar.SetObjectName(EDIT5)
	me.edit5Toolbar.QWidget.AddAction(me.editInsertWebLinkAction)
	me.edit5Toolbar.QWidget.AddAction(me.editInsertFileLinkAction)
	me.edit5Toolbar.QWidget.AddAction(me.editInsertSymbolAction)
	me.window.AddToolBarBreak()
}

func (me *App) makeCardToolbar() {
	const CARD = "Card"
	me.cardToolbar = me.window.AddToolBarWithTitle(CARD)
	me.cardToolbar.SetObjectName(CARD)
	me.cardToolbar.QWidget.AddAction(me.cardNewAction)
	me.cardToolbar.QWidget.AddAction(me.cardViewVisibleAction)
	me.cardToolbar.AddSeparator()
	me.cardToolbar.QWidget.AddAction(me.cardAddToBoxAction)
	me.cardToolbar.QWidget.AddAction(me.cardRemoveFromBoxAction)
	me.cardToolbar.AddSeparator()
	me.cardToolbar.QWidget.AddAction(me.cardExportAction)
	me.cardToolbar.AddSeparator()
	me.cardToolbar.QWidget.AddAction(me.cardUnhideAction)
	me.cardToolbar.QWidget.AddAction(me.cardHideAction)
}

func (me *App) makeBoxToolbar() {
	const BOX = "Box"
	me.boxToolbar = me.window.AddToolBarWithTitle(BOX)
	me.boxToolbar.SetObjectName(BOX)
	me.boxToolbar.QWidget.AddAction(me.boxNewAction)
	me.boxToolbar.QWidget.AddAction(me.boxViewAction)
	me.boxToolbar.AddSeparator()
	me.boxToolbar.QWidget.AddAction(me.boxAddFromSearchAction)
	me.boxToolbar.QWidget.AddAction(me.boxAddFromBoxAction)
}

func (me *App) makeSearchToolbar() {
	const SEARCH = "Search"
	me.searchToolbar = me.window.AddToolBarWithTitle(SEARCH)
	me.searchToolbar.SetObjectName(SEARCH)
	me.searchToolbar.QWidget.AddAction(me.searchNewAction)
	me.searchToolbar.QWidget.AddAction(me.searchViewAction)
}
