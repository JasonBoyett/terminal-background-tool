package main

import (
	files "github.com/JasonBoyett/terminal-background-tool/internal/files"
	tui "github.com/JasonBoyett/terminal-background-tool/internal/tui"
)

func main() {
	opts, err := files.GetValidOpts()
	if err != nil {
		if err.Error() == "No config.json file found." {
			err := tui.Setup()
			if err != nil {
				panic(err)
			}
			main()
		}
	} else {
		tui.EnterTui(opts)
	}
}
