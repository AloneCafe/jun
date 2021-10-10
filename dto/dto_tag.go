package dto

type Tag struct {
	ID   int64   `json:"t_id" db:"t_id"`
	Name *string `json:"t_name" db:"t_name"`
}
