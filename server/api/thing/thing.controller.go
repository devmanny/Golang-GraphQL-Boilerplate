package thing

import (
	"net/http"
	"strconv"
	"time"

	"github.com/graphql-go/graphql"
	"github.com/labstack/echo"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

var thingType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Thing",
	Fields: graphql.Fields{
		"id":        &graphql.Field{Type: graphql.String},
		"userId":    &graphql.Field{Type: graphql.String},
		"createdAt": &graphql.Field{Type: graphql.DateTime},
		"content":   &graphql.Field{Type: graphql.String},
	},
})

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

// CreateThing ...
func CreateThing(params graphql.ResolveParams) (interface{}, error) {

	// Get arguments
	name, _ := params.Args["content"].(string)
	content, _ := params.Args["content"].(string)
	userID, _ := params.Args["userId"].(string)

	thing := &Thing{
		UserID:    userID,
		Name:      name,
		Content:   content,
		CreatedAt: time.Now().UTC(),
	}

	key := datastore.NewIncompleteKey(Ctx, "Thing", nil)
	generatedKey, err := datastore.Put(Ctx, key, thing)

	if err != nil {
		return Thing{}, err
	}

	// Update thing's ID
	thing.ID = strconv.FormatInt(generatedKey.IntID(), 10)

	return thing, nil
}
