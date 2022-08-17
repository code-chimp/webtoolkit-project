package main

import "github.com/code-chimp/webtoolkit"

func main() {
	var tools webtoolkit.Tools
	testDir := "./test-parent/test-dir"

	_ = tools.CreateDirIfNotExists(testDir)
}
