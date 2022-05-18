package rpc

import (
	"context"
	"douyin-12306/config"
	"douyin-12306/kitex_gen/userKitex"
	"douyin-12306/kitex_gen/userKitex/userservice"
	"douyin-12306/logger"
	"douyin-12306/pkg/errno"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
)

var userClient userservice.Client

func initUserRpc() {
	r, err := etcd.NewEtcdResolver([]string{
		config.C.Etcd.Address,
	})
	if err != nil {
		logger.L.Panic(err)
	}
	c, err := userservice.NewClient(
		config.C.Services.User.Name,
		client.WithMuxConnection(config.C.Services.User.Client.MuxConnection),
		client.WithRPCTimeout(config.C.Services.User.Client.RpcTimeout*time.Second),
		client.WithConnectTimeout(config.C.Services.User.Client.ConnTimeout*time.Millisecond),
		client.WithFailureRetry(retry.NewFailurePolicy()),
		client.WithSuite(trace.NewDefaultClientSuite()),
		client.WithResolver(r),
	)
	if err != nil {
		logger.L.Panic(err)
	}
	userClient = c
	logger.L.Infow("Init [service:user] client success", map[string]interface{}{
		"etcd addr": config.C.Etcd.Address,
		"config":    config.C.Services.User.Client,
	})
}

func RegisterUser(ctx context.Context, req *userKitex.UserRegisterRequest) (*userKitex.UserRegisterResponse, error) {
	resp, err := userClient.RegisterUser(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.Response.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.Response.StatusCode, resp.Response.StatusMsg)
	}
	return resp, nil
}

func LoginUser(ctx context.Context, req *userKitex.UserLoginRequest) (*userKitex.UserLoginResponse, error) {
	resp, err := userClient.LoginUser(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.Response.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.Response.StatusCode, resp.Response.StatusMsg)
	}
	return resp, nil
}

func QueryUserInfo(ctx context.Context, req *userKitex.UserInfoRequest) (*userKitex.UserInfoResponse, error) {
	resp, err := userClient.QueryUserInfo(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.Response.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.Response.StatusCode, resp.Response.StatusMsg)
	}
	return resp, nil
}
