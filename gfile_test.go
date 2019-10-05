package gfile_test

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/trains629/gfile"
)

func TestGraphql(t *testing.T) {
	const dir = "./"
	flist, err := ioutil.ReadDir(dir)
	if err != nil {
		t.Fatal(err.Error())
	}
	query := `
	  query {
		  readDir(path:"%s") {
			  name
			  isDir
		  }
	  }
	
	`
	data, errors := gfile.Do(fmt.Sprintf(query, dir))

	if len(errors) > 0 {
		t.Fatal(errors)
	}

	if len(flist) != len(data) {
		t.Fatal("len error")
	}
}
