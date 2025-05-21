// Copyright © 2025 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import "time"

func (me *Model) ConfigCreated() (time.Time, error) {
	return me.configWhen(CREATED)
}

func (me *Model) ConfigUpdated() (time.Time, error) {
	return me.configWhen(UPDATED)
}

func (me *Model) configWhen(key string) (time.Time, error) {
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
