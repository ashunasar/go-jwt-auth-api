package models

type SignUpBody struct {
	Name     string `json:"name" validate:"required,min=3,max=15"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=15"`
}
