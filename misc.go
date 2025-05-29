// Copyright © 2025 Mark Summerfield. All rights reserved.

package main

func int8ToStr(raw []int8) string {
	data := make([]byte, 0, len(raw))
	for _, i := range raw {
		if i == 0 {
			break
		}
		data = append(data, byte(i))
	}
	return string(data)
}
