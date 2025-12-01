package lib

import (
	"os"
	"path/filepath"
)

func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}

// Read a file starting from the root directory
func ReadFile(path string) string {
	exePath, err := os.Getwd()
	CheckError(err)

	osPath := filepath.Join(filepath.Dir(exePath), path)
	dat, err := os.ReadFile(osPath)
	CheckError(err)

	return string(dat)
}
