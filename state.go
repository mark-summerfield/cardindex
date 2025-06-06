// Copyright © 2025 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import "github.com/mark-summerfield/cardindex/database"

type State struct {
	onError func(string)
	db      *database.Database
}

func NewState(onError func(string), db *database.Database) *State {
	return &State{onError, db}
}
