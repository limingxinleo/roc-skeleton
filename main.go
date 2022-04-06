package main

import (
	"github.com/hyperf/roc"
	"github.com/hyperf/roc/exception"
	"github.com/hyperf/roc/formatter"
	"github.com/hyperf/roc/log"
	"github.com/hyperf/roc/router"
	"github.com/hyperf/roc/server"
	"github.com/joho/godotenv"
	"github.com/limingxinleo/roc-skeleton/action/roc_version"
	"go.uber.org/zap"
	"time"
)

func SetUpRouters() *router.SimpleRouter {
	r := router.NewSimpleRouter()
	r.Add("/r_o_c_version/getVersion", &roc_version.GetVersion{})
	r.Add("/r_o_c_version/hash", &roc_version.Hash{})
	return r
}

func main() {
	_ = godotenv.Load()

	r := SetUpRouters()

	handler := server.NewTcpServerHandler(func(route *formatter.JsonRPCRoute, packet *roc.Packet, server *server.TcpServer) (any, exception.ExceptionInterface) {
		name := zap.String("name", "RPC")
		log.Logger().Info("RPC_RECEIVED", name, zap.String("path", route.Path))
		now := time.Now()

		action, ok := r.Routes[route.Path]
		if !ok {
			log.Logger().Warn("The route is not defined.", name, zap.String("path", route.Path))
			return nil, &exception.Exception{Code: exception.NOT_FOUND, Message: "The route is not defined."}
		}

		ret, e := action.Handle(packet, server.Serializer)

		// 记录接口调用时间
		t := time.Now().UnixMilli() - now.UnixMilli()
		if t > 200 {
			log.Logger().Error("RPC_TIME", name, zap.String("path", route.Path), zap.Int64("time", t))
		} else {
			log.Logger().Info("RPC_TIME", name, zap.String("path", route.Path), zap.Int64("time", t))
		}

		return ret, e
	})

	serv := server.NewTcpServer("0.0.0.0:9501", handler)

	serv.Start()
}
