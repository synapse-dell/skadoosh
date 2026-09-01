package models

type User struct {
	ID   int    `json:"id" gorm:"primaryKey"` //this tell gorm that ID is the primary key
	Name string `json:"name"`
	Age  int    `json:"age"`
}
