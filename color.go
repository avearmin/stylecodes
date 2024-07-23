package stylecodes

import (
	"fmt"

	"github.com/avearmin/stylecodes/internal/hex"
)

const (
	ColorBlack         = "\033[30m"
	ColorRed           = "\033[31m"
	ColorGreen         = "\033[32m"
	ColorYellow        = "\033[33m"
	ColorBlue          = "\033[34m"
	ColorMagenta       = "\033[35m"
	ColorCyan          = "\033[36m"
	ColorWhite         = "\033[37m"
	ColorBrightBlack   = "\033[90m"
	ColorBrightRed     = "\033[91m"
	ColorBrightGreen   = "\033[92m"
	ColorBrightYellow  = "\033[93m"
	ColorBrightBlue    = "\033[94m"
	ColorBrightMagenta = "\033[95m"
	ColorBrightCyan    = "\033[96m"
	ColorBrightWhite   = "\033[97m"

	ResetColor = "\033[39m"
)

func ColorHex(h string) string {
	r, g, b, err := hex.ToRGB(h)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
}

func ColorRGB(r, g, b int) string {
	if (r < 0 || g < 0 || b < 0) || (r > 255 || g > 255 || b > 255) {
		return ""
	}
	return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
}

func ColorAnsi256(n int) string {
	if n < 0 || n > 255 {
		return ""
	}
	return fmt.Sprintf("\033[38;5;%dm", n)
}
