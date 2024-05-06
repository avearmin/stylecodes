package hex

import (
	"errors"
	"strconv"
)

func isHex(hex string) bool {
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

func ToRGB(hex string) (int64, int64, int64, error) {
	if !isHex(hex) {
		return 0, 0, 0, errors.New("not a hex value")
	}

	hex = hex[1:]

	r, err := strconv.ParseInt(hex[0:2], 16, 64)
	if err != nil {
		return 0, 0, 0, errors.New("couldn't parse red value")
	}
	g, err := strconv.ParseInt(hex[2:4], 16, 64)
	if err != nil {
		return 0, 0, 0, errors.New("couldn't parse green value")
	}
	b, err := strconv.ParseInt(hex[4:6], 16, 64)
	if err != nil {
		return 0, 0, 0, errors.New("couldn't parse blue value")
	}

	return r, g, b, nil
}
