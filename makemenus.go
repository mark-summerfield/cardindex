// Copyright © 2025 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"github.com/mappu/miqt/qt"
)

func (me *App) MakeMainMenu() {
	menubar := me.window.MenuBar()
	me.makeMainFileMenu(menubar)
	me.makeMainEditMenu(menubar)
	me.makeMainCardMenu(menubar)
	me.makeMainBoxMenu(menubar)
	me.makeMainSearchMenu(menubar)
	me.makeMainWindowMenu(menubar)
	me.makeMainHelpMenu(menubar)
}

func (me *App) makeMainFileMenu(menubar *qt.QMenuBar) {
	me.fileMenu = menubar.AddMenuWithTitle("&File")
	me.fileMenu.QWidget.AddAction(me.fileNewAction)
	me.fileMenu.QWidget.AddAction(me.fileOpenAction)
	me.fileMenu.QWidget.AddAction(me.fileSaveAction)
	me.fileMenu.QWidget.AddAction(me.fileSaveAsAction)
	me.fileMenu.QWidget.AddAction(me.fileExportAction)
	me.fileMenu.AddSeparator()
	me.fileMenu.QWidget.AddAction(me.fileConfigureAction)
	me.fileMenu.AddSeparator()
	me.fileMenu.QWidget.AddAction(me.fileQuitAction)
	me.fileMenu.AddSeparator()
	for _, action := range me.fileOpenActions {
		me.fileMenu.QWidget.AddAction(action)
	}
}

func (me *App) makeMainEditMenu(menubar *qt.QMenuBar) {
	me.editMenu = menubar.AddMenuWithTitle("&Edit")
	me.editMenu.QWidget.AddAction(me.editUndoAction)
	me.editMenu.QWidget.AddAction(me.editRedoAction)
	me.editMenu.AddSeparator()
	me.editMenu.QWidget.AddAction(me.editCopyAction)
	me.editMenu.QWidget.AddAction(me.editCutAction)
	me.editMenu.QWidget.AddAction(me.editPasteAction)
	me.editMenu.AddSeparator()
	me.editMenu.QWidget.AddAction(me.editBoldAction)
	me.editMenu.QWidget.AddAction(me.editItalicAction)
	me.editMenu.QWidget.AddAction(me.editMonospaceAction)
	me.editMenu.AddSeparator()
	me.editMenu.QWidget.AddAction(me.editBulletListAction)
	me.editMenu.QWidget.AddAction(me.editNumberedListAction)
	me.editMenu.QWidget.AddAction(me.editEndListAction)
	me.editMenu.AddSeparator()
	me.editMenu.QWidget.AddAction(me.editInsertWebLinkAction)
	me.editMenu.QWidget.AddAction(me.editInsertFileLinkAction)
	me.editMenu.QWidget.AddAction(me.editInsertSymbolAction)
}

func (me *App) makeMainCardMenu(menubar *qt.QMenuBar) {
	me.cardMenu = menubar.AddMenuWithTitle("&Card")
	me.cardMenu.QWidget.AddAction(me.cardNewAction)
	me.cardMenu.QWidget.AddAction(me.cardViewVisibleAction)
	me.cardMenu.QWidget.AddAction(me.cardViewUnboxedAction)
	me.cardMenu.QWidget.AddAction(me.cardViewHiddenAction)
	me.cardMenu.AddSeparator()
	me.cardMenu.QWidget.AddAction(me.cardAddToBoxAction)
	me.cardMenu.QWidget.AddAction(me.cardRemoveFromBoxAction)
	me.cardMenu.AddSeparator()
	me.cardMenu.QWidget.AddAction(me.cardExportAction)
	me.cardMenu.AddSeparator()
	me.cardMenu.QWidget.AddAction(me.cardUnhideAction)
	me.cardMenu.QWidget.AddAction(me.cardHideAction)
	me.cardMenu.AddSeparator()
	me.cardMenu.QWidget.AddAction(me.cardDeleteAction)
}

func (me *App) makeMainBoxMenu(menubar *qt.QMenuBar) {
	me.boxMenu = menubar.AddMenuWithTitle("&Box")
	me.boxMenu.QWidget.AddAction(me.boxNewAction)
	me.boxMenu.QWidget.AddAction(me.boxViewAction)
	me.boxMenu.AddSeparator()
	me.boxMenu.QWidget.AddAction(me.boxAddFromSearchAction)
	me.boxMenu.QWidget.AddAction(me.boxAddFromBoxAction)
	me.boxMenu.AddSeparator()
	me.boxMenu.QWidget.AddAction(me.boxDeleteAction)
}

func (me *App) makeMainSearchMenu(menubar *qt.QMenuBar) {
	me.searchMenu = menubar.AddMenuWithTitle("&Search")
	me.searchMenu.QWidget.AddAction(me.searchNewAction)
	me.searchMenu.QWidget.AddAction(me.searchViewAction)
	me.searchMenu.AddSeparator()
	me.searchMenu.QWidget.AddAction(me.searchDeleteAction)
}

func (me *App) makeMainWindowMenu(menubar *qt.QMenuBar) {
	me.windowMenu = menubar.AddMenuWithTitle("&Window")
	me.windowMenu.QWidget.AddAction(me.windowNextAction)
	me.windowMenu.QWidget.AddAction(me.windowPrevAction)
	me.windowMenu.QWidget.AddAction(me.windowCascadeAction)
	me.windowMenu.QWidget.AddAction(me.windowTileAction)
	me.windowMenu.AddSeparator()
	me.windowMenu.QWidget.AddAction(me.windowCloseAction)
}

func (me *App) makeMainHelpMenu(menubar *qt.QMenuBar) {
	me.helpMenu = menubar.AddMenuWithTitle("&Help")
	me.helpMenu.QWidget.AddAction(me.helpHelpAction)
	me.helpMenu.QWidget.AddAction(me.helpAboutAction)
}
