// Copyright © 2025 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import _ "embed"

//go:embed Version.dat
var Version string

const (
	APPNAME = "CardIndex"
)
