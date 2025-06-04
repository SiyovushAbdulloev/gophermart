package order

import "time"

var (
	NewStatus        = "NEW"
	ProcessingStatus = "PROCESSING"
	InvalidStatus    = "INVALID"
	ProcessedStatus  = "PROCESSED"
)

//easyjson:json
type Order struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Points    float64   `json:"points"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (o *Order) StatusString() string {
	if o.Status == NewStatus {
		return "NEW"
	}
	if o.Status == ProcessingStatus {
		return "PROCESSING"
	}
	if o.Status == InvalidStatus {
		return "INVALID"
	}
	if o.Status == ProcessedStatus {
		return "PROCESSED"
	}

	return ""
}
