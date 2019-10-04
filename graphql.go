package gfile

import (
	"io/ioutil"
	"os"

	"github.com/graphql-go/graphql"
)

type TFileInfo struct {
	File os.FileInfo
	Path string
}

var FileInfoType = graphql.NewObject(graphql.ObjectConfig{
	Name: "FileInfo",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if value, ok := p.Source.(TFileInfo); ok {
					return value.File.Name(), nil
				}
				return "", nil
			},
		},

		"isDir": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Boolean),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if value, ok := p.Source.(TFileInfo); ok {
					return value.File.IsDir(), nil
				}
				return false, nil
			},
		},

		"size": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if value, ok := p.Source.(TFileInfo); ok {
					return value.File.Size(), nil
				}
				return 0, nil
			},
		},
	},
})

var queryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"findFile": &graphql.Field{
			Type: graphql.NewList(FileInfoType),
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
					return make([]TFileInfo, 0), nil
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
					ls2 := make([]TFileInfo, len(ls))
					for _, var1 := range ls {
						ls2 = append(ls2, TFileInfo{
							File: var1, Path: path,
						})
					}
					return ls2, nil
				}
				return make([]TFileInfo, 0), nil
			},
		},
		"exists": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Boolean),
			Args: graphql.FieldConfigArgument{
				"path": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},

			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				path, _ := p.Args["path"].(string)
				return Exists(path), nil
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
