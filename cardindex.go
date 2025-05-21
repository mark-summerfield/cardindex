// Copyright © 2025 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"fmt"
)

func main() {
	if version, err := SqliteVersion(); err == nil {
		fmt.Printf("%s v%s (SQLite v%s)\n", APPNAME,
			Version[:len(Version)-1], version)
	} else {
		fmt.Println(err)
	}
}
