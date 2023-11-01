package main

import (
	"RpcProvider/handler"
	service "RpcProvider/kitex_gen/service/provider"
	"log"
)

func main() {
	svr := service.NewServer(new(handler.ProviderImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
