package responses

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

func SuccessResponse(msg string) Response {
	return Response{
		StatusCode: 0,
		StatusMsg:  msg,
	}
}

func ErrorResponse(err error) Response {
	return Response{
		StatusCode: 1,
		StatusMsg:  err.Error(),
	}
}
