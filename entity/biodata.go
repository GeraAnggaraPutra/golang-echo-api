package entity

import "time"

type Biodata struct {
	ID        int
	NAME      string
	AGE       int
	ADDRESS   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
