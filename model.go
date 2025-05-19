// Copyright © 2025 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/mark-summerfield/ufile"
	_ "modernc.org/sqlite"
)

type Model struct {
	db *sql.DB
}

type Counts struct {
	Visible int
	Unboxed int
	Hidden  int
}

func NewModel(filename string) (*Model, error) {
	exists := ufile.FileExists(filename)
	db, err := sql.Open("sqlite", filename)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(SQLprepare)
	if err != nil {
		return nil, errors.Join(err, db.Close())
	}
	if !exists {
		_, err := db.Exec(SQLcreate)
		if err != nil {
			return nil, errors.Join(err, db.Close())
		}
	}
	return &Model{db}, nil
}

func (me *Model) Close() error {
	var err error
	if me.db != nil {
		err = me.db.Close()
		me.db = nil
	}
	return err
}

func (me *Model) Counts() (*Counts, error) {
	row := me.db.QueryRow("SELECT * FROM Counts")
	fmt.Println(row)
	return &Counts{}, nil
}
