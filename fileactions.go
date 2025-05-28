// Copyright © 2025 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"fmt"
)

func (me *App) fileNew() {
	//   Builtin Dialog: choose nonexistent filename
	// loadModel(filename)
	fmt.Println("fileNew") // TODO
}

func (me *App) fileOpen() {
	//   Builtin Dialog: choose existing filename
	// loadModel(filename)
	fmt.Println("fileOpen") // TODO
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
	//	- open model using new name
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
}
