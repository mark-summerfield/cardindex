// Copyright © 2025 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"strings"

	"github.com/mappu/miqt/qt"
)

func (me *App) MakeConnections() {
	me.makeFileConnections()
	me.makeEditConnections()
	me.makeCardConnections()
	me.makeBoxConnections()
	me.makeSearchConnections()
	me.makeWindowConnections()
	me.makeHelpConnections()
}

func (me *App) makeFileConnections() {
	me.fileNewAction.OnTriggered(func() { me.fileNew() })
	me.fileOpenAction.OnTriggered(func() { me.fileOpen() })
	me.fileSaveAction.OnTriggered(func() { me.fileSave() })
	me.fileSaveAsAction.OnTriggered(func() { me.fileSaveAs() })
	me.fileExportAction.OnTriggered(func() { me.fileExport() })
	me.fileConfigureAction.OnTriggered(func() { me.fileConfigure() })
	me.fileQuitAction.OnTriggered(func() { me.window.Close() })
	for _, action := range me.fileOpenActions {
		action.OnTriggered(func() {
			filename := action.ToolTip()
			if strings.HasPrefix(filename, "&") {
				i := strings.Index(filename, " ")
				filename = strings.TrimSpace(filename[i:])
			}
			me.fileOpenDatabase(filename)
		})
	}
}

func (me *App) makeEditConnections() {
	me.editUndoAction.OnTriggered(func() { me.editUndo() })
	me.editRedoAction.OnTriggered(func() { me.editRedo() })
	me.editCopyAction.OnTriggered(func() { me.editCopy() })
	me.editCutAction.OnTriggered(func() { me.editCut() })
	me.editPasteAction.OnTriggered(func() { me.editPaste() })
	me.editBoldAction.OnTriggered(func() { me.editBold() })
	me.editItalicAction.OnTriggered(func() { me.editItalic() })
	me.editMonospaceAction.OnTriggered(func() { me.editMonospace() })
	me.editBulletListAction.OnTriggered(func() { me.editBulletList() })
	me.editNumberedListAction.OnTriggered(func() { me.editNumberedList() })
	me.editEndListAction.OnTriggered(func() { me.editEndList() })
	me.editInsertWebLinkAction.OnTriggered(
		func() { me.editInsertWebLink() })
	me.editInsertFileLinkAction.OnTriggered(
		func() { me.editInsertFileLink() })
	me.editInsertSymbolAction.OnTriggered(func() { me.editInsertSymbol() })
}

func (me *App) makeCardConnections() {
	me.cardNewAction.OnTriggered(func() { me.cardNew() })
	me.cardViewVisibleAction.OnTriggered(func() { me.cardViewVisible() })
	me.cardViewUnboxedAction.OnTriggered(func() { me.cardViewUnboxed() })
	me.cardViewHiddenAction.OnTriggered(func() { me.cardViewHidden() })
	me.cardAddToBoxAction.OnTriggered(func() { me.cardAddToBox() })
	me.cardRemoveFromBoxAction.OnTriggered(
		func() { me.cardRemoveFromBox() })
	me.cardExportAction.OnTriggered(func() { me.cardExport() })
	me.cardUnhideAction.OnTriggered(func() { me.cardUnhide() })
	me.cardHideAction.OnTriggered(func() { me.cardHide() })
	me.cardDeleteAction.OnTriggered(func() { me.cardDelete() })
}

func (me *App) makeBoxConnections() {
	me.boxNewAction.OnTriggered(func() { me.boxNew() })
	me.boxViewAction.OnTriggered(func() { me.boxView() })
	me.boxAddFromSearchAction.OnTriggered(func() { me.boxAddFromSearch() })
	me.boxAddFromBoxAction.OnTriggered(func() { me.boxAddFromBox() })
	me.boxDeleteAction.OnTriggered(func() { me.boxDelete() })
}

func (me *App) makeSearchConnections() {
	me.searchNewAction.OnTriggered(func() { me.searchNew() })
	me.searchViewAction.OnTriggered(func() { me.searchView() })
	me.searchDeleteAction.OnTriggered(func() { me.searchDelete() })
}

func (me *App) makeWindowConnections() {
	me.windowNextAction.OnTriggered(
		func() { me.mdiArea.ActivateNextSubWindow() })
	me.windowPrevAction.OnTriggered(
		func() { me.mdiArea.ActivatePreviousSubWindow() })
	me.windowCascadeAction.OnTriggered(
		func() { me.mdiArea.CascadeSubWindows() })
	me.windowTileAction.OnTriggered(func() { me.mdiArea.TileSubWindows() })
	me.windowCloseAction.OnTriggered(
		func() { me.mdiArea.CloseActiveSubWindow() })
}

func (me *App) makeHelpConnections() {
	me.helpHelpAction.OnTriggered(func() { me.helpHelp() })
	me.helpAboutAction.OnTriggered(func() { me.helpAbout() })
	me.window.OnCloseEvent(func(super func(event *qt.QCloseEvent),
		event *qt.QCloseEvent,
	) {
		me.SaveSettings()
		me.fileCloseDatabase() // must be last since it closes current db
	})
}
