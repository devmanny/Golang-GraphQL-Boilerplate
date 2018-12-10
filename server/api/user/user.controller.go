package user

import (
	"net/http"
	"strconv"

	"github.com/graphql-go/graphql"
	"github.com/labstack/echo"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

// RequestData ...
type RequestData struct {
	Query     string                 `json:"query" query:"query"`
	Variables map[string]interface{} `json:"variables" query:"variables"`
}

// GraphQL ...
func GraphQL(c echo.Context) error {

	// Schema ...
	Schema, _ := graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    Query,
			Mutation: Mutation,
		},
	)

	Ctx = appengine.NewContext(c.Request())

	request := new(RequestData)
	c.Bind(request)

	response := graphql.Do(
		graphql.Params{
			Schema:        Schema,
			RequestString: request.Query,
		},
	)

	return c.JSON(http.StatusOK, response)
}

// CreateUser ...
func CreateUser(params graphql.ResolveParams) (interface{}, error) {

	user := &User{
		Name: params.Args["name"].(string),
	}

	key := datastore.NewIncompleteKey(Ctx, "User", nil)
	generatedKey, err := datastore.Put(Ctx, key, user)

	if err != nil {
		return User{}, err
	}

	user.ID = strconv.FormatInt(generatedKey.IntID(), 10)
	return user, nil
}
