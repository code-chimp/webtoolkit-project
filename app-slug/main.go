package main

import (
	"fmt"
	"github.com/code-chimp/webtoolkit"
	"log"
)

func main() {
	var tools webtoolkit.Tools

	slug, err := tools.Slugify("JUMPin`_jehosephat&hodor^jimmy")
	if err != nil {
		log.Println(err.Error())
	}

	fmt.Println(slug)
}
