// Copyright © 2025 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"log"

	qt "github.com/mappu/miqt/qt6"
	"github.com/mark-summerfield/ufile"
)

type App struct {
	config              *Config
	model               *Model
	window              *qt.QMainWindow
	mdiArea             *qt.QMdiArea
	statusIndicator     *qt.QLabel
	fileMenu            *qt.QMenu
	fileNewAction       *qt.QAction
	fileOpenAction      *qt.QAction
	fileSaveAction      *qt.QAction
	fileSaveAsAction    *qt.QAction
	fileExportAction    *qt.QAction
	fileConfigureAction *qt.QAction
	fileQuitAction      *qt.QAction
}

func NewApp() *App {
	var app App
	app.window = qt.NewQMainWindow(nil)
	app.window.SetWindowIcon(getIcon(SVG_ICON))
	app.LoadSettings()
	app.MakeMainWindow()
	if app.model != nil {
		app.LoadModel()
	}
	return &app
}

func (me *App) Show() { me.window.Show() }

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
	if err := me.config.Save(); err != nil {
		log.Printf("failed to save config in %q: %v\n", me.config.Filename,
			err)
	}
}
