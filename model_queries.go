// Copyright © 2025 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

func (me *Model) Query(qid int) (Query, error) {
	query := Query{qid: qid}
	var inboxes, notinboxes string
	row := me.db.QueryRow(SQL_QUERY_GET, qid)
	err := row.Scan(&query.name, &query.matchText, &inboxes, &notinboxes,
		&query.hidden, &query.oid)
	query.inBoxes = bidsForStr(inboxes)
	query.notInBoxes = bidsForStr(notinboxes)
	return query, err
}

func (me *Model) QueryAdd(query Query) (int, error) {
	reply, err := me.db.Exec(SQL_QUERY_INSERT, query.name, query.matchText,
		strForBids(query.inBoxes), strForBids(query.notInBoxes),
		query.hidden, query.oid)
	if err != nil {
		return INVALID_ID, err
	}
	if qid, err := reply.LastInsertId(); err != nil {
		return INVALID_ID, err
	} else {
		return int(qid), nil
	}
}

func (me *Model) QueryEdit(query Query) error {
	_, err := me.db.Exec(SQL_QUERY_UPDATE, query.qid, query.name,
		query.matchText, strForBids(query.inBoxes),
		strForBids(query.notInBoxes), query.hidden, query.oid)
	return err
}

func (me *Model) QueryDelete(qid int) error {
	_, err := me.db.Exec(SQL_QUERY_DELETE, qid)
	return err
}

func (me *Model) Queries() ([]Query, error) {
	rows, err := me.db.Query(SQL_QUERIES)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var queries []Query
	for rows.Next() {
		var query Query
		var inboxes, notinboxes string
		if err = rows.Scan(&query.qid, &query.name, &query.matchText,
			&inboxes, &notinboxes, &query.hidden, &query.oid); err != nil {
			return nil, err
		}
		query.inBoxes = bidsForStr(inboxes)
		query.notInBoxes = bidsForStr(notinboxes)
		queries = append(queries, query)
	}
	return queries, nil
}
