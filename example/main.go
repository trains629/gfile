package main

import (
	"fmt"
	"log"

	"github.com/trains629/gfile"
)

func main() {
	query := `
	  query {
		  readDir(path:"./") {
			  name
			  isDir
			  path
		  }
	  }
	
	`
	data, e := gfile.Do(query)
	if len(e) > 0 {
		fmt.Println(e)
	} else {
		fmt.Println(data)
	}

	value, ok := data.(map[string]interface{})

	if !ok {
		return
	}

	readdir := value["readDir"]

	t35, ok := readdir.([]interface{})

	if ok {
		log.Println(t35, len(t35))

		for _, var1 := range t35 {
			log.Println(var1)
			t43, ok := var1.(map[string]interface{})
			if ok {
				log.Println(t43["name"])
			}
		}
	}
}
