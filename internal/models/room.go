package models

type Room struct {
	ID string `json:"id" building:"required"`
}
