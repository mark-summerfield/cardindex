// Copyright © 2025 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"fmt"
)

func (me *App) cardNew() {
	fmt.Println("cardNew") // TODO
	me.updateUi()
}

func (me *App) cardViewVisible() {
	if me.cardListVisibleWindow == nil {
		me.cardListVisibleWindow = NewListWindow(me.db, me.onError,
			CARD_LIST_VISIBLE_KIND)
	} else {
		me.cardListVisibleWindow.db = me.db // in case it's changed
	}
	// refresh using default or last specified oid
	me.cardListVisibleWindow.Refresh(me.cardListVisibleWindow.oid)
	me.updateUi()
}

func (me *App) cardViewUnboxed() {
	me.updateUi()
	if me.cardListUnboxedWindow == nil {
		me.cardListUnboxedWindow = NewListWindow(me.db, me.onError,
			CARD_LIST_UNBOXED_KIND)
	} else {
		me.cardListUnboxedWindow.db = me.db // in case it's changed
	}
	// refresh using default or last specified oid
	me.cardListUnboxedWindow.Refresh(me.cardListUnboxedWindow.oid)
}

func (me *App) cardViewHidden() {
	if me.cardListHiddenWindow == nil {
		me.cardListHiddenWindow = NewListWindow(me.db, me.onError,
			CARD_LIST_HIDDEN_KIND)
	} else {
		me.cardListHiddenWindow.db = me.db // in case it's changed
	}
	// refresh using default or last specified oid
	me.cardListHiddenWindow.Refresh(me.cardListHiddenWindow.oid)
	me.updateUi()
}

func (me *App) cardAddToBox() {
	fmt.Println("cardAddToBox") // TODO
	me.updateUi()
}

func (me *App) cardRemoveFromBox() {
	fmt.Println("cardRemoveFromBox") // TODO
	me.updateUi()
}

func (me *App) cardExport() {
	//   HTML | CommonMark | ODT | PDF
	fmt.Println("cardExport") // TODO
}

func (me *App) cardUnhide() {
	fmt.Println("cardUnhide") // TODO
	me.updateUi()
}

func (me *App) cardHide() {
	fmt.Println("cardHide") // TODO
	me.updateUi()
}

func (me *App) cardDelete() {
	fmt.Println("cardDelete") // TODO
	me.updateUi()
}
