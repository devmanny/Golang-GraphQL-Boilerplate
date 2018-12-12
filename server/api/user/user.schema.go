package user

import (
	"context"

	"app.onca.api/server/api/thing"
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
		"id":     &graphql.Field{Type: graphql.String},
		"name":   &graphql.Field{Type: graphql.String},
		"things": thing.MakeListField(thing.MakeNodeListType("ThingList", thing.ThingType), QueryThingsByUser),
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
		"user": &graphql.Field{
			Type: UserType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			},
			Resolve: QueryUser,
		},
	},
})

// Schema ...
var Schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    Query,
		Mutation: Mutation,
	},
)
