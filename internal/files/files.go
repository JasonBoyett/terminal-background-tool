package files

import (
	"encoding/json"
	"errors"
	"strings"

	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"

	images "github.com/JasonBoyett/terminal-background-tool/internal/images"
)

// This struct holds the data for the config.json file
type Config struct {
  BgDirectory string `json:"bgDirectory"`
}



func SaveConfig(data string) error{
  root, err := os.Executable()
  if err != nil { return err }
  root = filepath.Clean(root)
  root = strings.Trim(root, "main")
  fileName := filepath.Join(root, "config.json") 
  config := Config{BgDirectory: data}
  configData, err := json.MarshalIndent(config, "", "  ")
  if err != nil {
    return err
  }
  if err := ioutil.WriteFile(fileName, configData, 0644); err != nil {
    return err
  }
  return nil
}

func LoadConfig() (Config, error){
  var config Config
  root, err:= os.Executable()
  root = filepath.Clean(root)
  root = strings.Trim(root, "main")
  if err != nil { 
    return config, errors.New("No config.json file found.") 
  }
  configFile := filepath.Join(root, "config.json")
  data, error := ioutil.ReadFile(configFile)
  if error != nil {
    return config, errors.New("No config.json file found.") 
  }

  if err := json.Unmarshal(data, &config); err != nil {
    return config, err
  }
  
  return config, nil
}


// GetValidOpts returns a list of options for the user to choose from.
func GetValidOpts() ([]string, error) {
	opts := []string{}

	config, err := LoadConfig()
	if err != nil {
		if err.Error() == "open config.json: no such file or directory" {
			return nil, errors.New("No config.json file found.")
		}

		return nil, err
	}

	imagesDir := filepath.Join(config.BgDirectory, "images")

	files, err := ioutil.ReadDir(imagesDir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if isImage(file.Name()){
		  opts = append(opts, file.Name())
		}
	}

	return opts, nil
}


// SetBg sets the background image to the given file in the directory provided by the config file.
func SetBg(bg string) error{
  config, err := LoadConfig()
  if err != nil {
    return err
  }
  
  pngPath := filepath.Join(config.BgDirectory, "png_images")
  jpgPath := filepath.Join(config.BgDirectory, "jpg_images")

  paths := []string{pngPath, jpgPath}

  for _, path := range paths {
    if err := os.MkdirAll(path, 0755); err != nil { panic(err) }

    if err := images.SetBgImage(
      bg,
      config.BgDirectory, 
      path,
    ); err != nil { return err }
  }


  return nil
}

// RandomBg sets the background image to a random image in the users image directory.
func RandomBg() error{
  choices, err := GetValidOpts()
  if err != nil { return err }
  
  num := rand.Intn(len(choices))

  bgErr := SetBg(choices[num])
  if bgErr != nil { return err }

  return nil
}

func isImage(filename string) bool{
  inputExt := filepath.Ext(filename)
  imageExtensions := []string{
    ".jpg",
    ".jpeg",
    ".png",
    ".gif",
    ".bmp",
    ".tiff",
    ".webp",
    ".svg",
    ".ico",
  }
  for _, ext := range imageExtensions {
    if inputExt == ext {
      return true
    }
  }
  return false
}
