// Copyright © 2025 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/mappu/miqt/qt"
	"github.com/mark-summerfield/cardindex/database"
	"github.com/mark-summerfield/ufile"
)

type App struct {
	config                   *Config
	db                       *database.Database
	window                   *qt.QMainWindow
	mdiArea                  *qt.QMdiArea
	statusIndicator          *qt.QLabel
	fileToolbar              *qt.QToolBar
	fileMenu                 *qt.QMenu
	fileNewAction            *qt.QAction
	fileOpenAction           *qt.QAction
	fileSaveAction           *qt.QAction
	fileSaveAsAction         *qt.QAction
	fileExportAction         *qt.QAction
	fileConfigureAction      *qt.QAction
	fileQuitAction           *qt.QAction
	fileOpenActions          []*qt.QAction
	edit1Toolbar             *qt.QToolBar
	edit2Toolbar             *qt.QToolBar
	edit3Toolbar             *qt.QToolBar
	edit4Toolbar             *qt.QToolBar
	edit5Toolbar             *qt.QToolBar
	editMenu                 *qt.QMenu
	editUndoAction           *qt.QAction
	editRedoAction           *qt.QAction
	editCopyAction           *qt.QAction
	editCutAction            *qt.QAction
	editPasteAction          *qt.QAction
	editBoldAction           *qt.QAction
	editItalicAction         *qt.QAction
	editMonospaceAction      *qt.QAction
	editBulletListAction     *qt.QAction
	editNumberedListAction   *qt.QAction
	editEndListAction        *qt.QAction
	editInsertWebLinkAction  *qt.QAction
	editInsertFileLinkAction *qt.QAction
	editInsertSymbolAction   *qt.QAction
	cardToolbar              *qt.QToolBar
	cardMenu                 *qt.QMenu
	cardNewAction            *qt.QAction
	cardViewVisibleAction    *qt.QAction
	cardViewUnboxedAction    *qt.QAction
	cardViewHiddenAction     *qt.QAction
	cardAddToBoxAction       *qt.QAction
	cardRemoveFromBoxAction  *qt.QAction
	cardExportAction         *qt.QAction
	cardUnhideAction         *qt.QAction
	cardHideAction           *qt.QAction
	cardDeleteAction         *qt.QAction
	boxToolbar               *qt.QToolBar
	boxMenu                  *qt.QMenu
	boxNewAction             *qt.QAction
	boxViewAction            *qt.QAction
	boxAddFromSearchAction   *qt.QAction
	boxAddFromBoxAction      *qt.QAction
	boxDeleteAction          *qt.QAction
	searchToolbar            *qt.QToolBar
	searchMenu               *qt.QMenu
	searchNewAction          *qt.QAction
	searchViewAction         *qt.QAction
	searchDeleteAction       *qt.QAction
	windowMenu               *qt.QMenu
	windowNextAction         *qt.QAction
	windowPrevAction         *qt.QAction
	windowCascadeAction      *qt.QAction
	windowTileAction         *qt.QAction
	windowCloseAction        *qt.QAction
	helpMenu                 *qt.QMenu
	helpHelpAction           *qt.QAction
	helpAboutAction          *qt.QAction
}

func NewApp() *App {
	var app App
	app.window = qt.NewQMainWindow(nil)
	app.window.SetWindowIcon(getIcon(SVG_ICON))
	app.MakeMainWindow()
	app.LoadSettings()
	return &app
}

func (me *App) Show() {
	me.window.Show()
	if me.config.MostRecentFile != "" {
		me.fileOpenDatabase(me.config.MostRecentFile)
	} else {
		me.updateUi()
	}
}

func (me *App) LoadSettings() {
	filename, exists := ufile.GetIniFile(DOMAIN, APPNAME)
	if exists {
		me.config = NewConfigFrom(filename)
	} else {
		me.config = NewConfig(filename)
	}
	if !me.config.CursorBlink {
		qt.QApplication_SetCursorFlashTime(0)
	}
	me.window.RestoreGeometry(me.config.WindowGeometry)
	me.window.RestoreState(me.config.WindowState)
}

func (me *App) SaveSettings() {
	me.config.WindowGeometry = me.window.SaveGeometry()
	me.config.WindowState = me.window.SaveState()
	if me.db != nil {
		me.config.MostRecentFile = me.db.Filename()
	}
	if err := me.config.Save(); err != nil {
		log.Printf("failed to save config in %q: %v\n", me.config.Filename,
			err)
	}
}

func (me *App) updateUi() {
	me.fileMenuUpdate()
	enable := me.db != nil
	for _, menu := range []*qt.QMenu{
		me.editMenu, me.cardMenu, me.boxMenu, me.searchMenu, me.windowMenu,
	} {
		menu.SetEnabled(enable)
	}
	for _, toolbar := range []*qt.QToolBar{
		me.edit1Toolbar, me.edit2Toolbar, me.edit3Toolbar, me.edit4Toolbar,
		me.edit5Toolbar, me.cardToolbar, me.boxToolbar, me.searchToolbar,
	} {
		toolbar.SetEnabled(enable)
	}
	for _, action := range []*qt.QAction{
		me.fileSaveAction, me.fileSaveAsAction,
		me.fileExportAction,
	} {
		action.SetEnabled(enable)
	}
	if enable {
		if counts, err := me.db.CardCounts(); err == nil {
			me.StatusIndicatorUpdate(counts.Visible, counts.Unboxed, true)
		} else {
			me.onError(fmt.Sprintf("Failed to read card counts:\n%s", err))
		}
		// TODO enable edit actions depending on visible windows etc.
	} else {
		me.StatusIndicatorUpdate(0, 0, false)
		me.statusIndicator.QWidget.SetToolTip("")
		// TODO disable edit actions depending on visible windows etc.
	}
}

func (me *App) onError(message string) {
	qt.QMessageBox_Warning3(me.window.QWidget, "Error — "+APPNAME, message,
		"&Close")
}

func (me *App) getDefaultDir() string {
	var dirname string
	if me.db != nil {
		dirname = filepath.Dir(me.db.Filename())
	}
	if dirname == "" {
		dirname = ufile.HomeDir()
	}
	return dirname
}
