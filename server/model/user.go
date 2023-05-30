package model

type User struct {
	Model
	Name   string `gorm:"varchat(32);" json:"name"`
	Age    uint8  `json:"age"`
	Gender uint8  `json:"gender"`
}
