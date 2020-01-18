package main

import (
	"mytest/handler"
//	"mytest/subscriber"

	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-plugins/broker/rabbitmq"

	example "mytest/proto/example"
)

func main() {

	// Init broker
	cmd.Init()
	rabbitmq.DefaultExchange = "cq"
	if err := broker.Init(); err != nil {
		log.Fatalf("Broker Init error: %v", err)
	}
	if err := broker.Connect(); err != nil {
		log.Fatalf("Broker Connect error: %v", err)
	}
	
	// New Service
	service := micro.NewService(
		micro.Name("cq.srv.test.service"),
		micro.Version("latest"),
	)

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Register Struct as Subscriber
//	micro.RegisterSubscriber("topic.cq.srv.test.service", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
//	micro.RegisterSubscriber("topic.cq.srv.test.service", service.Server(), subscriber.Handler)

	// Initialise service
	service.Init()

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
