package user

import "github.com/graphql-go/graphql"

// User ...
type User struct {
	ID   string `json:"id" datastore:"-"`
	Name string `json:"name"`
}

// UserType ...
var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id":   &graphql.Field{Type: graphql.String},
		"name": &graphql.Field{Type: graphql.String},
	},
})
