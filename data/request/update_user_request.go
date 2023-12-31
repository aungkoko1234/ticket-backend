package request

type UpdateUserRequest struct {
	Id   string    `validate:"required"`
	Name string `validate:"required,min=1,max=200" json:"name"`
	Email string `validate:"required,min=1,max=200" json:"email"`
}