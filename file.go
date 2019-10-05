package gfile

import (
	"os"
	"time"
)

type TFileInfo struct {
	File os.FileInfo
	Path string `json:path`
}

func (f TFileInfo) Name() string {
	return f.File.Name()
}

func (f TFileInfo) Size() int64 {
	return f.File.Size()
}

func (f TFileInfo) Mode() os.FileMode {
	return f.File.Mode()
}

func (f TFileInfo) ModTime() time.Time {
	return f.File.ModTime()
}

func (f TFileInfo) IsDir() bool {
	return f.File.IsDir()
}

func (f TFileInfo) Sys() interface{} {
	return f.File.Sys()
}

// Exists func
func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
