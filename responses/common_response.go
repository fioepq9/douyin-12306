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

type User struct {
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

type LoginInfo struct {
	Id    int64
	Token string
}
