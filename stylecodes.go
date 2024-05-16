package stylecodes

import (
	"strings"
)

const (
	EndAll = "\033[0m"
)

func Style(text string, codes ...string) string {
	lines := strings.Split(text, "\n")
	
	var codeStr strings.Builder
	for _, v := range codes {
		codeStr.WriteString(v)
	}

	var builder strings.Builder
	for i, v := range lines {
		builder.WriteString(codeStr.String() + v + EndAll)
		if i != len(lines) - 1 {
			builder.WriteString("\n")
		}
	}

	return builder.String()
}
