package main

import (
	"encoding/json"
	"er-rabbit-consumer/database"
	"er-rabbit-consumer/model"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

func init() {
	godotenv.Load()
}
func main() {
	fmt.Println("Consumer Application")
	host := os.Getenv("MQHOST")
	mqport := os.Getenv("MQPORT")
	user := os.Getenv("MQUSER")
	password := os.Getenv("MQPASSWORD")
	mqURL := fmt.Sprintf("amqp://%s:%s@%s:%s/", user, password, host, mqport)
	conn, err := amqp.Dial(mqURL)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer ch.Close()
	msgs, _ := ch.Consume(
		"ExchangeRate",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	var exchangeRates model.ExchangeRate
	forever := make(chan bool)
	go func() {
		for message := range msgs {
			bodyBytes := (message.Body)
			if err != nil {
				panic(err.Error())
			}
			json.Unmarshal(bodyBytes, &exchangeRates)
			if err != nil {
				panic(err.Error())
			}
			fmt.Println(exchangeRates)
			database.Connect(exchangeRates)
		}
	}()
	fmt.Println("Succesfully connected to our RabbitMQ Instance")
	<-forever
}
