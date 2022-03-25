package main

import (
	"fmt"
	"github.com/hyperf/roc"
	"github.com/hyperf/roc/exception"
	"github.com/hyperf/roc/formatter"
	"github.com/hyperf/roc/router"
	"github.com/hyperf/roc/server"
	"github.com/limingxinleo/roc-skeleton/action/roc_version"
)

func SetUpRouters() *router.SimpleRouter {
	r := router.NewSimpleRouter()
	r.Add("/r_o_c_version/getVersion", &roc_version.GetVersion{})
	r.Add("/r_o_c_version/hash", &roc_version.Hash{})
	return r
}

func main() {
	r := SetUpRouters()

	handler := server.NewTcpServerHandler(func(route *formatter.JsonRPCRoute, packet *roc.Packet, server *server.TcpServer) (any, exception.ExceptionInterface) {
		fmt.Println("Receive path: " + route.Path)

		action, ok := r.Routes[route.Path]
		if !ok {
			return nil, &exception.Exception{Code: exception.NOT_FOUND, Message: "The route is not defined."}
		}

		return action.Handle(packet, server.Serializer)
	})

	serv := server.NewTcpServer("0.0.0.0:9501", handler)

	serv.Start()
}
