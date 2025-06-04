package order

import "time"

var (
	NewStatus        = 0
	ProcessingStatus = 1
	InvalidStatus    = 2
	ProcessedStatus  = 3
)

//easyjson:json
type Order struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Points    int64     `json:"points"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
