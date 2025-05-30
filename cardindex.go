// Copyright © 2025 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"fmt"
	"os"

	"github.com/mappu/miqt/qt"
)

func main() {
	if !(len(os.Args) > 1 && os.Args[1] == "--debug") {
		defer func() {
			if r := recover(); r != nil {
				message := fmt.Sprintf("Unrecoverable error: %s", r)
				qt.QMessageBox_Critical5(nil, "Error — "+APPNAME, message,
					qt.QMessageBox__Close)
				fmt.Println(message)
			}
		}()
	}

	qt.NewQApplication(os.Args)
	qt.QCoreApplication_SetApplicationName(APPNAME)
	app := NewApp()
	app.Show()
	qt.QApplication_Exec()
}
