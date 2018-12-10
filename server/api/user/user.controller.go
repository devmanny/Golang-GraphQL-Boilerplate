package user

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/graphql-go/graphql"
	"github.com/labstack/echo"
	"google.golang.org/appengine/datastore"
)

// RequestData ...
type RequestData struct {
	Query     string                 `json:"query" query:"query"`
	Variables map[string]interface{} `json:"variables" query:"variables"`
}

// // DemoGraphQL ...
// func DemoGraphQL(c echo.Context) error {

// 	// body, err := ioutil.ReadAll(c.Request().Body)

// 	fields := graphql.Fields{
// 		"hello": &graphql.Field{
// 			Type: graphql.String,
// 			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
// 				fmt.Println(31, p.Args)
// 				return "world", nil
// 			},
// 		},
// 	}

// 	rootQuery := graphql.ObjectConfig{
// 		Name:   "RootQuery",
// 		Fields: fields,
// 	}

// 	schemaConfig := graphql.SchemaConfig{
// 		Query: graphql.NewObject(rootQuery),
// 	}

// 	schema, err := graphql.NewSchema(schemaConfig)

// 	request := new(RequestData)
// 	c.Bind(request)

// 	fmt.Println(request)

// 	params := graphql.Params{
// 		Schema:        schema,
// 		RequestString: request.Query,
// 	}

// 	resp := graphql.Do(params)

// 	return c.JSON(http.StatusOK, resp)
// }

// GraphQL ...
func GraphQL(c echo.Context) error {

	var (

		// UserType ...
		UserType = graphql.NewObject(graphql.ObjectConfig{
			Name: "User",
			Fields: graphql.Fields{
				"id":   &graphql.Field{Type: graphql.String},
				"name": &graphql.Field{Type: graphql.String},
			},
		})

		// UserArgs ...
		UserArgs = graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		}

		// UserMutations ...
		UserMutations = graphql.Fields{
			"createUser": &graphql.Field{
				Type:    UserType,
				Args:    UserArgs,
				Resolve: CreateUser,
			},
		}

		// Mutation ...
		Mutation = graphql.NewObject(graphql.ObjectConfig{
			Name:   "Mutations",
			Fields: UserMutations,
		})

		// Schema ...
		Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
			Mutation: Mutation,
		})
	)

	request := new(RequestData)
	c.Bind(request)

	fmt.Println(UserArgs.name)

	response := graphql.Do(graphql.Params{
		Schema:        Schema,
		RequestString: request.Query,
		// Context:       appengine.NewContext(c.Request()),
	})

	return c.JSON(http.StatusOK, response)
}

// CreateUser ...
func CreateUser(params graphql.ResolveParams) (interface{}, error) {
	ctx := params.Context

	name, _ := params.Args["name"].(string)
	user := &User{Name: name}
	key := datastore.NewIncompleteKey(ctx, "User", nil)

	fmt.Println(81, name)

	// Insert user into Datastore
	generatedKey, err := datastore.Put(ctx, key, user)

	if err != nil {
		return User{}, err
	}

	// Set user's auto-generated ID
	user.ID = strconv.FormatInt(generatedKey.IntID(), 10)
	return user, nil

}
