package stylecodes

import (
	"strings"
)

const (
	EndAll = "\033[0m"
)

func Style(text string, codes ...string) string {
	return strings.Join(codes, "") + text + EndAll
}
