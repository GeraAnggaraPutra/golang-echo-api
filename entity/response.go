package entity

import "time"

type Response struct {
	ID        int       `json:"id"`
	NAME      string    `json:"name"`
	AGE       int       `json:"age"`
	ADDRESS   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
