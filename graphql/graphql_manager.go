package graphql

import "github.com/graphql-go/graphql"

type GraphqlSchema struct{
	Scheam *graphql.Schema
	RootQuery *graphql.Object
	RootMutation *graphql.Object
}

func GetGraphqlSchema() *GraphqlSchema{
	schema:=&GraphqlSchema{}
	// root mutation
	schema.RootMutation = graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Description:"This is RootMutation",
		Fields: graphql.Fields{
		},
	})
	schema.RootQuery = graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Description:"This is RootQuery",
		Fields: graphql.Fields{
		},
		})

	WrapTodoQuerys(schema)

	// define schema, with our rootQuery and rootMutation
	graphqlSchema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query:    schema.RootQuery,
		Mutation: schema.RootMutation,
	})
	schema.Scheam=&graphqlSchema
	return schema

}