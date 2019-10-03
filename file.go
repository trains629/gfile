package gfile

import (
	"io/ioutil"
	"os"
)

// Exists func
func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func ReadDir(path string) (chan FileType, int) {
	if !Exists(path) {
		return nil, 0
	}
	ll, err2 := ioutil.ReadDir(path)

	if err2 != nil {
		return nil, 0
	}
	ch := make(chan FileType, len(ll))

	for _, var1 := range ll {
		go func(var1 os.FileInfo) {
			ch <- FileType{
				Name:  var1.Name(),
				Path:  path,
				isDir: var1.IsDir(),
				Size:  var1.Size(),
			}
		}(var1)
	}

	return ch, len(ll)

}
