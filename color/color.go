package color

import (
	"fmt"
	"github.com/avearmin/stylecodes/internal/hex"
)

const (
	Black   = "\033[30m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"
	End     = "\033[39m"
)

func Hex(h string) string {
	r, g, b, err := hex.ToRGB(h)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
}
