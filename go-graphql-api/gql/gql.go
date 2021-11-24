package gql

import (
	"fmt"

	"github.com/graphql-go/graphql"
)

// ExecuteQuery runs our graphql queries
func ExecuteQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	// Error check
	if len(result.Errors) > 0 {
		fmt.Printf("Unexpected errors inside ExecuteQuery: %v", result.Errors)
	}

	return result
}
// do we have to add a ExecuteMutation() function to
// handle GraphQL mutations 

// how do we run realize init, and what does this do?