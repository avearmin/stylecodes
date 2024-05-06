package color

import "testing"

func TestIsHex(t *testing.T) {
	tests := map[string]struct {
		input string
		want  bool
	}{
		"valid hex": {
			input: "#36FF00",
			want:  true,
		},
		"invalid hex: unsupported letter": {
			input: "#36fZ00",
			want:  false,
		},
		"invalid hex: too little chars": {
			input: "#36FF0",
			want:  false,
		},
		"invalid hex: no #": {
			input: "36FF000",
			want:  false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if got := IsHex(test.input); got != test.want {
				t.Fatalf("got=%t, but want=%t", got, test.want)
			}
		})
	}
}
