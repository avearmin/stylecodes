package box

import (
	"github.com/avearmin/stylecodes/color"
	"github.com/avearmin/stylecodes/bgcolor"
	"github.com/avearmin/stylecodes/modifier"
	"testing"
)

func TestStripANSI(t *testing.T) {
	tests := map[string]struct{
		input string
		want string
	} {
		"strip color": {
			input: color.Red + "foo" + color.End,
			want: "foo",
		},
		"strip bgcolor": {
			input: bgcolor.Red + "foo" + bgcolor.End,
			want: "foo",
		},
		"strip modifier": {
			input: modifier.Underline + "foo" + modifier.EndUnderline,
			want: "foo",
		},
		"strip hex": {
			input: color.Hex("#ffffff") + "foo" + color.End,
			want: "foo",
		},
		"strip Ansi256": {
			input: bgcolor.Ansi256(26) + "foo" + bgcolor.End,
			want: "foo",
		},
	}

	for k, v := range tests {
		t.Run(k, func(t *testing.T){
			
			if got := stripANSI(v.input); got != v.want {
				t.Fatalf("want=%s, but got=%s", v.want, got)
			}
		})
	}
}
