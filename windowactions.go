// Copyright © 2025 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"fmt"

	"github.com/mappu/miqt/qt"
)

func (me *App) windowMenuUpdate() {
	me.windowMenu.Clear()
	me.addWindowActions()
	me.makeWindowConnections()
	if me.mdiArea == nil {
		return
	}
	// ACCELS must not conflict with menu's Alt+{ACNPT}
	ACCELS := []rune("123456789BDEFGHIJKLMOQRSUVWXYZ")
	if children := me.mdiArea.SubWindowList(); len(children) > 0 {
		me.windowMenu.AddSeparator()
		var action *qt.QAction
		for i, child := range children {
			text := child.WindowTitle()
			if i < len(ACCELS) {
				text = fmt.Sprintf("&%c %s", ACCELS[i], text)
			}
			action = qt.NewQAction3(getIcon(SVG_WINDOW), text)
			action.SetToolTip("Activate window " + text)
			action.OnTriggered( // TODO check this works correctly
				func() { me.mdiArea.SetActiveSubWindow(child) })
			me.windowMenu.QWidget.AddAction(action)
		}
	}
}

func (me *App) addWindowActions() {
	me.windowMenu.QWidget.AddAction(me.windowNextAction)
	me.windowMenu.QWidget.AddAction(me.windowPrevAction)
	me.windowMenu.QWidget.AddAction(me.windowCascadeAction)
	me.windowMenu.QWidget.AddAction(me.windowTileAction)
	me.windowMenu.AddSeparator()
	me.windowMenu.QWidget.AddAction(me.windowCloseAction)
}

func (me *App) makeWindowConnections() {
	me.windowNextAction.OnTriggered(
		func() { me.mdiArea.ActivateNextSubWindow() })
	me.windowPrevAction.OnTriggered(
		func() { me.mdiArea.ActivatePreviousSubWindow() })
	me.windowCascadeAction.OnTriggered(
		func() { me.mdiArea.CascadeSubWindows() })
	me.windowTileAction.OnTriggered(
		func() { me.mdiArea.TileSubWindows() })
	me.windowCloseAction.OnTriggered(
		func() { me.mdiArea.CloseActiveSubWindow() })
}
