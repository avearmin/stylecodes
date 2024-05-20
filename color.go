package stylecodes

import (
	"fmt"

	"github.com/avearmin/stylecodes/internal/hex"
)

type color struct {
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

var Color = color{
	Black:         "\033[30m",
	Red:           "\033[31m",
	Green:         "\033[32m",
	Yellow:        "\033[33m",
	Blue:          "\033[34m",
	Magenta:       "\033[35m",
	Cyan:          "\033[36m",
	White:         "\033[37m",
	BrightBlack:   "\033[90m",
	BrightRed:     "\033[91m",
	BrightGreen:   "\033[92m",
	BrightYellow:  "\033[93m",
	BrightBlue:    "\033[94m",
	BrightMagenta: "\033[95m",
	BrightCyan:    "\033[96m",
	BrightWhite:   "\033[97m",
	Reset:         "\033[39m",
}

func (c color) Hex(h string) string {
	r, g, b, err := hex.ToRGB(h)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
}

func (c color) RGB(r, g, b int) string {
	if (r < 0 || g < 0 || b < 0) || (r > 255 || g > 255 || b > 255) {
		return ""
	}
	return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
}

func (c color) Ansi256(n int) string {
	if n < 0 || n > 255 {
		return ""
	}
	return fmt.Sprintf("\033[38;5;%dm", n)
}
