// Copyright © 2025 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/mappu/miqt/qt"
	"github.com/mark-summerfield/cardindex/model"
	"github.com/mark-summerfield/ufile"
)

func (me *App) fileMenuUpdate() {
	files := me.config.RecentFiles.Files()
	for i, action := range me.fileOpenActions {
		if i < len(files) {
			filename := files[i]
			text := filepath.Base(filename)
			if i <= MAX_RECENT_FILES {
				text = "&" + strconv.Itoa(i+1) + " " + text
			}
			action.SetText(text)
			action.SetToolTip(filename)
			action.SetVisible(true)
		} else {
			action.SetText("")
			action.SetToolTip("")
			action.SetVisible(false)
		}
	}
}

func (me *App) fileNew() {
	dirname := me.getDefaultDir()
	if filename := qt.QFileDialog_GetSaveFileName4(me.window.QWidget,
		"Create Card Index — "+APPNAME, dirname,
		FILE_FILTER); filename != "" {
		me.fileOpenModel(filename)
	}
}

func (me *App) fileOpen() {
	dirname := me.getDefaultDir()
	if filename := qt.QFileDialog_GetOpenFileName4(me.window.QWidget,
		"Open Card Index — "+APPNAME, dirname,
		FILE_FILTER); filename != "" {
		me.fileOpenModel(filename)
	} else {
		me.window.SetWindowTitle(APPNAME)
		me.StatusMessage("Click File→New or File→Open", TIMEOUT_LONG)
		me.updateUi()
	}
}

func (me *App) fileSave() {
	// save any windows with unsaved changes
	me.fileSaveMdiWindowsToModel()
	fmt.Println("fileSave") // TODO
}

func (me *App) fileSaveAs() {
	//   Builtin Dialog: choose save filename
	//	- save unsaved changes
	//	- copy .cix to new name
	//	- close model
	//	- open model using new name:
	// me.fileOpenModel(filename)
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
	me.updateUi()
}

func (me *App) fileOpenModel(filename string) {
	me.fileCloseModel()
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
		me.fileReadMdiWindowsFromModel()
		me.StatusMessage(action+" “"+filename+"”", TIMEOUT_LONG)
		me.window.SetWindowTitle(filepath.Base(filename) + " — " + APPNAME)
		me.statusIndicator.QWidget.SetToolTip(filename)
	} else {
		me.onError(fmt.Sprintf("Failed to open %s:\n%s", filename, err))
	}
	me.updateUi()
}

func (me *App) fileCloseModel() {
	me.window.SetWindowTitle(APPNAME)
	me.statusIndicator.QWidget.SetToolTip("")
	if me.model != nil {
		me.fileSaveMdiWindowsToModel()
		filename := me.model.Filename()
		if err := me.model.Close(); err != nil {
			me.onError(fmt.Sprintf("Error closing %s:\n%s", filename, err))
		}
		me.model = nil
	}
}

func (me *App) fileSaveMdiWindowsToModel() {
	// TODO save all MDI window states to me.model CONFIG table
	fmt.Println("fileSaveMdiWindowsToModel")
}

func (me *App) fileReadMdiWindowsFromModel() {
	// TODO load all MDI window states from me.model CONFIG table &
	// size & position MDI windows accordingly
	fmt.Println("fileReadMdiWindowsFromModel")
}
