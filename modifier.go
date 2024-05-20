package stylecodes

type modifier struct {
	Bold               string
	Dim                string
	Italic             string
	Underline          string
	Inverse            string
	Hidden             string
	Strikethrough      string
	Reset              string
	ResetBold          string
	ResetDim           string
	ResetItalic        string
	ResetUnderline     string
	ResetInverse       string
	ResetHidden        string
	ResetStrikethrough string
}

var Modifier = modifier{
	Bold:               "\033[1m",
	Dim:                "\033[2m",
	Italic:             "\033[3m",
	Underline:          "\033[4m",
	Inverse:            "\033[7m",
	Hidden:             "\033[8m",
	Strikethrough:      "\033[9m",
	Reset:              "\033[0m",
	ResetBold:          "\033[22m",
	ResetDim:           "\033[22m",
	ResetItalic:        "\033[23m",
	ResetUnderline:     "\033[24m",
	ResetInverse:       "\033[27m",
	ResetHidden:        "\033[28m",
	ResetStrikethrough: "\033[29m",
}
