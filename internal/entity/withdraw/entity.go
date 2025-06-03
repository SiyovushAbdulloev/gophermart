package withdraw

import "time"

//easyjson:json
type WithDraw struct {
	Id        int       `json:"id"`
	UserId    int       `json:"user_id"`
	Order     int64     `json:"order"`
	Sum       int64     `json:"sum"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
