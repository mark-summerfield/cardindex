// Copyright © 2025 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"fmt"

	"github.com/mark-summerfield/unum"
)

func (me *App) StatusMessage(message string, timeout int) {
	me.window.StatusBar().ShowMessage2(message, timeout)
}

func (me *App) StatusClear() { me.window.StatusBar().ClearMessage() }

func (me *App) StatusIndicatorUpdate(cards, unboxed int) {
	me.statusIndicator.SetText(fmt.Sprintf("%s Cards • %s Unboxed",
		unum.Commas(cards), unum.Commas(unboxed)))
}
