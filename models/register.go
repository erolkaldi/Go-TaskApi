package models

type RegisterDto struct {
	Username string `json:"user_name" validate:"required,min=2,max=255"`
	Email    string `json:"email"`
	Password string `json:"password" validate:"required"`
	UserType string `json:"user_type" validate:"required,eq=ADMIN|eq=GUEST|eq=USER"`
}
