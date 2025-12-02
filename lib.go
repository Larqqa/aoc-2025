package lib

import (
	"os"
	"path/filepath"
	"runtime"
)

func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}

// Get the adjacent directory to lib, which should be root
func getRootDir() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Dir(filename)
}

// Read a file starting from the root directory
func ReadFile(path string) string {
	root := getRootDir()
	osPath := filepath.Join(root, path)
	dat, err := os.ReadFile(osPath)
	CheckError(err)

	return string(dat)
}
