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
			input: Color.Red + "foo" + Color.Reset,
			want:  "foo",
		},
		"strip bgcolor": {
			input: BgColor.Red + "foo" + BgColor.Reset,
			want:  "foo",
		},
		"strip modifier": {
			input: Modifier.Underline + "foo" + Modifier.ResetUnderline,
			want:  "foo",
		},
		"strip hex": {
			input: Color.Hex("#ffffff") + "foo" + Color.Reset,
			want:  "foo",
		},
		"strip Ansi256": {
			input: BgColor.Ansi256(26) + "foo" + BgColor.Reset,
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
