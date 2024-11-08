package files

import (
	"os/exec"
	"path/filepath"
	"strings"
)

// This struct holds the data for the config.json file
type Config struct {
	BgDirectory string `json:"bgDirectory"`
	PostRun     string `json:"postRun"`
}

func (c *Config) GetPostRunScript() exec.Cmd {
	if c.PostRun == "" {
		return exec.Cmd{}
	}
	if strings.Contains(c.PostRun, "%T") {
		c.PostRun = strings.ReplaceAll(
			c.PostRun, "%T",
			filepath.Join(
				c.BgDirectory, "png_images", "current.png",
			),
		)
	}
	if strings.Contains(c.PostRun, "%t") {
		c.PostRun = strings.ReplaceAll(
			c.PostRun, "%T",
			filepath.Join(
				c.BgDirectory, "jpg_images", "current.jpg",
			),
		)
	}
	return *exec.Command(c.PostRun)
}
