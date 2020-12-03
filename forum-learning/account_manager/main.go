package main

import (
	"log"

	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/account_manager/utils"
)

func main() {
	server, err := utils.NewServer()

	if err != nil {
		log.Fatal(err)
	}

	server.Router.Run(":8080")
}

// func depedencyInjection(server *utils.Server) {
// 	dependecyinjection.Authentication(server)
// }
