package thing

import (
	"context"
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

// MakeListField ...
func MakeListField(listType graphql.Output, resolve graphql.FieldResolveFn) *graphql.Field {
	return &graphql.Field{
		Type:    listType,
		Resolve: resolve,
		Args: graphql.FieldConfigArgument{
			"limit":  &graphql.ArgumentConfig{Type: graphql.Int},
			"offset": &graphql.ArgumentConfig{Type: graphql.Int},
		},
	}
}

// MakeNodeListType ...
func MakeNodeListType(name string, nodeType *graphql.Object) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: name,
		Fields: graphql.Fields{
			"nodes":      &graphql.Field{Type: graphql.NewList(nodeType)},
			"totalCount": &graphql.Field{Type: graphql.Int},
		},
	})
}

// ListResult ...
type ListResult struct {
	Nodes      []Thing `json:"nodes"`
	TotalCount int     `json:"totalCount"`
}

// QueryThingList ...
func QueryThingList(ctx context.Context, query *datastore.Query) (ListResult, error) {
	// Order by creation time
	query = query.Order("-CreatedAt")
	var result ListResult
	// Run the query
	if keys, err := query.GetAll(ctx, &result.Nodes); err != nil {
		return result, err
	} else {
		// Set IDs
		for i, key := range keys {
			result.Nodes[i].ID = strconv.FormatInt(key.IntID(), 10)
		}
		// Set total count
		result.TotalCount = len(result.Nodes)
	}
	return result, nil
}

// QueryThings ...
func QueryThings(params graphql.ResolveParams) (interface{}, error) {
	ctx := params.Context
	query := datastore.NewQuery("Thing")
	if limit, ok := params.Args["limit"].(int); ok {
		query = query.Limit(limit)
	}
	if offset, ok := params.Args["offset"].(int); ok {
		query = query.Offset(offset)
	}
	return QueryThingList(ctx, query)
}
