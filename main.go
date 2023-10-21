package main

import (
	service "RpcProvider/kitex_gen/service/provider"
	"log"
)

func main() {
	svr := service.NewServer(new(ProviderImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
