package request

type RegisterRequest struct {
	Name string `validate:"required,min=1,max=200" json:"name"`
	Email string `validate:"required,min=1,max=200" json:"email"`
	Password string `validate:"required,min=1,max=200" json:"password"`
}