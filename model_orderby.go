// Copyright © 2025 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import "fmt"

func orderBy(by string) string {
	switch by {
	case BY_NAME:
		return "ORDER BY LOWER(Name)"
	case BY_NEW_TO_OLD:
		return "ORDER BY updated DESC"
	case BY_OLD_TO_NEW:
		return "ORDER BY created"
	}
	panic(fmt.Sprintf("invalid orderBy: %q", by))
}
