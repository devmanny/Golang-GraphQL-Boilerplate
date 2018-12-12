package thing

import (
	"context"
	"time"

	"github.com/graphql-go/graphql"
)

// Thing ...
type Thing struct {
	ID        string    `json:"id" datastore:"-"`
	UserID    string    `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
	Content   string    `json:"content"`
}

//Things ...
var Things []Thing

// Ctx ...
var Ctx context.Context

// ThingType ...
var ThingType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Thing",
	Fields: graphql.Fields{
		"id":        &graphql.Field{Type: graphql.String},
		"userId":    &graphql.Field{Type: graphql.String},
		"createdAt": &graphql.Field{Type: graphql.DateTime},
		"name":      &graphql.Field{Type: graphql.String},
		"content":   &graphql.Field{Type: graphql.String},
	},
})

// ThingArgs ...
var ThingArgs = graphql.FieldConfigArgument{
	"name": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"userId": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"content": &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
}

// ThingMutations ...
var ThingMutations = graphql.Fields{
	"create": &graphql.Field{
		Description: "Create new thing",
		Type:        ThingType,
		Args:        ThingArgs,
		Resolve:     CreateThing,
	},
}

// Mutation ...
var Mutation = graphql.NewObject(graphql.ObjectConfig{
	Name:   "Mutations",
	Fields: ThingMutations,
})

// Query ...
var Query = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"things": MakeListField(MakeNodeListType("ThingList", ThingType), QueryThings),
	},
})

// Schema ...
var Schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    Query,
		Mutation: Mutation,
	},
)
