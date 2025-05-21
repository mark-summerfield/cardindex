// Copyright © 2025 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"strings"
	"time"
)

func (me *Model) CardCounts() (CardCounts, error) {
	var counts CardCounts
	row := me.db.QueryRow(SQL_CARD_COUNTS)
	if err := row.Scan(&counts.Visible, &counts.Unboxed,
		&counts.Hidden); err != nil {
		return counts, err
	}
	return counts, nil
}

func (me *Model) Card(cid int) (Card, error) {
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

func (me *Model) CardAdd(body string) (int, error) {
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

func (me *Model) CardEdit(cid int, body string) error {
	_, err := me.db.Exec(SQL_CARD_UPDATE, body, cid)
	return err
}

func (me *Model) CardHidden(cid int) (bool, error) {
	var hidden bool
	row := me.db.QueryRow(SQL_CARD_GET_HIDDEN, cid)
	err := row.Scan(&hidden)
	return hidden, err
}

func (me *Model) CardHide(cid int) error {
	_, err := me.db.Exec(SQL_CARD_VISIBILITY, true, cid)
	return err
}

func (me *Model) CardUnhide(cid int) error {
	_, err := me.db.Exec(SQL_CARD_VISIBILITY, false, cid)
	return err
}

func (me *Model) CardDelete(cid int) error {
	_, err := me.db.Exec(SQL_CARD_DELETE, cid)
	return err
}

func (me *Model) CardNamesVisible(by string) ([]CardName, error) {
	return me.cardNames(SQL_CARD_NAMES_VISIBLE, by)
}

func (me *Model) CardNamesUnboxed(by string) ([]CardName, error) {
	return me.cardNames(SQL_CARD_NAMES_UNBOXED, by)
}

func (me *Model) CardNamesHidden(by string) ([]CardName, error) {
	return me.cardNames(SQL_CARD_NAMES_HIDDEN, by)
}

func (me *Model) cardNames(sql, by string) ([]CardName, error) {
	sql, _ = strings.CutSuffix(sql, ";")
	sql += " " + orderBy(by) + ";"
	rows, err := me.db.Query(sql)
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
