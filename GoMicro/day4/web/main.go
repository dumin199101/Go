package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/web"
	"gomic-etcd/web/handler"
	"net/http"
)

func main() {

	// create new web service
	service := web.NewService(
		web.Name("go.micro.web.web"),
		web.Version("latest"),
		web.Address(":7799"),
		web.Registry(etcd.NewRegistry(registry.Addrs("127.0.0.1:2379"))),
	)

	// initialise service
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	// register html handler
	service.Handle("/", http.FileServer(http.Dir("/GoCode/src/Go/GoMicro/day4/web/html")))

	// register call handler
	service.HandleFunc("/web/call", handler.WebCall)

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
