package main

import (
	"log"
	"sync"

	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/depedencyinjection"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/utils"
)

func main() {

	var wg sync.WaitGroup
	server, err := utils.NewServer()

	if err != nil {
		log.Fatal(err)
	}

	depedencyInjection(server)

	wg.Add(2)

	go func(server *utils.Server, wg *sync.WaitGroup) {
		defer wg.Done()
		server.Router.Run(":8080")
	}(server, &wg)

	go func(server *utils.Server, wg *sync.WaitGroup) {
		defer wg.Done()
		server.QueueServer.RunConsumerServer()
	}(server, &wg)

	wg.Wait()

}

func depedencyInjection(server *utils.Server) {
	depedencyinjection.EventConsumerDI(server)
}
