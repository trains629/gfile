package gfile

import (
	"io/ioutil"

	"github.com/graphql-go/graphql"
)

var FileInfoType = graphql.NewObject(graphql.ObjectConfig{
	Name: "FileInfo",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if value, ok := p.Source.(*TFileInfo); ok {
					return value.Name(), nil
				}
				return "", nil
			},
		},

		"isDir": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Boolean),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if value, ok := p.Source.(*TFileInfo); ok {
					return value.IsDir(), nil
				}
				return false, nil
			},
		},

		"size": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if value, ok := p.Source.(*TFileInfo); ok {
					return value.Size(), nil
				}
				return 0, nil
			},
		},

		"path": &graphql.Field{
			Type: graphql.String,
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
					return make([]*TFileInfo, 0), nil
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
				path, ok := p.Args["path"].(string)
				if !ok {
					return make([]*TFileInfo, 0), nil
				}
				ls, err := ioutil.ReadDir(path)
				if err != nil {
					return make([]*TFileInfo, 0), nil
				}
				ls2 := make([]*TFileInfo, len(ls))
				for ii, var1 := range ls {
					ls2[ii] = &TFileInfo{
						File: var1, Path: path,
					}
				}
				return ls2, nil
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
func Do(query string) (interface{}, []error) {
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
