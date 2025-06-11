package withdraw

import "time"

//easyjson:json
type WithDraw struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Order     string    `json:"order"`
	Sum       float64   `json:"sum"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
