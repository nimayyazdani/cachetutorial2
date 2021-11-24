package gql

import (
	"github.com/Homehublol/Cache-Server"
	"github.com/graphql-go/graphql"
)

// Resolver struct holds a connection to our database
type Resolver struct {
	db *postgres.Db
}

// UserResolver resolves our user query through a db call to GetRequestByEmail
func (r *Resolver) RequestResolver(p graphql.ResolveParams) (interface{}, error) {
	// Strip the name from arguments and assert that it's a string
	user_email, ok := p.Args["user_email"].(string)
	if ok {
		requests := r.db.GetRequestsByEmail(user_email)
		return requests, nil
	}

	return nil, nil
}
