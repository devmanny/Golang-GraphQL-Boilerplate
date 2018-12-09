package thing

import "time"

// Thing ...
type Thing struct {
	ID        string    `json:"id" datastore:"-"`
	UserID    string    `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
	Content   string    `json:"content"`
}
