package box

import (
	"regexp"
	"strings"
)

const (
	topLeft = "┌"
	topRight = "┐"
	botLeft = "└"
	botRight = "┘"
	mid = "─"
	side = "│"

	endStyle = "\033[0m"
)

func Border(text, style string) string {
	lines := strings.Split(text, "\n")
	
	maxLen := 0
	for _, v := range lines {
		if maxLen < len([]rune(stripANSI(v))) {
			maxLen = len([]rune(stripANSI(v))) 
		}
	}

	topBorder := style + topLeft + strings.Repeat(mid, maxLen) + topRight + endStyle
	botBorder := style + botLeft + strings.Repeat(mid, maxLen) + botRight + endStyle
	
	var builder strings.Builder
	builder.WriteString(topBorder + "\n")
	builder.WriteString(addSide(lines, maxLen, style))
	builder.WriteString(botBorder)

	return builder.String() 
}

func addSide(lines []string, maxLen int, style string) string {
	var text strings.Builder

	for _, v := range lines {
		padding := strings.Repeat(" ", maxLen - len([]rune(stripANSI(v))))
		styledSide := style + side + endStyle
		text.WriteString(styledSide + v + padding + styledSide + "\n")
	}
	
	return text.String()
}

func stripANSI(s string) string {
	re := regexp.MustCompile("[\u001B\u009B][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[a-zA-Z\\d]*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PRZcf-ntqry=><~]))")
	return re.ReplaceAllString(s, "")
}
