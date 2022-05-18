package main

import (
	"douyin-12306/cmd/user/dal"
	"douyin-12306/config"
	"douyin-12306/kitex_gen/userKitex/userservice"
	"douyin-12306/logger"
	"douyin-12306/pkg/middleware"
	"douyin-12306/pkg/tracer"
	"net"

	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
)

func Init() {
	tracer.InitJaeger(config.C.Services.User.Name)
	dal.Init()
}

func main() {
	logger.L = logger.NewZapLogger(config.C.Log.Out, config.C.Log.Level)
	defer logger.L.Sync()

	r, err := etcd.NewEtcdRegistry([]string{config.C.Etcd.Address})
	if err != nil {
		logger.L.Panic(err)
	}

	addr, err := net.ResolveTCPAddr("tcp", config.C.Services.User.Addr)
	if err != nil {
		logger.L.Panic(err)
	}
	Init()
	svr := userservice.NewServer(
		new(UserServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: config.C.Services.User.Name,
		}),
		server.WithMiddleware(middleware.CommonMiddleware),
		server.WithMiddleware(middleware.ServerMiddleware),
		server.WithServiceAddr(addr),
		server.WithLimit(&limit.Option{
			MaxConnections: config.C.Services.User.Server.MaxConnections,
			MaxQPS:         config.C.Services.User.Server.MaxQPS,
		}),
		server.WithMuxTransport(),
		server.WithSuite(trace.NewDefaultServerSuite()),
		server.WithRegistry(r),
	)
	err = svr.Run()
	if err != nil {
		logger.L.Fatal(err)
	}
}
