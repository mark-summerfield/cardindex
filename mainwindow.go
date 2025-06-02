// Copyright © 2025 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"github.com/mappu/miqt/qt"
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
	// TODO
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
	me.fileMenuUpdate()
}

func (me *App) makeMainEditMenu(menubar *qt.QMenuBar) {
	me.editMenu = menubar.AddMenuWithTitle("&Edit")
	// TODO &Undo
	// TODO &Redo
	me.editMenu.QWidget.AddAction(me.editCopyAction)
	me.editMenu.QWidget.AddAction(me.editCutAction)
	me.editMenu.QWidget.AddAction(me.editPasteAction)
	// TODO &Bold
	// TODO &Italic
	// TODO &Monospace
	// TODO &Bullet List
	// TODO &Numbered List
	// TODO &Clear List
	// TODO Insert &Web Link…
	// TODO Insert &File Link…
	// TODO Insert &Symbol…
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
	me.windowMenuUpdate()
}

func (me *App) makeMainHelpMenu(menubar *qt.QMenuBar) {
	me.helpMenu = menubar.AddMenuWithTitle("&Help")
	me.helpMenu.QWidget.AddAction(me.helpHelpAction)
	me.helpMenu.QWidget.AddAction(me.helpAboutAction)
}

func (me *App) MakeToolbars() {
	const FILE = "File"
	fileToolbar := me.window.AddToolBarWithTitle(FILE)
	fileToolbar.SetObjectName(FILE)
	fileToolbar.QWidget.AddAction(me.fileNewAction)
	fileToolbar.QWidget.AddAction(me.fileOpenAction)
	fileToolbar.QWidget.AddAction(me.fileSaveAction)
	const EDIT = "Edit"
	editToolbar := me.window.AddToolBarWithTitle(EDIT)
	editToolbar.SetObjectName(EDIT)
	editToolbar.QWidget.AddAction(me.editCopyAction)
	editToolbar.QWidget.AddAction(me.editCutAction)
	editToolbar.QWidget.AddAction(me.editPasteAction)
	const CARD = "Card"
	cardToolbar := me.window.AddToolBarWithTitle(CARD)
	cardToolbar.SetObjectName(CARD)
	cardToolbar.QWidget.AddAction(me.cardNewAction)
	cardToolbar.QWidget.AddAction(me.cardViewVisibleAction)
	cardToolbar.AddSeparator()
	cardToolbar.QWidget.AddAction(me.cardAddToBoxAction)
	cardToolbar.QWidget.AddAction(me.cardRemoveFromBoxAction)
	cardToolbar.AddSeparator()
	cardToolbar.QWidget.AddAction(me.cardExportAction)
	cardToolbar.AddSeparator()
	cardToolbar.QWidget.AddAction(me.cardUnhideAction)
	cardToolbar.QWidget.AddAction(me.cardHideAction)
	const BOX = "Box"
	boxToolbar := me.window.AddToolBarWithTitle(BOX)
	boxToolbar.SetObjectName(BOX)
	boxToolbar.QWidget.AddAction(me.boxNewAction)
	boxToolbar.QWidget.AddAction(me.boxViewAction)
	boxToolbar.AddSeparator()
	boxToolbar.QWidget.AddAction(me.boxAddFromSearchAction)
	boxToolbar.QWidget.AddAction(me.boxAddFromBoxAction)
	const SEARCH = "Search"
	searchToolbar := me.window.AddToolBarWithTitle(SEARCH)
	searchToolbar.SetObjectName(SEARCH)
	searchToolbar.QWidget.AddAction(me.searchNewAction)
	searchToolbar.QWidget.AddAction(me.searchViewAction)
}

func (me *App) MakeWidgets() {
	me.mdiArea = qt.NewQMdiArea2()
	me.windowMenuUpdate()
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

// For file actions see fileactions.go;
// for window action see windowactions.go
func (me *App) MakeConnections() {
	// TODO
	me.editCopyAction.OnTriggered(func() { me.editCopy() })
	me.editCutAction.OnTriggered(func() { me.editCut() })
	me.editPasteAction.OnTriggered(func() { me.editPaste() })
	// TODO
	me.cardNewAction.OnTriggered(func() { me.cardNew() })
	me.cardViewVisibleAction.OnTriggered(func() { me.cardViewVisible() })
	me.cardViewUnboxedAction.OnTriggered(func() { me.cardViewUnboxed() })
	me.cardViewHiddenAction.OnTriggered(func() { me.cardViewHidden() })
	me.cardAddToBoxAction.OnTriggered(func() { me.cardAddToBox() })
	me.cardRemoveFromBoxAction.OnTriggered(
		func() { me.cardRemoveFromBox() })
	me.cardExportAction.OnTriggered(func() { me.cardExport() })
	me.cardUnhideAction.OnTriggered(func() { me.cardUnhide() })
	me.cardHideAction.OnTriggered(func() { me.cardHide() })
	me.cardDeleteAction.OnTriggered(func() { me.cardDelete() })
	me.boxNewAction.OnTriggered(func() { me.boxNew() })
	me.boxViewAction.OnTriggered(func() { me.boxView() })
	me.boxAddFromSearchAction.OnTriggered(func() { me.boxAddFromSearch() })
	me.boxAddFromBoxAction.OnTriggered(func() { me.boxAddFromBox() })
	me.boxDeleteAction.OnTriggered(func() { me.boxDelete() })
	me.searchNewAction.OnTriggered(func() { me.searchNew() })
	me.searchViewAction.OnTriggered(func() { me.searchView() })
	me.searchDeleteAction.OnTriggered(func() { me.searchDelete() })
	me.helpHelpAction.OnTriggered(func() { me.helpHelp() })
	me.helpAboutAction.OnTriggered(func() { me.helpAbout() })
	me.window.OnCloseEvent(func(super func(event *qt.QCloseEvent),
		event *qt.QCloseEvent,
	) {
		me.SaveSettings()
	})
}
