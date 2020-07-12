package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/util/log"
	_ "github.com/micro/go-plugins/transport/nats/v2"
	masterpoto "tionyxtrack/masterservice/proto"
	api "tionyxtrack/masterservice/protoapi"
)

func main() {

	service := micro.NewService(
		micro.Name("go.micro.srv.master"),
	)
	service.Init()

	err := masterpoto.RegisterMasterServiceHandler(service.Server(), new(api.MasterService))
	if err != nil {
		log.Log(err)
	}
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
