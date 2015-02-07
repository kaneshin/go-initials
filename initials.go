package initials

import (
	"bytes"
	"strings"
)

func getInitialIfASCII(s string) string {
	if len(s) > 0 {
		r := []rune(s)
		r0 := r[0]
		if (0x41 <= r0 && r0 <= 0x5A) || (0x61 <= r0 && r0 <= 0x7A) {
			return strings.ToUpper(string(r0))
		}
	}
	return ""
}

func GetInitialsIfASCII(ss ...string) string {
	var buf bytes.Buffer
	for _, s := range ss {
		buf.WriteString(getInitialIfASCII(s))
	}
	return buf.String()
}
