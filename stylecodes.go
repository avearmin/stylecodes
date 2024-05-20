package stylecodes

import (
	"regexp"
	"strings"
)

const (
	ResetAll = "\033[0m"
)

func Style(text string, codes ...string) string {
	lines := strings.Split(text, "\n")

	var codeStr strings.Builder
	for _, v := range codes {
		codeStr.WriteString(v)
	}

	var builder strings.Builder
	for i, v := range lines {
		builder.WriteString(codeStr.String() + v + ResetAll)
		if i != len(lines)-1 {
			builder.WriteString("\n")
		}
	}

	return builder.String()
}

func StripANSI(s string) string {
	re := regexp.MustCompile("[\u001B\u009B][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[a-zA-Z\\d]*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PRZcf-ntqry=><~]))")
	return re.ReplaceAllString(s, "")
}
