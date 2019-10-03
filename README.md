# gfile
使用graphql进行文件或目录查询

graphql定义的结构:

```graphql

type File {
    name:String! # 文件名称,包含后缀
    path:String! # 文件路径，默认从/开始
    isDir:Boolean
    size:Int
}

type Time {

}

type FileInfo {
  Name():String!
  Size(): Int
  Mode(): Int
  ModTime(): Time
  IsDir(): Boolean
}

type Query {
  readDir(path:String!):[FileInfo]!
  ls(path:String):[File]! # 返回指定目录下的文件列表
  findFile(path:!String,exts:[String],current:Boolean):[File]! # 按后缀名查找文件，默认只查找当前目录
  lsByFile(file:File!):[File]!
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
