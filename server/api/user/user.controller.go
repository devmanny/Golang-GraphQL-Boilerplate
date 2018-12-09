package user

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/graphql-go/graphql"
	"github.com/labstack/echo"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

// var schema graphql.Schema

// Init ...
func Init() {
	fmt.Println("INIT - - - -")

}

// Index ...
func Index(c echo.Context) error {

	// rootMutation ...
	rootMutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "rootMutation",
		Fields: graphql.Fields{
			"createUser": &graphql.Field{
				Type: UserType,
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: createUser,
			},
		},
	})

	fmt.Println(rootMutation)

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Mutation: rootMutation,
	})

	fmt.Println(err)

	ctx := appengine.NewContext(c.Request())
	body, err := ioutil.ReadAll(c.Request().Body)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	query := string(body)

	fmt.Println(schema, query, ctx)

	resp := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	if len(resp.Errors) > 0 {
		return c.JSON(http.StatusBadRequest, resp.Errors)
	}

	return c.JSON(http.StatusOK, resp)
}

func createUser(params graphql.ResolveParams) (interface{}, error) {
	ctx := params.Context

	name, _ := params.Args["name"].(string)
	user := &User{Name: name}
	key := datastore.NewIncompleteKey(ctx, "User", nil)

	// Insert user into Datastore
	if generatedKey, err := datastore.Put(ctx, key, user); err != nil {
		return User{}, err
	} else {
		// Set user's auto-generated ID
		user.ID = strconv.FormatInt(generatedKey.IntID(), 10)
		return user, nil
	}
}
