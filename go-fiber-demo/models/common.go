package models

import "time"

type TimeStamps struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TimeStampsH struct {
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
