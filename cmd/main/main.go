package main

import (
	files "github.com/JasonBoyett/terminal-background-tool/internal/files"
	selector "github.com/JasonBoyett/terminal-background-tool/internal/selector"
)


func main() {
	opts, err := files.GetValidOpts()
	if err != nil {
    if err.Error() == "No config.json file found." {
      err := selector.Setup()
      if err != nil { panic(err) }
      main()
    }
	} else {
		selector.EnterTui(opts)
	}
}
