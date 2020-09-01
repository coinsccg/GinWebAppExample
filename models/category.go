package models

type Category struct {
	CategoryID   int32  `json:"id" db:"category_id"`
	CategoryName string `json:"name" db:"category_name"`
}
