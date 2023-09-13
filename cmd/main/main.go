package main

import (
	"os"

	files "github.com/JasonBoyett/terminal-background-tool/internal/files"
	selector "github.com/JasonBoyett/terminal-background-tool/internal/selector"
)


func main() {
	opts, err := files.GetValidOpts()
	if err != nil {
    if err.Error() == "No config.json file found." {
      err := selector.Setup()
      os.Exit(0)
      if err != nil { panic(err) }
    }
	} else {
		selector.EnterTui(opts)
	}
}
