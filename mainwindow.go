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
	me.makeFileActions()
	me.makeEditActions() // TODO complete
	me.makeCardActions()
	me.makeBoxActions()
	// TODO Search View Window Help
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
}

func (me *App) makeEditActions() {
	me.editCopyAction = qt.NewQAction3(getIcon(SVG_EDIT_COPY), "&Copy")
	me.editCopyAction.SetToolTip("Copy the selected text to the clipboard")
	me.editCopyAction.SetShortcutsWithShortcuts(qt.QKeySequence__Copy)
	me.editCutAction = qt.NewQAction3(getIcon(SVG_EDIT_CUT), "C&ut")
	me.editCutAction.SetToolTip("Cut the selected text to the clipboard")
	me.editCutAction.SetShortcutsWithShortcuts(qt.QKeySequence__Cut)
	me.editPasteAction = qt.NewQAction3(getIcon(SVG_EDIT_PASTE), "&Paste")
	me.editPasteAction.SetToolTip(
		"Paste the clipboard text at the cursor position")
	me.editPasteAction.SetShortcutsWithShortcuts(qt.QKeySequence__Paste)
}

func (me *App) makeCardActions() {
	me.cardNewAction = qt.NewQAction3(getIcon(SVG_CARD_NEW), "&New")
	me.cardNewAction.SetToolTip("Create a new card")
	me.cardNewAction.SetShortcut(qt.NewQKeySequence3(int(qt.Key_F7)))
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
	// TODO &Undo
	// TODO &Redo
	me.editMenu.AddAction(me.editCopyAction)
	me.editMenu.AddAction(me.editCutAction)
	me.editMenu.AddAction(me.editPasteAction)
	// TODO &Bold
	// TODO &Italic
	// TODO &Monospace
	// TODO &Bullet List
	// TODO &Numbered List
	// TODO &Clear List
	// TODO Insert &Symbol…

	me.cardMenu = menubar.AddMenuWithTitle("&Card")
	me.cardMenu.AddAction(me.cardNewAction)
	me.cardMenu.AddSeparator()
	me.cardMenu.AddAction(me.cardAddToBoxAction)
	me.cardMenu.AddAction(me.cardRemoveFromBoxAction)
	me.cardMenu.AddSeparator()
	me.cardMenu.AddAction(me.cardExportAction)
	me.cardMenu.AddSeparator()
	me.cardMenu.AddAction(me.cardUnhideAction)
	me.cardMenu.AddAction(me.cardHideAction)
	me.cardMenu.AddSeparator()
	me.cardMenu.AddAction(me.cardDeleteAction)
	me.boxMenu = menubar.AddMenuWithTitle("&Box")
	me.boxMenu.AddAction(me.boxNewAction)
	me.boxMenu.AddSeparator()
	me.boxMenu.AddAction(me.boxAddFromSearchAction)
	me.boxMenu.AddAction(me.boxAddFromBoxAction)
	me.boxMenu.AddSeparator()
	me.boxMenu.AddAction(me.boxDeleteAction)
	me.searchMenu = menubar.AddMenuWithTitle("&Search")
	// TODO   &New F9
	// TODO   &Delete
	me.viewMenu = menubar.AddMenuWithTitle("&View")
	// TODO   &Visible Cards
	// TODO   &Unboxed Cards
	// TODO   &Hidden Cards
	// TODO   &Card… # choice of cards
	// TODO   &Box… # choice of boxes
	// TODO   &Search… # choice of searches
	me.windowMenu = menubar.AddMenuWithTitle("&Window")
	// TODO   &Next Ctrl+Tab
	// TODO   &Previous Ctrl+Shift+Tab
	// TODO   &Cascade
	// TODO   &Tile
	// TODO   &Windows → 1. | 2. | … | 9. | A. | … | Z.
	// TODO   &Close Ctrl+W
	me.helpMenu = menubar.AddMenuWithTitle("&Help")
	// TODO  &Help F1
	// TODO  &About # VERSION & miQt version & Qt version & SqliteVersion()
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
	const CARD = "Card"
	cardToolbar := me.window.AddToolBarWithTitle(CARD)
	cardToolbar.SetObjectName(*qt.NewQAnyStringView3(CARD))
	cardToolbar.AddAction(me.cardNewAction)
	cardToolbar.AddSeparator()
	cardToolbar.AddAction(me.cardAddToBoxAction)
	cardToolbar.AddAction(me.cardRemoveFromBoxAction)
	cardToolbar.AddSeparator()
	cardToolbar.AddAction(me.cardExportAction)
	cardToolbar.AddSeparator()
	cardToolbar.AddAction(me.cardUnhideAction)
	cardToolbar.AddAction(me.cardHideAction)
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
	me.fileNewAction.OnTriggered(func() { me.fileNew() })
	me.fileOpenAction.OnTriggered(func() { me.fileOpen() })
	me.fileSaveAction.OnTriggered(func() { me.fileSave() })
	me.fileSaveAsAction.OnTriggered(func() { me.fileSaveAs() })
	me.fileExportAction.OnTriggered(func() { me.fileExport() })
	me.fileConfigureAction.OnTriggered(func() { me.fileConfigure() })
	me.fileQuitAction.OnTriggered(func() { me.window.Close() })
	me.editCopyAction.OnTriggered(func() { me.editCopy() })
	me.editCutAction.OnTriggered(func() { me.editCut() })
	me.editPasteAction.OnTriggered(func() { me.editPaste() })
	me.cardNewAction.OnTriggered(func() { me.cardNew() })
	me.cardAddToBoxAction.OnTriggered(func() { me.cardAddToBox() })
	me.cardRemoveFromBoxAction.OnTriggered(
		func() { me.cardRemoveFromBox() })
	me.cardExportAction.OnTriggered(func() { me.cardExport() })
	me.cardUnhideAction.OnTriggered(func() { me.cardUnhide() })
	me.cardHideAction.OnTriggered(func() { me.cardHide() })
	me.cardDeleteAction.OnTriggered(func() { me.cardDelete() })
	me.window.OnCloseEvent(func(super func(event *qt.QCloseEvent),
		event *qt.QCloseEvent,
	) {
		me.SaveSettings()
	})
}
