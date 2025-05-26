// Copyright © 2025 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"log"

	qt "github.com/mappu/miqt/qt6"
)

func getIcon(name string) *qt.QIcon {
	name = "images/" + name
	if data, err := ICONS.ReadFile(name); err != nil {
		log.Printf("failed to read icon %q: %v\n", name, err)
		return qt.NewQIcon()
	} else {
		image := qt.QImage_FromDataWithData(data)
		pixmap := qt.QPixmap_FromImage(image)
		return qt.NewQIcon2(pixmap)
	}
}
