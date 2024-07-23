package stylecodes

import (
	"testing"
)

func TestStripANSI(t *testing.T) {
	tests := map[string]struct {
		input string
		want  string
	}{
		"strip color": {
			input: ColorRed + "foo" + ResetColor,
			want:  "foo",
		},
		"strip bgcolor": {
			input: BgColorRed + "foo" + ResetBgColor,
			want:  "foo",
		},
		"strip modifier": {
			input: Underline + "foo" + ResetUnderline,
			want:  "foo",
		},
		"strip hex": {
			input: ColorHex("#ffffff") + "foo" + ResetColor,
			want:  "foo",
		},
		"strip Ansi256": {
			input: BgColorAnsi256(26) + "foo" + ResetBgColor,
			want:  "foo",
		},
	}

	for k, v := range tests {
		t.Run(k, func(t *testing.T) {

			if got := StripANSI(v.input); got != v.want {
				t.Fatalf("want=%s, but got=%s", v.want, got)
			}
		})
	}
}
