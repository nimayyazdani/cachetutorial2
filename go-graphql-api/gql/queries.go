package gql

import (
	"github.com/Homehublol/Cache-Server" // how to change this to our local computer //
	"github.com/graphql-go/graphql" 
)

// Root holds a pointer to a graphql object
type Root struct {
	Query *graphql.Object
}

// NewRoot returns base query type. This is where we add all the base queries
func NewRoot(db *postgres.Db) *Root {
	// Create a resolver holding our database. Resolver can be found in resolvers.go
	resolver := Resolver{db: db}

	// Create a new Root that describes our base query set up. In this
	// example we have a request query that takes one argument called user_email
	root := Root{
		Query: graphql.NewObject(
			graphql.ObjectConfig{
				Name: "Query",
				Fields: graphql.Fields{
					"Request": &graphql.Field{
						// Slice of Request type which can be found in types.go
						Type: graphql.NewList(Request),
						Args: graphql.FieldConfigArgument{
							"User_email": &graphql.ArgumentConfig{
								Type: graphql.String,
							},
						},
						Resolve: resolver.RequestResolver,
					},
				},
			},
		),
	}
	return &root
}
