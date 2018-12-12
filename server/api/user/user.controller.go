package user

import (
	"errors"
	"net/http"
	"strconv"

	"app.onca.api/server/api/thing"

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

// QueryUser ...
func QueryUser(params graphql.ResolveParams) (interface{}, error) {
	if strID, ok := params.Args["id"].(string); ok {
		// Parse ID argument
		id, errID := strconv.ParseInt(strID, 10, 64)
		if errID != nil {
			return nil, errors.New("Invalid id")
		}
		user := &User{ID: strID}
		key := datastore.NewKey(Ctx, "User", "", id, nil)
		// Fetch user by ID

		err := datastore.Get(Ctx, key, user)
		if err != nil {
			// Assume not found
			return nil, errors.New("User not found")
		}
		return user, nil
	}
	return User{}, nil
}

// QueryThingsByUser ...
func QueryThingsByUser(params graphql.ResolveParams) (interface{}, error) {
	query := datastore.NewQuery("Thing")
	if limit, ok := params.Args["limit"].(int); ok {
		query = query.Limit(limit)
	}
	if offset, ok := params.Args["offset"].(int); ok {
		query = query.Offset(offset)
	}
	// Check user's ID against post's UserID field

	userData, ok := params.Source.(*User)
	if ok {
		query = query.Filter("UserID =", userData.ID)
	}
	return thing.QueryThingList(Ctx, query)
}
