// Copyright © 2025 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	_ "embed"
)

//go:embed Version.dat
var Version string

//go:embed sql/prepare.sql
var SQLprepare string

//go:embed sql/create.sql
var SQLcreate string

const (
	APPNAME = "CardIndex"

	MaxOpens = 11
)

// var E100 = errors.New("E100: failed to get counts")
