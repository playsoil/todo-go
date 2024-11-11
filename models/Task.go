package models

type Task struct {
	ID    uint   `gorm:"primarykey"`
	Title string `json:"title" gorm:"string; not null"`
}
