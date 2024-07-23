package stylecodes

import (
	"fmt"

	"github.com/avearmin/stylecodes/internal/hex"
)

const (
	BgColorBlack         = "\033[40m"
	BgColorRed           = "\033[41m"
	BgColorGreen         = "\033[42m"
	BgColorYellow        = "\033[43m"
	BgColorBlue          = "\033[44m"
	BgColorMagenta       = "\033[45m"
	BgColorCyan          = "\033[46m"
	BgColorWhite         = "\033[47m"
	BgColorBrightBlack   = "\033[100m"
	BgColorBrightRed     = "\033[101m"
	BgColorBrightGreen   = "\033[102m"
	BgColorBrightYellow  = "\033[103m"
	BgColorBrightBlue    = "\033[104m"
	BgColorBrightMagenta = "\033[105m"
	BgColorBrightCyan    = "\033[106m"
	BgColorBrightWhite   = "\033[107m"

	ResetBgColor = "\033[49m"
)

func BgColorHex(h string) string {
	r, g, b, err := hex.ToRGB(h)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("\033[48;2;%d;%d;%dm", r, g, b)
}

func BgColorRGB(r, g, b int) string {
	if (r < 0 || g < 0 || b < 0) || (r > 255 || g > 255 || b > 255) {
		return ""
	}
	return fmt.Sprintf("\033[48;2;%d;%d;%dm", r, g, b)
}

func BgColorAnsi256(n int) string {
	if n < 0 || n > 255 {
		return ""
	}
	return fmt.Sprintf("\033[48;5;%dm", n)
}
