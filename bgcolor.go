package stylecodes

import (
	"fmt"

	"github.com/avearmin/stylecodes/internal/hex"
)

type bgColor struct {
	Black         string
	Red           string
	Green         string
	Yellow        string
	Blue          string
	Magenta       string
	Cyan          string
	White         string
	BrightBlack   string
	BrightRed     string
	BrightGreen   string
	BrightYellow  string
	BrightBlue    string
	BrightMagenta string
	BrightCyan    string
	BrightWhite   string
	Reset         string
}

var BgColor = bgColor{
	Black:         "\033[40m",
	Red:           "\033[41m",
	Green:         "\033[42m",
	Yellow:        "\033[43m",
	Blue:          "\033[44m",
	Magenta:       "\033[45m",
	Cyan:          "\033[46m",
	White:         "\033[47m",
	BrightBlack:   "\033[100m",
	BrightRed:     "\033[101m",
	BrightGreen:   "\033[102m",
	BrightYellow:  "\033[103m",
	BrightBlue:    "\033[104m",
	BrightMagenta: "\033[105m",
	BrightCyan:    "\033[106m",
	BrightWhite:   "\033[107m",
	Reset:         "\033[49m",
}

func (bg bgColor) Hex(h string) string {
	r, g, b, err := hex.ToRGB(h)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("\033[48;2;%d;%d;%dm", r, g, b)
}

func (bg bgColor) RGB(r, g, b int) string {
	if (r < 0 || g < 0 || b < 0) || (r > 255 || g > 255 || b > 255) {
		return ""
	}
	return fmt.Sprintf("\033[48;2;%d;%d;%dm", r, g, b)
}

func (bg bgColor) Ansi256(n int) string {
	if n < 0 || n > 255 {
		return ""
	}
	return fmt.Sprintf("\033[48;5;%dm", n)
}
