package requests

type FeedRequest struct {
	LatestTime int64  `form:"latest_time" binding:"min=1,max=64"`
	Token      string `form:"token" binding:"min=1,max=64"`
}
