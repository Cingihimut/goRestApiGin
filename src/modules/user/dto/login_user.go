package user

type LoginInput struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"required,min=8"`
}
