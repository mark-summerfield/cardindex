// Copyright © 2025 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	qt "github.com/mappu/miqt/qt6"
)

func (me *App) MakeMainWindow() {
	me.MakeActions()
	me.MakeMainMenu()
	me.MakeToolbars()
	me.MakeWidgets()
	me.MakeConnections()
}

func (me *App) MakeActions() {
	me.fileNewAction = qt.NewQAction3(getIcon(SVG_FILE_NEW), "&New…")
	me.fileNewAction.SetShortcutsWithShortcuts(qt.QKeySequence__New)
	me.fileOpenAction = qt.NewQAction3(getIcon(SVG_FILE_OPEN), "&Open…")
	me.fileOpenAction.SetShortcutsWithShortcuts(qt.QKeySequence__Open)
	me.fileSaveAction = qt.NewQAction3(getIcon(SVG_FILE_SAVE), "&Save")
	me.fileSaveAction.SetShortcutsWithShortcuts(qt.QKeySequence__Save)
	me.fileSaveAsAction = qt.NewQAction3(getIcon(SVG_FILE_SAVE_AS),
		"Save &As…")
	me.fileExportAction = qt.NewQAction3(getIcon(SVG_FILE_EXPORT),
		"&Export…")
	me.fileConfigureAction = qt.NewQAction3(getIcon(SVG_FILE_CONFIGURE),
		"&Configure…")
	me.fileQuitAction = qt.NewQAction3(getIcon(SVG_FILE_QUIT), "&Quit")
	me.fileQuitAction.SetShortcutsWithShortcuts(qt.QKeySequence__Quit)
	me.editCopyAction = qt.NewQAction3(getIcon(SVG_EDIT_COPY), "&Copy")
	me.editCopyAction.SetShortcutsWithShortcuts(qt.QKeySequence__Copy)
	me.editCutAction = qt.NewQAction3(getIcon(SVG_EDIT_CUT), "C&ut")
	me.editCutAction.SetShortcutsWithShortcuts(qt.QKeySequence__Cut)
	me.editPasteAction = qt.NewQAction3(getIcon(SVG_EDIT_PASTE), "&Paste")
	me.editPasteAction.SetShortcutsWithShortcuts(qt.QKeySequence__Paste)

	// TODO
}

func (me *App) MakeMainMenu() {
	menubar := me.window.MenuBar()
	me.fileMenu = menubar.AddMenuWithTitle("&File")
	me.fileMenu.AddAction(me.fileNewAction)
	me.fileMenu.AddAction(me.fileOpenAction)
	me.fileMenu.AddAction(me.fileSaveAction)
	me.fileMenu.AddAction(me.fileSaveAsAction)
	me.fileMenu.AddAction(me.fileExportAction)
	me.fileMenu.AddSeparator()
	me.fileMenu.AddAction(me.fileConfigureAction)
	me.fileMenu.AddSeparator()
	me.fileMenu.AddAction(me.fileQuitAction)
	me.editMenu = menubar.AddMenuWithTitle("&Edit")
	me.editMenu.AddAction(me.editCopyAction)
	me.editMenu.AddAction(me.editCutAction)
	me.editMenu.AddAction(me.editPasteAction)
	// TODO
}

func (me *App) MakeToolbars() {
	const FILE = "File"
	fileToolbar := me.window.AddToolBarWithTitle(FILE)
	fileToolbar.SetObjectName(*qt.NewQAnyStringView3(FILE))
	fileToolbar.AddAction(me.fileNewAction)
	fileToolbar.AddAction(me.fileOpenAction)
	fileToolbar.AddAction(me.fileSaveAction)
	const EDIT = "Edit"
	editToolbar := me.window.AddToolBarWithTitle(EDIT)
	editToolbar.SetObjectName(*qt.NewQAnyStringView3(EDIT))
	editToolbar.AddAction(me.editCopyAction)
	editToolbar.AddAction(me.editCutAction)
	editToolbar.AddAction(me.editPasteAction)
	// TODO
}

func (me *App) MakeWidgets() {
	me.mdiArea = qt.NewQMdiArea2()
	me.MakeStatusBar()
	me.window.SetCentralWidget(me.mdiArea.QWidget)
}

func (me *App) MakeStatusBar() {
	me.statusIndicator = qt.NewQLabel3("0 Cards • 0 Unboxed")
	me.statusIndicator.SetFrameShadow(qt.QFrame__Sunken)
	me.statusIndicator.SetFrameShape(qt.QFrame__StyledPanel)
	statusbar := me.window.StatusBar()
	statusbar.SetSizeGripEnabled(false)
	statusbar.AddPermanentWidget(me.statusIndicator.QWidget)
	me.StatusMessage("Click File→New or File→Open", TIMEOUT_LONG)
}

func (me *App) MakeConnections() {
	// TODO
	me.fileQuitAction.OnTriggered(func() { me.window.Close() })
	me.window.OnCloseEvent(func(super func(event *qt.QCloseEvent),
		event *qt.QCloseEvent,
	) {
		me.SaveSettings()
	})
}
