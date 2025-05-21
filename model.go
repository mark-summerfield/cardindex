// Copyright © 2025 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"database/sql"
	"errors"

	"github.com/mark-summerfield/ufile"
	_ "modernc.org/sqlite"
)

type Model struct {
	filename string
	db       *sql.DB
}

func NewModel(filename string) (*Model, error) {
	exists := ufile.FileExists(filename)
	db, err := sql.Open(DRIVER, filename)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(SQL_PREPARE)
	if err != nil {
		return nil, errors.Join(err, db.Close())
	}
	if !exists {
		_, err := db.Exec(SQL_CREATE)
		if err != nil {
			return nil, errors.Join(err, db.Close())
		}
	}
	return &Model{filename, db}, nil
}

func (me *Model) Close() error {
	var err error
	if me.db != nil {
		sql := SQL_BEGIN + SQL_CONFIG_UPDATE
		n := 0
		row := me.db.QueryRow(SQL_CONFIG_GET_N)
		if err = row.Scan(&n); err == nil {
			if n >= MAX_OPENS {
				sql += SQL_CONFIG_ZERO_N
			}
			sql += SQL_COMMIT
			if n >= MAX_OPENS {
				sql += SQL_OPTIMIZE
			}
			_, err = me.db.Exec(sql)
		}
		err = errors.Join(err, me.db.Close())
		me.db = nil
	}
	return err
}

func (me *Model) Filename() string { return me.filename }
