package request

type RegisterRequest struct {
	Name string `json:"name" binding:"required,min=1,max=60"`
	Email string `json:"email" binding:"required,email,min=1,max=60"`
	Password string `json:"password" binding:"required,min=8,max=20"`
}