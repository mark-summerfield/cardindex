// Copyright © 2025 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

func orderBy(by string) string {
	switch by {
	case NAME, "N": // NAME="Name" (from UI); "N" (from db Query)
		return "ORDER BY LOWER(Name)"
	case UPDATED, "U":
		return "ORDER BY updated DESC"
	case CREATED, "C":
		return "ORDER BY created"
	}
	return "" // unordered
}

func orderById(by string) string {
	switch by {
	case NAME, "N":
		return "N"
	case UPDATED, "U":
		return "U"
	case CREATED, "C":
		return "C"
	}
	return "I" // ignore (unordered)
}
