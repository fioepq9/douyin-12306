package requests

type UserRegisterRequest struct {
	Username string `form:"username" binding:"required,min=6,max=32"`
	Password string `form:"password" binding:"required,min=6,max=32"`
}