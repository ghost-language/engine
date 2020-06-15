package engine

import (
	"os"
	"os/exec"
	"path"
)

// GetPath returns the current working directory.
func GetPath() string {
	file, _ := exec.LookPath(os.Args[0])
	dir, _ := path.Split(file)

	os.Chdir(dir)

	path, _ := os.Getwd()

	return path + "/"
}

func FileExists(_file string) bool {
	info, err := os.Stat(_file)

	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}
