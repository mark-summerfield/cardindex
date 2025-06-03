// Copyright © 2025 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/mappu/miqt/qt"
	"github.com/mark-summerfield/cardindex/model"
	"github.com/mark-summerfield/ufile"
)

func (me *App) db(msg string, model *model.Model) {
	extra := "(model is "
	if model != nil {
		extra += model.Filename() + ")"
	} else {
		extra += "nil)"
	}
	fmt.Println("DB", msg, extra, me.inFileNew)
}

func (me *App) fileMenuUpdate() {
	me.db("fileMenuUpdate A", me.model)
	// ACCELS must not conflict with menu's Alt+{ACENOSQ}
	ACCELS := []rune("123456789BDFGHIJKLMPRTUVWXYZ")
	me.fileMenu.Clear()
	me.addFileActions()
	me.makeFileConnections()
	if me.config != nil {
		if files := me.config.RecentFiles.Files(); len(files) > 0 {
			me.fileMenu.AddSeparator()
			var action *qt.QAction
			for i, filename := range files {
				text := filepath.Base(filename)
				if i < len(ACCELS) {
					text = fmt.Sprintf("&%c %s", ACCELS[i], text)
				}
				action = qt.NewQAction3(getIcon(SVG_FILE_OPEN), text)
				action.SetToolTip("Open " + filename)
				action.OnTriggered(func() { me.openModel(filename) })
				me.fileMenu.QWidget.AddAction(action)
			}
		}
	}
	me.db("fileMenuUpdate B", me.model)
}

func (me *App) addFileActions() {
	me.fileMenu.QWidget.AddAction(me.fileNewAction)
	me.fileMenu.QWidget.AddAction(me.fileOpenAction)
	me.fileMenu.QWidget.AddAction(me.fileSaveAction)
	me.fileMenu.QWidget.AddAction(me.fileSaveAsAction)
	me.fileMenu.QWidget.AddAction(me.fileExportAction)
	me.fileMenu.AddSeparator()
	me.fileMenu.QWidget.AddAction(me.fileConfigureAction)
	me.fileMenu.AddSeparator()
	me.fileMenu.QWidget.AddAction(me.fileQuitAction)
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
	if me.inFileNew {
		return
	}
	me.inFileNew = true
	defer func() { me.inFileNew = false }()
	me.db("fileNew A", me.model)
	dirname := me.getDefaultDir()
	if filename := qt.QFileDialog_GetSaveFileName3(me.window.QWidget,
		"Create Card Index — "+APPNAME, dirname); filename != "" {
		me.openModel(filename)
	}
	me.db("fileNew B", me.model)
}

func (me *App) fileOpen() {
	dirname := me.getDefaultDir()
	if filename := qt.QFileDialog_GetOpenFileName3(me.window.QWidget,
		"Open Card Index — "+APPNAME, dirname); filename != "" {
		me.openModel(filename)
	} else {
		me.window.SetWindowTitle(APPNAME)
		me.StatusMessage("Click File→New or File→Open", TIMEOUT_LONG)
	}
}

func (me *App) fileSave() {
	// save any windows with unsaved changes
	me.saveMdiWindowsToModel()
	fmt.Println("fileSave") // TODO
}

func (me *App) fileSaveAs() {
	//   Builtin Dialog: choose save filename
	//	- save unsaved changes
	//	- copy .cix to new name
	//	- close model
	//	- open model using new name:
	// me.openModel(filename)
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

func (me *App) openModel(filename string) {
	me.db("openModel A "+filename+"!", me.model)
	me.closeModel()
	if me.mdiArea != nil {
		me.mdiArea.CloseAllSubWindows()
	}
	if !strings.HasSuffix(filename, ".cix") {
		filename += ".cix"
	}
	action := "Opened"
	if !ufile.FileExists(filename) {
		action = "Created"
	}
	if model, err := model.NewModel(filename); err == nil {
		me.model = model
		me.config.RecentFiles.Add(filename)
		me.readMdiWindowsFromModel()
		me.StatusMessage(action+" “"+filename+"”", TIMEOUT_LONG)
		if counts, err := me.model.CardCounts(); err == nil {
			me.StatusIndicatorUpdate(counts.Visible, counts.Unboxed)
		} else {
			me.onError(fmt.Sprintf("Failed to read card counts:\n%s", err))
		}
		me.window.SetWindowTitle(filepath.Base(filename) + " — " + APPNAME)
		me.statusIndicator.QWidget.SetToolTip(filename)
	} else {
		me.onError(fmt.Sprintf("Failed to open %s:\n%s", filename, err))
	}
	me.fileMenuUpdate()
	me.db("openModel B", me.model)
}

func (me *App) closeModel() {
	me.db("closeModel A", me.model)
	me.window.SetWindowTitle(APPNAME)
	me.StatusIndicatorUpdate(0, 0)
	me.statusIndicator.QWidget.SetToolTip("")
	if me.model != nil {
		me.saveMdiWindowsToModel()
		filename := me.model.Filename()
		if err := me.model.Close(); err != nil {
			me.onError(fmt.Sprintf("Error closing %s:\n%s", filename, err))
		}
		me.model = nil
	}
	me.db("closeModel B", me.model)
}

func (me *App) saveMdiWindowsToModel() {
	// TODO save all MDI window states to me.model CONFIG table
	fmt.Println("saveMdiWindowsToModel")
}

func (me *App) readMdiWindowsFromModel() {
	// TODO load all MDI window states from me.model CONFIG table &
	// size & position MDI windows accordingly
	fmt.Println("readMdiWindowsFromModel")
}
