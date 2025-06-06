// Copyright © 2025 Mark Summerfield. All rights reserved.
// License: GPL-3

package database

import (
	"context"
	"errors"
)

func (me *Database) Box(bid int) (Box, error) {
	box := Box{bid: bid}
	row := me.db.QueryRow(SQL_BOX_GET, bid)
	err := row.Scan(&box.name)
	return box, err
}

func (me *Database) BoxAdd(name string) (int, error) {
	reply, err := me.db.Exec(SQL_BOX_INSERT, name)
	if err != nil {
		return INVALID_ID, err
	}
	if bid, err := reply.LastInsertId(); err != nil {
		return INVALID_ID, err
	} else {
		return int(bid), nil
	}
}

func (me *Database) BoxEdit(bid int, name string) error {
	_, err := me.db.Exec(SQL_BOX_UPDATE, name, bid)
	return err
}

func (me *Database) BoxInUse(bid int) (bool, error) {
	var in_use bool
	row := me.db.QueryRow(SQL_BOX_IN_USE, bid)
	err := row.Scan(&in_use)
	return in_use, err
}

// Will fail if the box is in use
func (me *Database) BoxDelete(bid int) error {
	_, err := me.db.Exec(SQL_BOX_DELETE, bid)
	return err
}

func (me *Database) BoxAddCards(bid int, cids ...int) error {
	if tx, err := me.db.BeginTx(context.Background(), nil); err == nil {
		for _, cid := range cids {
			if _, err = tx.Exec(SQL_BOX_ADD_CARD, cid, bid); err != nil {
				break
			}
		}
		if err == nil {
			return tx.Commit()
		} else {
			return errors.Join(err, tx.Rollback())
		}
	} else {
		return err
	}
}

func (me *Database) BoxRemoveCard(bid, cid int) error {
	_, err := me.db.Exec(SQL_BOX_REMOVE_CARD, cid, bid)
	return err
}

func (me *Database) Boxes() ([]Box, error) {
	rows, err := me.db.Query(SQL_BOXES)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var boxes []Box
	for rows.Next() {
		var box Box
		if err = rows.Scan(&box.bid, &box.name); err != nil {
			return nil, err
		}
		boxes = append(boxes, box)
	}
	return boxes, nil
}
