package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro/go-plugins/transport/nats/v2"
	nat "github.com/nats-io/nats.go"
	masterpoto "tionyxtrack/masterservice/proto"
	api "tionyxtrack/masterservice/protoapi"
)

func main() {
	transport := nats.NewTransport(nats.Options(nat.GetDefaultOptions()))
	service := micro.NewService(
		micro.Name("go.micro.srv.master"),
		micro.Transport(transport),
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
