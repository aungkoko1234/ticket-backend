package request

type LoginRequest struct {
	Email string `json:"email" binding:"required,email,min=1,max=200" `
	Password string `json:"password" binding:"required,min=8"`
}