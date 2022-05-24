package requests

type FeedRequest struct {
	LatestTime int64  `form:"latest_time" binding:"min=0,max=64"`
	Token      string `form:"token" binding:"min=0,max=64"`
}
