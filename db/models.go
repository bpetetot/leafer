package db

// Library model
type Library struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
	Path string `json:"path"`
}
