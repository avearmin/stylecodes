package color

import (
	"fmt"
	"strconv"
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

func IsHex(hex string) bool {
	if len(hex) != 7 {
		return false
	}
	if hex[0] != '#' {
		return false
	}

	hex = hex[1:]

	for i := 0; i < len(hex); i++ {
		if !isSupportedChar(hex[i]) {
			return false
		}
	}

	return true
}
func isSupportedChar(b byte) bool {
	return ('A' <= b && b <= 'F') || ('a' <= b && b <= 'f') || ('0' <= b && b <= '9')
}

func Hex(hex string) string {
	if !IsHex(hex) {
		return ""
	}

	hex = hex[1:]

	r, err := strconv.ParseInt(hex[0:2], 16, 64)
	if err != nil {
		return ""
	}
	g, err := strconv.ParseInt(hex[2:4], 16, 64)
	if err != nil {
		return ""
	}
	b, err := strconv.ParseInt(hex[4:6], 16, 64)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
}
