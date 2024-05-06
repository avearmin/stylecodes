package bgcolor

import (
	"fmt"
	"github.com/avearmin/stylecodes/internal/hex"
)

const (
	Black   = "\033[40m"
	Red     = "\033[41m"
	Green   = "\033[42m"
	Yellow  = "\033[43m"
	Blue    = "\033[44m"
	Magenta = "\033[45m"
	Cyan    = "\033[46m"
	White   = "\033[47m"
	End     = "\033[49m"
)

func Hex(h string) string {
	r, g, b, err := hex.ToRGB(h)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("\033[48;2;%d;%d;%dm", r, g, b)
}

func RGB(r, g, b int) string {
	if (r < 0 || g < 0 || b < 0) || (r > 255 || g > 255 || b > 255) {
		return ""
	}
	return fmt.Sprintf("\033[48;2;%d;%d;%dm", r, g, b)
}
