# gfile
使用graphql进行文件或目录查询

graphql定义的结构:

```graphql

type Time {

}

type FileInfo {
  Name():String!
  Size(): Int
  Mode(): Int
  ModTime(): Time
  IsDir(): Boolean
  path: String!
}

type Query {
  readDir(path:String!):[FileInfo]!
  exists(path:String):Boolean!
  findFile(path:!String,exts:[String],current:Boolean):[FileInfo]!
}

```

演示范例:

```go
package main

import (
	"github.com/trains629/gfile"
)

func main() {
	query := `
	  query {
		  readDir(path:"F:\\") {
			  name
			  isDir
		  }
	  }
	
	`
	gfile.Run(query)
}
```
