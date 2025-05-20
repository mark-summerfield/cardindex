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
	return &Model{filename, db}, nil
}

func (me *Model) Close() error {
	var err error
	if me.db != nil {
		sql := `BEGIN;
				UPDATE Config SET Value = JULIANDAY('NOW')
					WHERE Key = 'Updated';`
		n := 0
		row := me.db.QueryRow("SELECT Value FROM Config WHERE Key = N")
		if err = row.Scan(&n); err == nil {
			if n >= MaxOpens {
				sql += "UPDATE Config SET Value = 0 WHERE Key = N;"
			}
			sql += "COMMIT;"
			if n >= MaxOpens {
				sql += `INSERT INTO vt_fts_cards(vt_fts_cards)
							VALUES ('optimize');
						VACUUM;`
			}
			_, err = me.db.Exec(sql)
		}
		err = errors.Join(err, me.db.Close())
		me.db = nil
	}
	return err
}

func (me *Model) Version() (string, error) {
	row := me.db.QueryRow("SELECT SQLITE_VERSION()")
	var data string
	if err := row.Scan(&data); err != nil {
		return data, err
	}
	return data, nil
}

func (me *Model) Filename() string { return me.filename }

func (me *Model) Counts() (*Counts, error) {
	counts := &Counts{}
	row := me.db.QueryRow("SELECT Visible, Unboxed, Hidden FROM Counts")
	if err := row.Scan(&counts.Visible, &counts.Unboxed,
		&counts.Hidden); err != nil {
		return counts, err
	}
	return counts, nil
}

func (me *Model) ConfigCreated() (time.Time, error) {
	return me.configWhen("Created")
}

func (me *Model) ConfigUpdated() (time.Time, error) {
	return me.configWhen("Updated")
}

func (me *Model) configWhen(key string) (time.Time, error) {
	var when time.Time
	var data string
	row := me.db.QueryRow(
		"SELECT DATETIME(Value) FROM Config WHERE Key = ?", key)
	if err := row.Scan(&data); err != nil {
		return when, err
	}
	if when, err := time.Parse(time.DateTime, data); err != nil {
		return when, err
	} else {
		return when, nil
	}
}

// TODO
//
//  2. ConfigInt(key) → int value
//  3. ConfigRaw(key) → []byte value
//  4. ConfigStr(key) → string value
//  5. SetConfigItem(key, []byte value)
//
//  6. CardAdd(string) → cid
//  7. CardEdit(cid, string)
//  8. CardHide(cid)
//  9. CardUnhide(cid)
// 10. CardDelete(cid)
//
// 11. BoxAdd(string) → bid
// 12. BoxEdit(bid, string)
// 13. BoxDelete(bid)
//
// 14. AddCardToBox(cid, bid)
// 15. RemoveCardFromBox(cid, bid)
//
// 16. QueryAdd(query) → qid
// 17. QueryEdit(qid, query)
// 18. QueryDelete(qid)
//
// iterators:
//		19. AllVisibleCards() → iter.Seq…
//		20. AllUnboxedCards() → iter.Seq…
//		21. AllHiddenCards() → iter.Seq…
//		22. AllQueriedCards(query)() → iter.Seq…
//
//		23. AllBoxes() → iter.Seq…
