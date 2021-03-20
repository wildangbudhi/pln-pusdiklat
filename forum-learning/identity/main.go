package main

import (
	"log"
	"os"

	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/depedencyinjection"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/utils"
)

func main() {

	var server *utils.Server
	var err error
	var filePointer *os.File

	server, err = utils.NewServer()

	if err != nil {
		log.Fatalf("Error Initiate Server: %v", err)
	}

	depedencyInjection(server)
	setLogToFile("./log/system.log", filePointer)
	defer filePointer.Close()

	server.Router.Run(":8080")

}

func depedencyInjection(server *utils.Server) {
	depedencyinjection.AuthenticationDI(server)
	depedencyinjection.AccountManagerDI(server)
}

func setLogToFile(filePath string, filePointer *os.File) {

	_, err := os.Stat(filePath)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(filePath)

		if err != nil {
			log.Fatalf("Error Creating Log File: %v", err)
		}

		defer file.Close()
	}

	filePointer, err = os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalf("Error Opening Log File: %v", err)
	}

	log.SetOutput(filePointer)

	filePointer, err = os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalf("Error Opening Log File: %v", err)
	}

	log.SetOutput(filePointer)
}
