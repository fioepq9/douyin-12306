package main

import (
	"context"
	"douyin-12306/kitex_gen/videoKitex"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// Feed implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Feed(ctx context.Context, req *videoKitex.FeedRequest) (resp *videoKitex.FeedResponse, err error) {
	// TODO: Your code here...
	return
}

// PublishAction implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishAction(ctx context.Context, req *videoKitex.PublishActionRequest) (resp *videoKitex.PublishActionResponse, err error) {
	// TODO: Your code here...
	return
}

// QueryPublishList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) QueryPublishList(ctx context.Context, req *videoKitex.PublishListRequest) (resp *videoKitex.PublishListResponse, err error) {
	// TODO: Your code here...
	return
}

// FavoriteAction implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) FavoriteAction(ctx context.Context, req *videoKitex.FavoriteActionRequest) (resp *videoKitex.FavoriteActionResponse, err error) {
	// TODO: Your code here...
	return
}

// QueryFavoriteList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) QueryFavoriteList(ctx context.Context, req *videoKitex.FavoriteListRequest) (resp *videoKitex.FavoriteListResponse, err error) {
	// TODO: Your code here...
	return
}

// CommentAction implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) CommentAction(ctx context.Context, req *videoKitex.CommentActionRequest) (resp *videoKitex.CommentActionResponse, err error) {
	// TODO: Your code here...
	return
}

// QueryCommentList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) QueryCommentList(ctx context.Context, req *videoKitex.CommentListRequset) (resp *videoKitex.CommentListResponse, err error) {
	// TODO: Your code here...
	return
}
