package utils

import (
	"fmt"
	"log"
	"sync"

	"github.com/streadway/amqp"
)

type consumerFunction struct {
	queueName    string
	routingKey   string
	exchangeName string
	exchangeType string
	function     func(message amqp.Delivery)
}

// MessageQueueServer is Object for Message Queue Service
type MessageQueueServer struct {
	queueConnection      *amqp.Connection
	consumerFunctionList []consumerFunction
	waitGroupGoRoutine   sync.WaitGroup
}

// NewMessageQueueServer is a constructor for MessageQueueServer
func NewMessageQueueServer(host, port, user, password string) (*MessageQueueServer, error) {

	var err error
	server := new(MessageQueueServer)

	connectionString := fmt.Sprintf("amqp://%s:%s@%s:%s/", user, password, host, port)
	server.queueConnection, err = amqp.Dial(connectionString)

	if err != nil {
		return nil, err
	}

	server.consumerFunctionList = make([]consumerFunction, 0)

	return server, nil
}

// RegisterConsumerController is a function handler for registering function for consumer
func (server *MessageQueueServer) RegisterConsumerController(queueName string, routingKey string, exchangeName string, exchangeType string, controllerFunc func(message amqp.Delivery)) {
	function := consumerFunction{queueName: queueName, routingKey: routingKey, exchangeName: exchangeName, exchangeType: exchangeType, function: controllerFunc}
	server.consumerFunctionList = append(server.consumerFunctionList, function)
	server.waitGroupGoRoutine.Add(1)
}

// RunConsumerServer to start Server consuming data from Queue using function that has been registered
func (server *MessageQueueServer) RunConsumerServer() {

	for i := 0; i < len(server.consumerFunctionList); i++ {
		log.Println("Run Queue Consumer : " + server.consumerFunctionList[i].routingKey)
		go server.consumeQueue(server.consumerFunctionList[i])
	}

	server.waitGroupGoRoutine.Wait()
}

func (server *MessageQueueServer) consumeQueue(consumerFunc consumerFunction) {
	defer server.waitGroupGoRoutine.Done()

	ch, err := server.queueConnection.Channel()

	if err != nil {
		log.Fatal("Failed to Create Channel : " + err.Error())
	}

	err = ch.ExchangeDeclare("event", "topic", true, false, false, false, nil)

	if err != nil {
		log.Fatal("Failed to Declare Exchange : " + err.Error())
	}

	defer ch.Close()

	queue, err := ch.QueueDeclare(consumerFunc.queueName, false, false, false, false, nil)

	err = ch.QueueBind(queue.Name, consumerFunc.routingKey, "event", false, nil)

	if err != nil {
		log.Fatal("Failed to Bind Queue : " + err.Error())
	}

	consumer, err := ch.Consume(queue.Name, "", false, false, false, false, nil)

	if err != nil {
		log.Fatal("Failed to Consume Message from Queue : " + err.Error())
	}

	for msg := range consumer {
		consumerFunc.function(msg)
		msg.Ack(false)
	}
}

// Publish is Function to Publish new data to queue
func (server *MessageQueueServer) Publish(routingKey string, exchangeName string, exchangeType string, msg []byte, contentType string) error {

	ch, err := server.queueConnection.Channel()

	if err != nil {
		return err
	}

	defer ch.Close()

	err = ch.ExchangeDeclare(exchangeName, exchangeType, true, false, false, false, nil)

	if err != nil {
		return err
	}

	if err != nil {
		return err
	}

	payload := amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  contentType,
		Body:         msg,
	}

	err = ch.Publish(exchangeName, routingKey, false, false, payload)

	if err != nil {
		return err
	}

	return nil
}
