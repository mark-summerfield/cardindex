// Copyright © 2025 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"fmt"
	"path/filepath"

	qt "github.com/mappu/miqt/qt6"
)

func (me *App) fileMenuUpdate() {
	me.fileMenu.Clear()
	me.addFileActions()
	me.makeFileConnections()
	if files := me.config.RecentFiles.Files(); len(files) > 0 {
		me.fileMenu.AddSeparator()
		var action *qt.QAction
		for i, filename := range files {
			text := filepath.Base(filename)
			if i < 9 {
				text = fmt.Sprintf("&%d %s", i+1, text)
			}
			action = qt.NewQAction3(getIcon(SVG_FILE_OPEN), text)
			action.SetToolTip("Open " + filename)
			action.OnTriggered(func() { me.fileOpenRecent(filename) })
			me.fileMenu.AddAction(action)
		}
	}
}

func (me *App) addFileActions() {
	me.fileMenu.AddAction(me.fileNewAction)
	me.fileMenu.AddAction(me.fileOpenAction)
	me.fileMenu.AddAction(me.fileSaveAction)
	me.fileMenu.AddAction(me.fileSaveAsAction)
	me.fileMenu.AddAction(me.fileExportAction)
	me.fileMenu.AddSeparator()
	me.fileMenu.AddAction(me.fileConfigureAction)
	me.fileMenu.AddSeparator()
	me.fileMenu.AddAction(me.fileQuitAction)
}

func (me *App) makeFileConnections() {
	me.fileNewAction.OnTriggered(func() { me.fileNew() })
	me.fileOpenAction.OnTriggered(func() { me.fileOpen() })
	me.fileSaveAction.OnTriggered(func() { me.fileSave() })
	me.fileSaveAsAction.OnTriggered(func() { me.fileSaveAs() })
	me.fileExportAction.OnTriggered(func() { me.fileExport() })
	me.fileConfigureAction.OnTriggered(func() { me.fileConfigure() })
	me.fileQuitAction.OnTriggered(func() { me.window.Close() })
}

func (me *App) fileNew() {
	//   Builtin Dialog: choose nonexistent filename
	// me.loadModel(filename)
	fmt.Println("fileNew") // TODO
}

func (me *App) fileOpen() {
	//   Builtin Dialog: choose existing filename
	// me.loadModel(filename)
	fmt.Println("fileOpen") // TODO
}

func (me *App) fileOpenRecent(filename string) {
	me.loadModel(filename)
}

func (me *App) fileSave() {
	// save any windows with unsaved changes
	fmt.Println("fileSave") // TODO
}

func (me *App) fileSaveAs() {
	//   Builtin Dialog: choose save filename
	//	- save unsaved changes
	//	- copy .cix to new name
	//	- close model
	//	- open model using new name:
	// me.loadModel(filename)
	fmt.Println("fileSaveAs") // TODO
}

func (me *App) fileExport() {
	//   Dialog
	//		All | Box (choice) | Search (choice) | Card (choice)
	//		HTML | CommonMark | ODF | PDF
	//      One card per page | One box per page | Continuous
	fmt.Println("fileExport") // TODO
}

func (me *App) fileConfigure() {
	//   Dialog
	//		+-- Show ---------------------------------+
	//		| [X] File Toolbar [X] Edit Toolbar …     |
	//		+-----------------------------------------+
	//		[X] Cursor Blink # set config.CursorBlink directly)
	fmt.Println("fileConfigure") // TODO
}

func (me *App) loadModel(filename string) {
	// - save & close any existing model
	// - close all MDI windows
	// - create new model with this file
	// - update window title to filename — APPNAME
	// - create & size & position MDI windows as per new model's UI config
	fmt.Println("loadModel", filename) // TODO
	if filename != "" {
		me.config.RecentFiles.Add(filename)
	}
}
