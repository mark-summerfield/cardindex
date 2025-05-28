// Copyright © 2025 Mark Summerfield. All rights reserved.
// License: GPL-3

package model

func (me *Model) Search(sid int) (Search, error) {
	search := Search{sid: sid}
	row := me.db.QueryRow(SQL_SEARCH_GET, sid)
	err := row.Scan(&search.searchText, &search.hidden, &search.oid)
	return search, err
}

func (me *Model) SearchAdd(search Search) (int, error) {
	reply, err := me.db.Exec(SQL_SEARCH_INSERT, search.searchText,
		search.hidden, search.oid)
	if err != nil {
		return INVALID_ID, err
	}
	if sid, err := reply.LastInsertId(); err != nil {
		return INVALID_ID, err
	} else {
		return int(sid), nil
	}
}

func (me *Model) SearchEdit(search Search) error {
	_, err := me.db.Exec(SQL_SEARCH_UPDATE, search.sid, search.searchText,
		search.hidden, search.oid)
	return err
}

func (me *Model) SearchDelete(sid int) error {
	_, err := me.db.Exec(SQL_SEARCH_DELETE, sid)
	return err
}

func (me *Model) Searches() ([]Search, error) {
	rows, err := me.db.Query(SQL_SEARCHES)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var searches []Search
	for rows.Next() {
		var search Search
		if err = rows.Scan(&search.sid, &search.searchText, &search.hidden,
			&search.oid); err != nil {
			return nil, err
		}
		searches = append(searches, search)
	}
	return searches, nil
}
