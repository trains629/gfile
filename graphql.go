package gfile

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/graphql-go/graphql"
)

var FileKindEnum = graphql.NewEnum(graphql.EnumConfig{
	Name: "FileKind",
	Values: graphql.EnumValueConfigMap{
		"DIR": &graphql.EnumValueConfig{
			Value:       1,
			Description: "dir",
		},
		"FILE": &graphql.EnumValueConfig{
			Value:       2,
			Description: "file",
		},
	},
})

type FileType struct {
	Name  string `json:"name"`
	Path  string `json:"path"`
	isDir bool   `json:"isDir"`
	Size  int64  `json:"size"`
}

var fileType = graphql.NewObject(graphql.ObjectConfig{
	Name: "File",
	Fields: graphql.Fields{
		"name":  &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
		"path":  &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
		"isDir": &graphql.Field{Type: graphql.Boolean},
		"size":  &graphql.Field{Type: graphql.Int},
	},
})

var FileInfoType = graphql.NewObject(graphql.ObjectConfig{
	Name: "FileInfo",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if value, ok := p.Source.(os.FileInfo); ok {
					return value.Name(), nil
				}
				return "", nil
			},
		},

		"isDir": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Boolean),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if value, ok := p.Source.(os.FileInfo); ok {
					return value.IsDir(), nil
				}
				return false, nil
			},
		},

		"size": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if value, ok := p.Source.(os.FileInfo); ok {
					return value.Size(), nil
				}
				return 0, nil
			},
		},
	},
})

var queryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"ls": &graphql.Field{
			Type: graphql.NewList(fileType),
			Args: graphql.FieldConfigArgument{
				"path": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				path, err := p.Args["path"].(string)
				log.Println(p.Source)
				if err {
					path = ""
				}
				ll3, count := ReadDir(path)
				log.Println(count)
				if count <= 0 {
					return []interface{}{}, nil
				}

				return func() (interface{}, error) {
					var ll2 []FileType
					log.Println(count)
					for index := 0; index < count; index++ {
						ll2 = append(ll2, <-ll3)
					}
					return ll2, nil
				}, nil
			},
		},
		"findFile": &graphql.Field{
			Type: graphql.NewList(fileType),
			Args: graphql.FieldConfigArgument{
				"path": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"exts": &graphql.ArgumentConfig{
					Type: graphql.NewList(graphql.String),
				},
				"current": &graphql.ArgumentConfig{
					Type:         graphql.Boolean,
					DefaultValue: true,
				},
			},

			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				path, e1 := p.Args["path"].(string)

				if e1 {
					path = ""
				}

				ls, err := ioutil.ReadDir(path)

				if err != nil {
					return make([]os.FileInfo, 0), nil
				}

				return ls, nil
			},
		},
		"readDir": &graphql.Field{
			Type: graphql.NewList(FileInfoType),
			Args: graphql.FieldConfigArgument{
				"path": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},

			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				path, e1 := p.Args["path"].(string)
				if !e1 {
					path = ".\\"
				}

				if ls, err := ioutil.ReadDir(path); err == nil {
					return ls, nil
				}
				return make([]os.FileInfo, 0), nil
			},
		},
	},
})

// Run query
func Run(query string) (interface{}, []error) {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})

	errors := []error{}

	if err != nil {
		return nil, errors
	}

	result := graphql.Do(graphql.Params{
		RequestString: query,
		Schema:        schema,
	})

	if result.HasErrors() {
		for _, var1 := range result.Errors {
			errors = append(errors, var1.OriginalError())
		}
		return nil, errors
	}

	return result.Data, errors
}
