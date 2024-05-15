package main

import (
	"fmt"
	"github.com/avearmin/stylecodes/color"
	"github.com/avearmin/stylecodes/box"
)

func main() {
	fmt.Println(box.Border(color.Red + "this" + color.End + "\n" + color.Red + "has" + color.End + "\n" + color.Red + "a" + color.End + "\n" + color.Red + "border" + color.End, color.Gold))
}
