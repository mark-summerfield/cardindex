// Copyright © 2025 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"fmt"

	qt "github.com/mappu/miqt/qt6"
)

type App struct {
	window *qt.QMainWindow
}

func NewApp() *App {
	var app App
	app.window = qt.NewQMainWindow(nil)
	app.window.SetWindowIcon(getIcon(Icon))
	app.ReadSettings()
	app.MakeWidgets()
	app.MakeLayout()
	app.MakeConnections()
	return &app
}

func (me *App) ReadSettings() {
	qt.QCoreApplication_SetApplicationName(APPNAME)
	settings := qt.NewQSettings7(ORGANIZATION, APPNAME)
	if cursorblink := settings.Value(*OPT_CURSOR_BLINK,
		DEF_CURSOR_BLINK).ToBool(); !cursorblink {
		qt.QApplication_SetCursorFlashTime(0)
	}
	// TODO window size & pos
}

func (me *App) MakeWidgets() {
	//
}

func (me *App) MakeLayout() {
	// widget := qt.NewQWidget(me.window.QWidget)
	// box := qt.NewQHBoxLayout(widget)
	// box.AddWidget(me.helloButton.QWidget)
	// box.AddWidget(me.aboutButton.QWidget)
	// box.AddWidget(me.quitButton.QWidget)
	// me.window.SetCentralWidget(widget)
}

func (me *App) MakeConnections() {
	// me.quitButton.OnPressed(func() { me.window.Close() })
	me.window.OnCloseEvent(func(super func(event *qt.QCloseEvent),
		event *qt.QCloseEvent,
	) {
		me.SaveState()
	})
}

func (me *App) Show() { me.window.Show() }

func (me *App) SaveState() {
	fmt.Println("CardIndex: save state")
}

func getIcon(data []byte) *qt.QIcon {
	image := qt.QImage_FromDataWithData(data)
	pixmap := qt.QPixmap_FromImage(image)
	return qt.NewQIcon2(pixmap)
}
