// Copyright © 2025 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import "fmt"

func orderBy(by string) string {
	switch by {
	case NAME:
		return "ORDER BY LOWER(Name)"
	case UPDATED:
		return "ORDER BY updated DESC"
	case CREATED:
		return "ORDER BY created"
	}
	panic(fmt.Sprintf("invalid order by: %q", by))
}
