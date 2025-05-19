package withdraw

import "time"

//easyjson:json
type WithDraw struct {
	Id        int       `json:"id"`
	UserId    int       `json:"user_id"`
	OrderId   int       `json:"order_id"`
	Sum       int64     `json:"sum"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
