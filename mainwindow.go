// Copyright © 2025 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"github.com/mappu/miqt/qt"
)

func (me *App) MakeMainWindow() {
	me.MakeActions()
	me.MakeMainMenu()
	me.MakeToolbars()
	me.MakeWidgets()
	me.MakeConnections()
}

func (me *App) MakeWidgets() {
	me.mdiArea = qt.NewQMdiArea2()
	me.makeStatusBar()
	me.window.SetCentralWidget(me.mdiArea.QWidget)
}

func (me *App) makeStatusBar() {
	me.statusIndicator = qt.NewQLabel3("")
	me.statusIndicator.SetFrameShadow(qt.QFrame__Sunken)
	me.statusIndicator.SetFrameShape(qt.QFrame__StyledPanel)
	statusbar := me.window.StatusBar()
	statusbar.SetSizeGripEnabled(false)
	statusbar.AddPermanentWidget(me.statusIndicator.QWidget)
	me.StatusMessage("Click File→New or File→Open", TIMEOUT_LONG)
}
