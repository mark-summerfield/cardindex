// Copyright © 2025 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

func (me *App) windowMenuUpdate() {
	me.windowMenu.Clear()
	me.addWindowActions()
	me.makeWindowConnections()
	// Add windows
	//if files := me.config.RecentFiles.Files(); len(files) > 0 {
	//	me.fileMenu.AddSeparator()
	//	var action *qt.QAction
	//	for i, filename := range files {
	//		text := filepath.Base(filename)
	//		if i < 9 {
	//			text = fmt.Sprintf("&%d %s", i+1, text)
	//		}
	//		action = qt.NewQAction3(getIcon(SVG_FILE_OPEN), text)
	//		action.SetToolTip("Open " + filename)
	//		action.OnTriggered(func() { me.fileOpenRecent(filename) })
	//		me.fileMenu.QWidget.AddAction(action)
	//	}
	//}
}

func (me *App) addWindowActions() {
	// me.fileMenu.QWidget.AddAction(me.fileNewAction)
	// me.fileMenu.QWidget.AddAction(me.fileOpenAction)
	// me.fileMenu.QWidget.AddAction(me.fileSaveAction)
	// me.fileMenu.QWidget.AddAction(me.fileSaveAsAction)
	// me.fileMenu.QWidget.AddAction(me.fileExportAction)
	// me.fileMenu.AddSeparator()
	// me.fileMenu.QWidget.AddAction(me.fileConfigureAction)
	// me.fileMenu.AddSeparator()
	// me.fileMenu.QWidget.AddAction(me.fileQuitAction)
}

func (me *App) makeWindowConnections() {
	// me.fileNewAction.OnTriggered(func() { me.fileNew() })
	// me.fileOpenAction.OnTriggered(func() { me.fileOpen() })
	// me.fileSaveAction.OnTriggered(func() { me.fileSave() })
	// me.fileSaveAsAction.OnTriggered(func() { me.fileSaveAs() })
	// me.fileExportAction.OnTriggered(func() { me.fileExport() })
	// me.fileConfigureAction.OnTriggered(func() { me.fileConfigure() })
	// me.fileQuitAction.OnTriggered(func() { me.window.Close() })
}

// func (me *App) fileNew() {
// 	//   Builtin Dialog: choose nonexistent filename
// 	// me.loadModel(filename)
// 	fmt.Println("fileNew") // TODO
// }
