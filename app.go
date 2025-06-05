// Copyright © 2025 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"log"
	"path/filepath"

	"github.com/mappu/miqt/qt"
	"github.com/mark-summerfield/cardindex/model"
	"github.com/mark-summerfield/ufile"
)

type App struct {
	config                   *Config
	model                    *model.Model
	window                   *qt.QMainWindow
	mdiArea                  *qt.QMdiArea
	statusIndicator          *qt.QLabel
	fileMenu                 *qt.QMenu
	fileNewAction            *qt.QAction
	fileOpenAction           *qt.QAction
	fileSaveAction           *qt.QAction
	fileSaveAsAction         *qt.QAction
	fileExportAction         *qt.QAction
	fileConfigureAction      *qt.QAction
	fileQuitAction           *qt.QAction
	fileOpenActions          []*qt.QAction
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
	boxMenu                  *qt.QMenu
	boxNewAction             *qt.QAction
	boxViewAction            *qt.QAction
	boxAddFromSearchAction   *qt.QAction
	boxAddFromBoxAction      *qt.QAction
	boxDeleteAction          *qt.QAction
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
		me.openModel(me.config.MostRecentFile)
	}
	me.fileMenuUpdate()
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
	if me.model != nil {
		me.config.MostRecentFile = me.model.Filename()
	}
	if err := me.config.Save(); err != nil {
		log.Printf("failed to save config in %q: %v\n", me.config.Filename,
			err)
	}
}

func (me *App) onError(message string) {
	qt.QMessageBox_Warning3(me.window.QWidget, "Error — "+APPNAME, message,
		"&Close")
}

func (me *App) getDefaultDir() string {
	var dirname string
	if me.model != nil {
		dirname = filepath.Dir(me.model.Filename())
	}
	if dirname == "" {
		dirname = ufile.HomeDir()
	}
	return dirname
}
