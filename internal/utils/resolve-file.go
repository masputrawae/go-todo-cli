package utils

import (
	"os"
	"path/filepath"
)

func ResolveFile(fp string) (*os.File, error) {
	dir := filepath.Dir(fp)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}
	return os.OpenFile(fp, os.O_RDWR|os.O_CREATE, 0644)
}
