package main

import (
	"strconv"
)

func IntToHex(i int64) []byte {
	return []byte(strconv.FormatInt(i, 16))
}
