package user

import (
	"context"

	"github.com/graphql-go/graphql"
)

// User ...
type User struct {
	ID   string `json:"id" datastore:"-"`
	Name string `json:"name"`
}

//Users ...
var Users []User

// Ctx ...
var Ctx context.Context

// UserType ...
var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"name": &graphql.Field{Type: graphql.String},
	},
})

// UserArgs ...
var UserArgs = graphql.FieldConfigArgument{
	"name": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
}

// UserMutations ...
var UserMutations = graphql.Fields{
	"create": &graphql.Field{
		Description: "Create new user",
		Type:        UserType,
		Args:        UserArgs,
		Resolve:     CreateUser,
	},
}

// Mutation ...
var Mutation = graphql.NewObject(graphql.ObjectConfig{
	Name:   "Mutations",
	Fields: UserMutations,
})

// Query ...
var Query = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
	},
})
