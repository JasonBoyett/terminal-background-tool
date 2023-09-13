package images

import (
	"os"
	"path/filepath"
	"strings"

	process "github.com/sunshineplan/imgconv"
)

func SetBgImage(bg, from, dst string) error {
  err := clearDir(dst)
  if err != nil { return err }
  src, err := process.Open(filepath.Join(from, "images", bg))
  if err != nil { return err }

  if strings.HasSuffix(dst, "png_images"){
    err := process.Save(
      filepath.Join(dst, "current.png"),
      src, 
      &process.FormatOption{Format: process.PNG})
    if err != nil { return err }

  }else if strings.HasSuffix(dst, "jpg_images"){
    err := process.Save(
      filepath.Join(dst, "current.jpg"),
      src, 
      &process.FormatOption{Format: process.JPEG},
    )
    if err != nil { return err }
  }


  return nil
}

func clearDir(tgt string) error{
  err := os.RemoveAll(tgt)  
  if err != nil { return err }
  return os.MkdirAll(tgt, 0755)
}
