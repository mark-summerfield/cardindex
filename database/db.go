// Copyright © 2025 Mark Summerfield. All rights reserved.
// License: GPL-3

package database

import "database/sql"

func SqliteVersion() (string, error) {
	db, err := sql.Open(DRIVER, ":memory:")
	if err != nil {
		return "", err
	}
	defer db.Close()
	row := db.QueryRow(SQL_VERSION)
	var version string
	if err := row.Scan(&version); err != nil {
		return version, err
	}
	return version, nil
}
