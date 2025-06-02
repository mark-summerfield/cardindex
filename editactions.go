// Copyright © 2025 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"fmt"
)

// TODO &Undo	editUndoAction // applies to text inside card
// TODO &Redo	editRedoAction // applies to text inside card

func (me *App) editCopy() {
	fmt.Println("editCopy") // TODO
}

func (me *App) editCut() {
	fmt.Println("editCut") // TODO
}

func (me *App) editPaste() {
	fmt.Println("editPaste") // TODO
}

// TODO &Bold			editBoldAction
// TODO &Italic			editItalicAction
// TODO &Monospace		editMonospaceAction
// TODO &Bullet List	editBulletListAction
// TODO &Numbered List	editNumberedListAction
// TODO &End List		editEndListAction

// TODO Insert &Symbol…	editInsertSymbolAction
// dialog that shows common symbols (and recently inserted ones?) plus the
// option of choosing by code point?
