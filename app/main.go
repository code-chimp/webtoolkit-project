package main

import (
	"fmt"
	"github.com/code-chimp/webtoolkit"
)

func main() {
	var tools webtoolkit.Tools
	s := tools.RandomString(10)
	fmt.Println("Random string:", s)

	s = tools.RandomString(25)
	fmt.Println("Random string:", s)
}
