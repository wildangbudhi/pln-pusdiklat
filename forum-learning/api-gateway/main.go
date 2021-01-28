package main

import (
	"log"
	"sync"

	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/api-gateway/utils"
)

func main() {
	server, err := utils.NewServer()

	if err != nil {
		log.Fatal(err)
	}

	defer server.DB.Disconnect(*server.DBContext)

	serverAdmin, err := utils.NewServer()

	if err != nil {
		log.Fatal(err)
	}

	defer serverAdmin.DB.Disconnect(*serverAdmin.DBContext)

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		server.Router.Run(":8080")
	}()

	go func() {
		defer wg.Done()
		serverAdmin.Router.Run(":8081")
	}()

	wg.Wait()

}

// func depedencyInjection(server *utils.Server) {
// 	dependecyinjection.Authentication(server)
// }
