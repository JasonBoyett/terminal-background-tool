package tui

import (
	"bufio"
	"fmt"
	"os"

	files "github.com/JasonBoyett/terminal-background-tool/internal/files"
)

func Setup() error {

	var path string
	var postRun string

	fmt.Println("Let's get your image folder set up")
	fmt.Println("Where would you like your image folder to be?")
	fmt.Println("Please provide a path")

	reader := bufio.NewScanner(os.Stdin)
	if reader.Scan() {
		path = reader.Text()
	}
	postRun = setupPostRun()
	if err := files.SaveConfig(path, postRun); err != nil {
		fmt.Println("error in save")
		return err
	}

	return nil
}

func setupPostRun() string {
	var answer string
	var postRun string
	fmt.Println("Would you like to run a script after the image changes? (y/n)")
	reader := bufio.NewScanner(os.Stdin)
	if reader.Scan() {
		answer = reader.Text()
	}
	if answer == "y" {
		fmt.Println(
			`Please provide the script to run. 
      Use %T to represent the png image file or %t to represent the jpg image file.`,
		)
		if reader.Scan() {
			postRun = reader.Text()
		}
	}
	return postRun
}
