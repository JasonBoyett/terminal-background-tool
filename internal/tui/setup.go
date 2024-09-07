package tui

import (
	"bufio"
	"fmt"
	"os"

	files "github.com/JasonBoyett/terminal-background-tool/internal/files"
)

func Setup() error {

	var path string

	fmt.Println("Let's get your image folder set up")
	fmt.Println("Where would you like your image folder to be?")
	fmt.Println("Please provide a path")

	reader := bufio.NewScanner(os.Stdin)
	if reader.Scan() {
		path = reader.Text()
	}
	if err := files.SaveConfig(path); err != nil {
		fmt.Println("error in save")
		return err
	}

	return nil
}
