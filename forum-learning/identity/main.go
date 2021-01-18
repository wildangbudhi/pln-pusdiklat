package main

import (
	"log"

	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/depedencyinjection"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/utils"
)

func main() {
	server, err := utils.NewServer()

	if err != nil {
		log.Fatal(err)
	}

	depedencyInjection(server)

	server.Router.Run(":8080")
}

func depedencyInjection(server *utils.Server) {
	depedencyinjection.AuthenticationDI(server)
}
