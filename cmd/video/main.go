package main

import (
	"douyin-12306/config"
	videoKitex "douyin-12306/kitex_gen/videoKitex/videoservice"
	"douyin-12306/logger"
	"douyin-12306/pkg/middleware"
	"net"

	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	logger.L = logger.NewZapLogger(config.C.Log.Out, config.C.Log.Level)
	defer logger.L.Sync()

	r, err := etcd.NewEtcdRegistry([]string{config.C.Etcd.Address})
	if err != nil {
		logger.L.Panic(err)
	}

	addr, err := net.ResolveTCPAddr("tcp", config.C.Services.Video.Addr)
	if err != nil {
		logger.L.Panic(err)
	}

	svr := videoKitex.NewServer(
		new(VideoServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: config.C.Services.Video.Name,
		}),
		server.WithMiddleware(middleware.CommonMiddleware),
		server.WithMiddleware(middleware.ServerMiddleware),
		server.WithServiceAddr(addr),
		server.WithLimit(&limit.Option{
			MaxConnections: config.C.Services.Video.Server.MaxConnections,
			MaxQPS:         config.C.Services.Video.Server.MaxQPS,
		}),
		server.WithMuxTransport(),
		server.WithRegistry(r),
	)
	logger.L.Infow("Init [service:user] serever", map[string]interface{}{
		"config": config.C.Services.Video.Server,
	})

	err = svr.Run()
	if err != nil {
		logger.L.Fatal(err)
	}
}
