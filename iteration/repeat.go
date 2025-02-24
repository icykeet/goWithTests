package iteration

import "strings"

func Repeat(ch string, repeatCount int) string {
	var repeated strings.Builder

	for i := 0; i < repeatCount; i++ {
		repeated.WriteString(ch)
	}

	return repeated.String()
}
