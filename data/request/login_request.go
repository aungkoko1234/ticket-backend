package request

type LoginRequest struct {
	Email string `validate:"required,min=1,max=200" json:"email"`
	Password string `validate:"required,min=1,max=200" json:"password"`
}