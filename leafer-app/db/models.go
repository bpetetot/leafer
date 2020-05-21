package db

import "time"

// Library model
type Library struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	Path      string    `json:"path"`
	Medias    *[]Media  `json:"medias,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
