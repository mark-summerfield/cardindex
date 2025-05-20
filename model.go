// Copyright © 2025 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"database/sql"
	"errors"
	"time"

	"github.com/mark-summerfield/ufile"
	_ "modernc.org/sqlite"
)

type Model struct {
	filename string
	db       *sql.DB
}

type Counts struct {
	Visible int
	Unboxed int
	Hidden  int
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
		sql := SQL_BEGIN + SQL_UPDATE_CONFIG_UPDATED
		n := 0
		row := me.db.QueryRow(SQL_GET_CONFIG_N)
		if err = row.Scan(&n); err == nil {
			if n >= MAX_OPENS {
				sql += SQL_ZERO_CONFIG_N
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

func (me *Model) Version() (string, error) {
	row := me.db.QueryRow(SQL_VERSION)
	var data string
	if err := row.Scan(&data); err != nil {
		return data, err
	}
	return data, nil
}

func (me *Model) Filename() string { return me.filename }

func (me *Model) ConfigCreated() (time.Time, error) {
	return me.configWhen(CREATED)
}

func (me *Model) ConfigUpdated() (time.Time, error) {
	return me.configWhen(UPDATED)
}

func (me *Model) configWhen(key string) (time.Time, error) {
	var when time.Time
	var data string
	row := me.db.QueryRow(SQL_GET_WHEN, key)
	if err := row.Scan(&data); err != nil {
		return when, err
	}
	if when, err := time.Parse(time.DateTime, data); err != nil {
		return when, err
	} else {
		return when, nil
	}
}

func (me *Model) Counts() (*Counts, error) {
	counts := &Counts{}
	row := me.db.QueryRow(SQL_GET_COUNTS)
	if err := row.Scan(&counts.Visible, &counts.Unboxed,
		&counts.Hidden); err != nil {
		return counts, err
	}
	return counts, nil
}

func (me *Model) CardAdd(body string) (int, error) {
	reply, err := me.db.Exec(SQL_INSERT_CARD, body)
	if err != nil {
		return -1, err
	}
	if cid, err := reply.LastInsertId(); err != nil {
		return -1, err
	} else {
		return int(cid), nil
	}
}

func (me *Model) CardEdit(cid int, body string) error {
	_, err := me.db.Exec(SQL_UPDATE_CARD, body, cid)
	return err
}

func (me *Model) CardBody(cid int) (string, error) {
	var body string
	row := me.db.QueryRow(SQL_CARD_BODY, cid)
	err := row.Scan(&body)
	return body, err
}

func (me *Model) CardHidden(cid int) (bool, error) {
	var hidden bool
	row := me.db.QueryRow(SQL_CARD_HIDDEN, cid)
	err := row.Scan(&hidden)
	return hidden, err
}

func (me *Model) CardHide(cid int) error {
	_, err := me.db.Exec(SQL_SET_CARD_VISIBILITY, true, cid)
	return err
}

func (me *Model) CardUnhide(cid int) error {
	_, err := me.db.Exec(SQL_SET_CARD_VISIBILITY, false, cid)
	return err
}

func (me *Model) CardDelete(cid int) error {
	_, err := me.db.Exec(SQL_DELETE_CARD, cid)
	return err
}

// TODO
// iterators:
//		AllVisibleCards() → iter.Seq…
//		AllUnboxedCards() → iter.Seq…
//		AllHiddenCards() → iter.Seq…
//
// BoxAdd(string) → bid
// BoxEdit(bid, string)
// BoxDelete(bid)
//
// AddCardToBox(cid, bid)
// RemoveCardFromBox(cid, bid)
//
// QueryAdd(query) → qid
// QueryEdit(qid, query)
// QueryDelete(qid)
//
// iterators:
//		AllQueriedCards(query)() → iter.Seq…
//
//		AllBoxes() → iter.Seq…
