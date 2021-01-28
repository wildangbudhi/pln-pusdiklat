package main

import (
	"log"
	"os"
	"sync"

	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/depedencyinjection"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/utils"
)

func main() {

	var wg sync.WaitGroup
	var server *utils.Server
	var err error
	var filePointer *os.File

	server, err = utils.NewServer()

	if err != nil {
		log.Fatal(err)
	}

	depedencyInjection(server)
	setLogToFile("./log/system.log", filePointer)
	defer filePointer.Close()

	wg.Add(2)

	// Run HTTP Server
	go func(server *utils.Server, wg *sync.WaitGroup) {
		defer wg.Done()
		server.Router.Run(":8080")
	}(server, &wg)

	// Run RabbitMQ Consumer
	go func(server *utils.Server, wg *sync.WaitGroup) {
		defer wg.Done()
		server.QueueServer.RunConsumerServer()
	}(server, &wg)

	wg.Wait()

}

func depedencyInjection(server *utils.Server) {
	depedencyinjection.ForumDI(server)
}

func setLogToFile(filePath string, filePointer *os.File) {
	filePointer, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalf("Error Opening Log File: %v", err)
	}

	log.SetOutput(filePointer)
}
