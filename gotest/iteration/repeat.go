package iteration

import "strings"

func Repeat(c string, n int) string {
	var repeated string
	for i := 0; i < n; i++ {
		repeated += c
	}
	return strings.ToUpper(repeated)
}
