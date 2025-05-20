package models

type Book struct {
	ID     int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Title  string `json:"title"`
	Author string `json:"author"`
}
