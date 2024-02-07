package domain

type Role struct {
	ID          int          `json:"id"`
	Name        string       `json:"name"`
	Permissions []Permission `json:"permissions"`
}