package entity

import "time"

type Article struct {
	Id        int64     `json:"id"`
	AccountId int64     `json:"id_account"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
