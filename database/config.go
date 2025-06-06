// Copyright © 2025 Mark Summerfield. All rights reserved.
// License: GPL-3

package database

import "time"

func (me *Database) ConfigCreated() (time.Time, error) {
	return me.configWhen(CREATED)
}

func (me *Database) ConfigUpdated() (time.Time, error) {
	return me.configWhen(UPDATED)
}

func (me *Database) configWhen(key string) (time.Time, error) {
	var when time.Time
	var text string
	row := me.db.QueryRow(SQL_CONFIG_GET_WHEN, key)
	if err := row.Scan(&text); err != nil {
		return when, err
	}
	if when, err := time.Parse(time.DateTime, text); err != nil {
		return when, err
	} else {
		return when, nil
	}
}
