package user

// User ...
type User struct {
	ID   string `json:"id" datastore:"-"`
	Name string `json:"name"`
}
