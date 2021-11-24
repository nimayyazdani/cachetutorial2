package gql

import "github.com/graphql-go/graphql"

// User describes a graphql object containing a User
var Request = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Request",
		Fields: graphql.Fields{
			"Request_id": &graphql.Field{
				Type: graphql.Int,
			},
			"User_email": &graphql.Field{
				Type: graphql.String,
			},
			"Request_time": &graphql.Field{
				Type: graphql.DateTime,
			},
			"Request_type": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
