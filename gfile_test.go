package gfile_test

import (
	"fmt"
	"testing"

	"github.com/trains629/gfile"
)

func TestGraphql(t *testing.T) {
	const dir = "F://"
	query := `
	  query {
		  readDir(path:"%s") {
			  name
			  isDir
		  }
	  }
	
	`
	data, errors := gfile.Run(fmt.Sprintf(query, dir))

	if len(errors) > 0 {
		t.Fatal(errors)
	}

	t.Log(data)
}
