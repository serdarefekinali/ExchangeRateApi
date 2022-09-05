package main

import (
	"encoding/json"
	"er-api-consumer/model"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

var conn *amqp.Connection
var ch *amqp.Channel
var err error

func init() {
	godotenv.Load()
}
func MySecretKey() (string, string) {
	key := os.Getenv("APIKEY")
	link := os.Getenv("APILINK")
	return key, link
}
func ApiUrl() (string, string, string, string, string, string, string, string, string) {
	mykey, apilink := MySecretKey()
	eurtry := fmt.Sprintf("%s/%s/pair/EUR/TRY", apilink, mykey)
	tryeur := fmt.Sprintf("%s/%s/pair/TRY/EUR", apilink, mykey)
	eureur := fmt.Sprintf("%s/%s/pair/EUR/EUR", apilink, mykey)
	usdtry := fmt.Sprintf("%s/%s/pair/USD/TRY", apilink, mykey)
	tryusd := fmt.Sprintf("%s/%s/pair/TRY/USD", apilink, mykey)
	usdusd := fmt.Sprintf("%s/%s/pair/USD/USD", apilink, mykey)
	eurusd := fmt.Sprintf("%s/%s/pair/EUR/USD", apilink, mykey)
	usdeur := fmt.Sprintf("%s/%s/pair/USD/EUR", apilink, mykey)
	trytry := fmt.Sprintf("%s/%s/pair/TRY/TRY", apilink, mykey)
	return eurtry, tryeur, eureur, usdtry, tryusd, usdusd, eurusd, usdeur, trytry
}

// et, te, ee, ut, tu, uu, eu, ue, tt := ApiUrl()
func EurTry() (model.ExchangeRate, error) {
	et, _, _, _, _, _, _, _, _ := ApiUrl()
	response, err := http.Get(et)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer response.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(response.Body)
	var exchangerates model.ExchangeRate
	err = json.Unmarshal(bodyBytes, &exchangerates)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(exchangerates)
	return exchangerates, nil
}

func TryEur() (model.ExchangeRate, error) {
	_, te, _, _, _, _, _, _, _ := ApiUrl()
	response, err := http.Get(te)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer response.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(response.Body)
	var exchangerates model.ExchangeRate
	err = json.Unmarshal(bodyBytes, &exchangerates)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(exchangerates)
	return exchangerates, nil
}

func EurEur() (model.ExchangeRate, error) {
	_, _, ee, _, _, _, _, _, _ := ApiUrl()
	response, err := http.Get(ee)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer response.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(response.Body)
	var exchangerates model.ExchangeRate
	err = json.Unmarshal(bodyBytes, &exchangerates)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(exchangerates)
	return exchangerates, nil
}

func UsdTry() (model.ExchangeRate, error) {
	_, _, _, ut, _, _, _, _, _ := ApiUrl()
	response, err := http.Get(ut)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer response.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(response.Body)
	var exchangerates model.ExchangeRate
	err = json.Unmarshal(bodyBytes, &exchangerates)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(exchangerates)
	return exchangerates, nil
}

func TryUsd() (model.ExchangeRate, error) {
	_, _, _, _, tu, _, _, _, _ := ApiUrl()
	response, err := http.Get(tu)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer response.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(response.Body)
	var exchangerates model.ExchangeRate
	err = json.Unmarshal(bodyBytes, &exchangerates)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(exchangerates)
	return exchangerates, nil
}

func UsdUsd() (model.ExchangeRate, error) {
	_, _, _, _, _, uu, _, _, _ := ApiUrl()
	response, err := http.Get(uu)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer response.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(response.Body)
	var exchangerates model.ExchangeRate
	err = json.Unmarshal(bodyBytes, &exchangerates)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(exchangerates)
	return exchangerates, nil
}

func EurUsd() (model.ExchangeRate, error) {
	_, _, _, _, _, _, eu, _, _ := ApiUrl()
	response, err := http.Get(eu)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer response.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(response.Body)
	var exchangerates model.ExchangeRate
	err = json.Unmarshal(bodyBytes, &exchangerates)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(exchangerates)
	return exchangerates, nil
}

func UsdEur() (model.ExchangeRate, error) {
	_, _, _, _, _, _, _, eu, _ := ApiUrl()
	response, err := http.Get(eu)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer response.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(response.Body)
	var exchangerates model.ExchangeRate
	err = json.Unmarshal(bodyBytes, &exchangerates)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(exchangerates)
	return exchangerates, nil
}

func TryTry() (model.ExchangeRate, error) {
	_, _, _, _, _, _, _, _, tt := ApiUrl()
	response, err := http.Get(tt)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer response.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(response.Body)
	var exchangerates model.ExchangeRate
	err = json.Unmarshal(bodyBytes, &exchangerates)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(exchangerates)
	return exchangerates, nil
}

func init() {
	godotenv.Load()
}
func main() {
	host := os.Getenv("MQHOST")
	mqport := os.Getenv("MQPORT")
	user := os.Getenv("MQUSER")
	password := os.Getenv("MQPASSWORD")
	mqURL := fmt.Sprintf("amqp://%s:%s@%s:%s/", user, password, host, mqport)
	conn, err = amqp.Dial(mqURL)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Succesfully Connected To our RabbitMQ Instance")
	}
	defer conn.Close()
	ch, err = conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer ch.Close()
	i := 1
	for {
		now := time.Now()
		fmt.Printf("%s %d %s %d %d:%d:%d.%d\n", now.Weekday(), now.Day(), now.Month(), now.Year(), now.Hour(), now.Minute(), now.Second(), now.Nanosecond())
		q, err := ch.QueueDeclare(
			"ExchangeRate",
			false,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		fmt.Println(q)
		for j := 0; j < 9; j++ {
			if j == 0 {
				calledParity, err := EurTry()
				if err != nil {
					fmt.Println(err)
					panic(err)
				}
				convertedByte, err := json.Marshal(calledParity)
				if err != nil {
					fmt.Println(err)
					panic(err)
				}
				PublishFunc(convertedByte)
				if err != nil {
					fmt.Println(err)
					panic(err)
				}
			}
			if j == 1 {
				calledParity, err := TryEur()
				if err != nil {
					fmt.Println(err)
					panic(err)
				}
				convertedByte, err := json.Marshal(calledParity)
				if err != nil {
					fmt.Println(err)
					panic(err)
				}
				PublishFunc(convertedByte)
				if err != nil {
					fmt.Println(err)
					panic(err)
				}
			}
			if j == 2 {
				calledParity, err := EurEur()
				if err != nil {
					fmt.Println(err)
					panic(err)
				}
				convertedByte, err := json.Marshal(calledParity)
				if err != nil {
					fmt.Println(err)
					panic(err)
				}
				PublishFunc(convertedByte)
				if err != nil {
					fmt.Println(err)
					panic(err)
				}
			}
			if j == 3 {
				calledParity, err := UsdTry()

				if err != nil {
					fmt.Println(err)
					panic(err)
				}
				convertedByte, err := json.Marshal(calledParity)
				if err != nil {
					fmt.Println(err)
					panic(err)
				}
				PublishFunc(convertedByte)
				if err != nil {
					fmt.Println(err)
					panic(err)
				}
			}
			if j == 4 {
				calledParity, err := TryUsd()
				if err != nil {
					fmt.Println(err)
					panic(err)
				}
				convertedByte, err := json.Marshal(calledParity)
				if err != nil {
					fmt.Println(err)
					panic(err)
				}
				PublishFunc(convertedByte)
				if err != nil {
					fmt.Println(err)
					panic(err)
				}
			}
			if j == 5 {
				calledParity, err := UsdUsd()
				if err != nil {
					fmt.Println(err)
					panic(err)
				}
				convertedByte, err := json.Marshal(calledParity)
				if err != nil {
					fmt.Println(err)
					panic(err)
				}
				PublishFunc(convertedByte)
				if err != nil {
					fmt.Println(err)
					panic(err)
				}
			}
			if j == 6 {
				calledParity, err := EurUsd()
				if err != nil {
					fmt.Println(err)
					panic(err)
				}
				convertedByte, err := json.Marshal(calledParity)
				if err != nil {
					fmt.Println(err)
					panic(err)
				}
				PublishFunc(convertedByte)
				if err != nil {
					fmt.Println(err)
					panic(err)
				}
			}
			if j == 7 {
				calledParity, err := UsdEur()
				if err != nil {
					fmt.Println(err)
					panic(err)
				}
				convertedByte, err := json.Marshal(calledParity)
				if err != nil {
					fmt.Println(err)
					panic(err)
				}
				PublishFunc(convertedByte)
				if err != nil {
					fmt.Println(err)
					panic(err)
				}
			}
			if j == 8 {
				calledParity, err := TryTry()
				if err != nil {
					fmt.Println(err)
					panic(err)
				}
				convertedByte, err := json.Marshal(calledParity)
				if err != nil {
					fmt.Println(err)
					panic(err)
				}
				PublishFunc(convertedByte)
				if err != nil {
					fmt.Println(err)
					panic(err)
				}
			}
		}
		time.Sleep(59 * time.Second)
		if i == 10 {
			break
		}
	}
}

func PublishFunc(convertedByte []byte) error {
	err := ch.Publish(
		"",
		"ExchangeRate",
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         convertedByte},
	)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return nil
}
