// Copyright © 2025 Mark Summerfield. All rights reserved.
// License: GPL-3

package database

import (
	"database/sql"
	"errors"

	"github.com/mark-summerfield/ufile"
	_ "modernc.org/sqlite"
)

type Database struct {
	filename string
	db       *sql.DB
}

func NewDatabase(filename string) (*Database, error) {
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
	return &Database{filename, db}, nil
}

func (me *Database) Close() error {
	var err error
	if me.db != nil {
		query := SQL_BEGIN + SQL_CONFIG_UPDATE
		n := 0
		row := me.db.QueryRow(SQL_CONFIG_GET_N)
		if err = row.Scan(&n); err == nil {
			if n >= MAX_OPENS {
				query += SQL_CONFIG_ZERO_N
			} else {
				query += SQL_CONFIG_INC_N
			}
			query += SQL_COMMIT
			if n >= MAX_OPENS {
				query += SQL_OPTIMIZE
			}
			_, err = me.db.Exec(query)
		}
		err = errors.Join(err, me.db.Close())
		me.db = nil
	}
	return err
}

func (me *Database) Filename() string { return me.filename }
