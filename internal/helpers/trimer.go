package helpers

import "errors"

func TrimFileExtension(filename string) (string, error) {
	for i := len(filename) - 1; i >= 0; i-- {
		if filename[i] == '.' {
			return filename[:i], nil
		}
	}

	return filename, errors.New("No file extension found")
}
