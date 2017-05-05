package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra/doc"
	"go.dev.bloomberg.com/bbgo/gobas/cmd"
)

func main() {
	docsDir := filepath.Join(getRepoRoot(), "docs")
	doc.GenMarkdownTreeCustom(cmd.RootCmd, docsDir, filePrepender, linkHandler)
}

func filePrepender(filename string) string { return "---\n---\n" }

// this link fixing wouldn't be necessary if the
// jekyll-relative-links plugin would start working
func linkHandler(s string) string { return strings.TrimSuffix(s, ".md") }

func getRepoRoot() string {
	cwd, err := os.Getwd()
	check(err)

	for dir, base := filepath.Split(cwd); dir != "/"; dir, base = filepath.Split(filepath.Clean(dir)) {
		if base == "gobas" {
			return filepath.Join(dir, base)
		}
	}

	fmt.Println("This command must be run from within the gobas repo")
	os.Exit(1)
	return ""
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
