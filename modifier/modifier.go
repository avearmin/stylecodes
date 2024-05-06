package modifier

type modifier struct {
	Start string
	End   string
}

var (
	Bold          = modifier{Start: "\033[1m", End: "\033[22m"}
	Dim           = modifier{Start: "\033[2m", End: "\033[22m"}
	Italic        = modifier{Start: "\033[3m", End: "\033[23m"}
	Underline     = modifier{Start: "\033[4m", End: "\033[24m"}
	Inverse       = modifier{Start: "\033[7m", End: "\033[27m"}
	Hidden        = modifier{Start: "\033[8m", End: "\033[28m"}
	Strikethrough = modifier{Start: "\033[9m", End: "\033[29m"}
	End           = "\033[0m"
)
