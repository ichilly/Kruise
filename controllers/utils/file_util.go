package utils

import (
	"os"
	"path"
)

func GetFilePath(dir string, name string) string {
	filePath := path.Join(dir, name)
	_, err := os.Stat(filePath)
	if err != nil {
		return ""
	}
	return filePath
}