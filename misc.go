// Copyright © 2025 Mark Summerfield. All rights reserved.
// License: GPL-3

package main

import (
	"slices"
	"strconv"
	"strings"

	"github.com/mark-summerfield/ufunc"
)

func strForBids(bids []int) string {
	return strings.Join(slices.Collect(ufunc.Map(bids,
		func(i int) (string, bool) {
			return strconv.Itoa(i), true
		})), ",")
}

func bidsForStr(text string) []int {
	return slices.Collect(ufunc.Map(strings.Split(text, ","),
		func(s string) (int, bool) {
			if i, err := strconv.Atoi(s); err == nil {
				return i, true
			} else {
				return 0, false
			}
		}))
}
