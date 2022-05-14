package responses

type UserRegisterResponse struct {
	Response
	LoginInfo
}

type UserLoginResponse struct {
	Response
	LoginInfo
}

type UserInfoResponse struct {
	Response
	User User `json:"user"`
}
