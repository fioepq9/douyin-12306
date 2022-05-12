package requests

type UserRegisterRequest struct {
	Username string `form:"username"`
	Password string `form:"password"`
}
