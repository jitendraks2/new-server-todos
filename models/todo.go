package models

import "time"

type Todos struct {
	Id          string    `json:"id"`
	Todo        string    `json:"todo"`
	IsCompleted bool      `json:"completed"`
	CreatedAt   time.Time `json:"createdAt"`
	// CompletedAt time.Time `json:"deletedAt"`
}
