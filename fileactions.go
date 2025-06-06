// Copyright © 2025 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/mappu/miqt/qt"
	"github.com/mark-summerfield/cardindex/database"
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
		me.fileOpenDatabase(filename)
	}
}

func (me *App) fileOpen() {
	dirname := me.getDefaultDir()
	if filename := qt.QFileDialog_GetOpenFileName4(me.window.QWidget,
		"Open Card Index — "+APPNAME, dirname,
		FILE_FILTER); filename != "" {
		me.fileOpenDatabase(filename)
	} else {
		me.window.SetWindowTitle(APPNAME)
		me.StatusMessage("Click File→New or File→Open", TIMEOUT_LONG)
		me.updateUi()
	}
}

func (me *App) fileSave() {
	// save any windows with unsaved changes
	me.fileSaveMdiWindowsToDatabase()
	fmt.Println("fileSave") // TODO
}

func (me *App) fileSaveAs() {
	//   Builtin Dialog: choose save filename
	//	- save unsaved changes
	//	- copy .cix to new name
	//	- close model
	//	- open model using new name:
	// me.fileOpenDatabase(filename)
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

func (me *App) fileOpenDatabase(filename string) {
	me.fileCloseDatabase()
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
	if db, err := database.NewDatabase(filename); err == nil {
		me.db = db
		me.config.RecentFiles.Add(filename)
		me.fileReadMdiWindowsFromDatabase()
		me.StatusMessage(action+" “"+filename+"”", TIMEOUT_LONG)
		me.window.SetWindowTitle(filepath.Base(filename) + " — " + APPNAME)
		me.statusIndicator.QWidget.SetToolTip(filename)
	} else {
		me.onError(fmt.Sprintf("Failed to open %s:\n%s", filename, err))
	}
	me.updateUi()
}

func (me *App) fileCloseDatabase() {
	me.window.SetWindowTitle(APPNAME)
	me.statusIndicator.QWidget.SetToolTip("")
	if me.db != nil {
		me.fileSaveMdiWindowsToDatabase()
		filename := me.db.Filename()
		if err := me.db.Close(); err != nil {
			me.onError(fmt.Sprintf("Error closing %s:\n%s", filename, err))
		}
		me.db = nil
	}
}

func (me *App) fileSaveMdiWindowsToDatabase() {
	// TODO save all MDI window states to me.db CONFIG table
	fmt.Println("fileSaveMdiWindowsToDatabase")
}

func (me *App) fileReadMdiWindowsFromDatabase() {
	// TODO load all MDI window states from me.db CONFIG table &
	// size & position MDI windows accordingly
	fmt.Println("fileReadMdiWindowsFromDatabase")
}
