// Copyright © 2025 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

func (me *Model) Box(bid int) (Box, error) {
	box := Box{bid: bid}
	row := me.db.QueryRow(SQL_BOX_GET, bid)
	err := row.Scan(&box.name)
	return box, err
}

func (me *Model) BoxAdd(name string) (int, error) {
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

func (me *Model) BoxEdit(bid int, name string) error {
	_, err := me.db.Exec(SQL_BOX_UPDATE, name, bid)
	return err
}

func (me *Model) BoxInUse(bid int) (bool, error) {
	var in_use bool
	row := me.db.QueryRow(SQL_BOX_IN_USE, bid)
	err := row.Scan(&in_use)
	return in_use, err
}

// Will fail if the box is in use
func (me *Model) BoxDelete(bid int) error {
	_, err := me.db.Exec(SQL_BOX_DELETE, bid)
	return err
}

func (me *Model) BoxAddCard(cid, bid int) error {
	_, err := me.db.Exec(SQL_BOX_ADD_CARD, cid, bid)
	return err
}

func (me *Model) BoxRemoveCard(cid, bid int) error {
	_, err := me.db.Exec(SQL_BOX_REMOVE_CARD, cid, bid)
	return err
}

func (me *Model) Boxes() ([]Box, error) {
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
