// Copyright © 2025 Mark Summerfield. All rights reserved.
// License: GPL-3

package database

import (
	"strings"
	"time"
)

func (me *Database) CardCounts() (CardCounts, error) {
	var counts CardCounts
	row := me.db.QueryRow(SQL_CARD_COUNTS)
	if err := row.Scan(&counts.Visible, &counts.Unboxed,
		&counts.Hidden); err != nil {
		return counts, err
	}
	return counts, nil
}

func (me *Database) Card(cid int) (Card, error) {
	card := Card{cid: cid}
	row := me.db.QueryRow(SQL_CARD_GET, cid)
	var created, updated string
	var err error
	if err = row.Scan(&card.body, &card.hidden, &created,
		&updated); err == nil {
		var when time.Time
		if when, err = time.Parse(time.DateTime, created); err == nil {
			card.created = when
		}
		if when, err = time.Parse(time.DateTime, updated); err == nil {
			card.updated = when
		}
	}
	return card, err
}

func (me *Database) CardAdd(body string) (int, error) {
	reply, err := me.db.Exec(SQL_CARD_INSERT, body)
	if err != nil {
		return INVALID_ID, err
	}
	if cid, err := reply.LastInsertId(); err != nil {
		return INVALID_ID, err
	} else {
		return int(cid), nil
	}
}

func (me *Database) CardEdit(cid int, body string) error {
	_, err := me.db.Exec(SQL_CARD_UPDATE, body, cid)
	return err
}

func (me *Database) CardHidden(cid int) (bool, error) {
	var hidden bool
	row := me.db.QueryRow(SQL_CARD_GET_HIDDEN, cid)
	err := row.Scan(&hidden)
	return hidden, err
}

func (me *Database) CardHide(cid int) error {
	_, err := me.db.Exec(SQL_CARD_VISIBILITY, true, cid)
	return err
}

func (me *Database) CardUnhide(cid int) error {
	_, err := me.db.Exec(SQL_CARD_VISIBILITY, false, cid)
	return err
}

func (me *Database) CardDelete(cid int) error {
	_, err := me.db.Exec(SQL_CARD_DELETE, cid)
	return err
}

func (me *Database) CardNamesVisible(oid Oid) ([]CardName, error) {
	return me.cardNames(SQL_CARD_NAMES_VISIBLE, oid)
}

func (me *Database) CardNamesUnboxed(oid Oid) ([]CardName, error) {
	return me.cardNames(SQL_CARD_NAMES_UNBOXED, oid)
}

func (me *Database) CardNamesHidden(oid Oid) ([]CardName, error) {
	return me.cardNames(SQL_CARD_NAMES_HIDDEN, oid)
}

func (me *Database) CardNamesForSid(sid int) ([]CardName, error) {
	if search, err := me.Search(sid); err == nil {
		return me.CardNamesForSearch(search)
	} else {
		return nil, err
	}
}

func (me *Database) CardNamesForSearch(search Search) ([]CardName, error) {
	query, args := search.Query(false)
	return me.cardNames(query, search.oid, args...)
}

func (me *Database) cardNames(query string, oid Oid, args ...any) ([]CardName,
	error,
) {
	query, _ = strings.CutSuffix(query, ";")
	query += " " + oid.Query() + ";"
	rows, err := me.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var cardnames []CardName
	for rows.Next() {
		var cardname CardName
		if err = rows.Scan(&cardname.cid, &cardname.name); err != nil {
			return nil, err
		}
		cardnames = append(cardnames, cardname)
	}
	return cardnames, nil
}
