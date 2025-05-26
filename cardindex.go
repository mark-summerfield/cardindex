// Copyright © 2025 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"fmt"
	"os"

	qt "github.com/mappu/miqt/qt6"
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
	app := NewApp()
	app.Show()
	qt.QApplication_Exec()
}
