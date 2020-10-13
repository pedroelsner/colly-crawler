package format

import (
	"strconv"
	"strings"
)

func Trim(s string) string {
	return strings.TrimSpace(s)
}

func Currency(s string) float64 {
	var value string

	value = strings.TrimSpace(s)
	value = strings.Trim(value, "R$ ")
	value = strings.ReplaceAll(value, ".", "")
	value = strings.ReplaceAll(value, ",", ".")

	n, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0.00
	}
	return n
}
