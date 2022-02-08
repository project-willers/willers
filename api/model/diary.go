package model

import "time"

type Diary struct {
	Name      string    `json:"name"`
	Time      time.Time `json:"time"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
