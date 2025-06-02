// Copyright © 2025 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"log"
	"runtime/debug"
	"strings"

	"github.com/mappu/miqt/qt"
)

func getIcon(filename string) *qt.QIcon {
	filename = "images/" + filename
	if data, err := ICONS.ReadFile(filename); err != nil {
		log.Printf("failed to read icon %q: %v\n", filename, err)
		return qt.NewQIcon()
	} else {
		image := qt.QImage_FromDataWithData(data)
		pixmap := qt.QPixmap_FromImage(image)
		return qt.NewQIcon2(pixmap)
	}
}

// TODO replace with QtVersion() when available
func miqtVersion() string {
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, dependency := range info.Deps {
			if strings.Contains(dependency.Path, "miqt") {
				version, _ := strings.CutPrefix(dependency.Version, "v")
				return version
			}
		}
	}
	return ""
}
