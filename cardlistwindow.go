// Copyright © 2025 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"fmt"

	"github.com/mappu/miqt/qt"
	"github.com/mark-summerfield/cardindex/database"
)

type CardListWindow struct {
	db      *database.Database
	onError OnError
	window  *qt.QMdiSubWindow
	kind    CardListKind
	oid     database.Oid
	// TODO widgets for makeListWindow & Refresh
}

func NewListWindow(db *database.Database, onError OnError,
	kind CardListKind,
) *CardListWindow {
	switch kind {
	case CARD_LIST_VISIBLE_KIND, CARD_LIST_UNBOXED_KIND,
		CARD_LIST_HIDDEN_KIND: // OK
	default:
		panic("internal error: invalid list kind: " + kind.String())
	}
	window := qt.NewQMdiSubWindow2()
	window.SetWindowTitle(kind.String() + " Cards")
	window.SetWindowIcon(getIcon(kind.IconName()))
	window.SetWidget(makeListWindow())
	return &CardListWindow{
		db: db, onError: onError, window: window, kind: kind,
		oid: database.NewOid(database.NAME),
	}
}

func makeListWindow() *qt.QWidget {
	widget := qt.NewQWidget2()
	orderByGroupBox := qt.NewQGroupBox3("&Order By")
	nameRadioButton := qt.NewQRadioButton3("&" + database.NAME)
	nameRadioButton.SetChecked(true)
	updatedRadioButton := qt.NewQRadioButton3("&" + database.UPDATED)
	createdRadioButton := qt.NewQRadioButton3("&" + database.CREATED)
	vbox := qt.NewQVBoxLayout2()
	hbox := qt.NewQHBoxLayout2()
	hbox.AddWidget(nameRadioButton.QWidget)
	hbox.AddWidget(updatedRadioButton.QWidget)
	hbox.AddWidget(createdRadioButton.QWidget)
	vbox.AddWidget(orderByGroupBox.QWidget)
	widget.SetLayout(vbox.QLayout) // TODO change to whole layout
	// TODO make connections
	return widget
}

// 'l' List
// Name (of List)        (X)
// Order by [   Name v] # Updated | Created
// ... use a vertical splitter
// Card Name1
// Card Name2
//     :
// 0 cards in list
// # Context Menu: (none)

func (me *CardListWindow) Refresh(oid database.Oid) {
	var cardNames []database.CardName
	var err error
	switch me.kind {
	case CARD_LIST_VISIBLE_KIND:
		cardNames, err = me.db.CardNamesVisible(oid)
	case CARD_LIST_UNBOXED_KIND:
		cardNames, err = me.db.CardNamesUnboxed(oid)
	case CARD_LIST_HIDDEN_KIND:
		cardNames, err = me.db.CardNamesHidden(oid)
	}
	if err != nil {
		me.onError(fmt.Sprintf("failed to show %s cards: %s", me.kind, err))
		return
	}
	// TODO clear widgets & repopulate them
	fmt.Println("CardListWindow.Update", cardNames)
}
