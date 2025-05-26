// Copyright © 2025 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"log"

	qt "github.com/mappu/miqt/qt6"
	"github.com/mark-summerfield/ufile"
)

type App struct {
	config           *Config
	model            *Model
	window           *qt.QMainWindow
	fileMenu         *qt.QMenu
	fileNewAction    *qt.QAction
	fileOpenAction   *qt.QAction
	fileSaveAction   *qt.QAction
	fileSaveAsAction *qt.QAction
	fileQuitAction   *qt.QAction
}

func NewApp() *App {
	var app App
	app.window = qt.NewQMainWindow(nil)
	app.window.SetWindowIcon(getIcon(SVG_ICON))
	app.LoadSettings()
	app.MakeActions()
	app.MakeMainMenu()
	app.MakeToolbars()
	app.MakeWidgets()
	app.MakeLayout()
	app.MakeConnections()
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

func (me *App) MakeActions() {
	me.fileNewAction = qt.NewQAction3(getIcon(SVG_FILE_NEW), "&New…")
	me.fileNewAction.SetShortcutsWithShortcuts(qt.QKeySequence__New)
	me.fileOpenAction = qt.NewQAction3(getIcon(SVG_FILE_OPEN), "&Open…")
	me.fileOpenAction.SetShortcutsWithShortcuts(qt.QKeySequence__Open)
	me.fileSaveAction = qt.NewQAction3(getIcon(SVG_FILE_SAVE), "&Save")
	me.fileSaveAction.SetShortcutsWithShortcuts(qt.QKeySequence__Save)
	me.fileSaveAsAction = qt.NewQAction3(getIcon(SVG_FILE_SAVE_AS),
		"Save &As…")
	me.fileQuitAction = qt.NewQAction3(getIcon(SVG_FILE_QUIT), "&Quit")
	me.fileQuitAction.SetShortcutsWithShortcuts(qt.QKeySequence__Quit)
	// TODO
}

func (me *App) MakeMainMenu() {
	menubar := me.window.MenuBar()
	me.fileMenu = menubar.AddMenuWithTitle("&File")
	me.fileMenu.AddAction(me.fileNewAction)
	me.fileMenu.AddAction(me.fileOpenAction)
	me.fileMenu.AddAction(me.fileSaveAction)
	me.fileMenu.AddAction(me.fileSaveAsAction)
	me.fileMenu.AddSeparator()
	me.fileMenu.AddAction(me.fileQuitAction)
	// TODO
}

func (me *App) MakeToolbars() {
	const FILE = "File"
	fileToolbar := me.window.AddToolBarWithTitle(FILE)
	fileToolbar.SetObjectName(*qt.NewQAnyStringView3(FILE))
	fileToolbar.AddAction(me.fileNewAction)
	fileToolbar.AddAction(me.fileOpenAction)
	fileToolbar.AddAction(me.fileSaveAction)
	// TODO
}

func (me *App) MakeWidgets() {
	// TODO
}

func (me *App) MakeLayout() {
	// TODO
	// widget := qt.NewQWidget(me.window.QWidget)
	// box := qt.NewQHBoxLayout(widget)
	// box.AddWidget(me.helloButton.QWidget)
	// box.AddWidget(me.aboutButton.QWidget)
	// box.AddWidget(me.quitButton.QWidget)
	// me.window.SetCentralWidget(widget)
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
