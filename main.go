package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const charset = "abcdefghijklmnopqrstuvwxyz0123456789"

var (
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func randomString(length int) string {
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[r.Intn(len(charset))]
	}
	return string(result)
}

func Connect() {
	for {
		opts := mqtt.NewClientOptions()
		opts.AddBroker("wss://o.popdeng.click:443/mqtt")
		opts.SetHTTPHeaders(http.Header{
			"origin":     {"https://popdeng.click"},
			"user-agent": {"Mozilla/5.0"},
		})
		opts.SetClientID(fmt.Sprintf("popdeng_%s", randomString(8)))
		opts.SetCleanSession(true)
		opts.SetDefaultPublishHandler(func(client mqtt.Client, msg mqtt.Message) {
			fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
		})
		opts.OnConnect = func(client mqtt.Client) {
			fmt.Println("Connected")
		}
		opts.OnConnectionLost = func(client mqtt.Client, err error) {
			fmt.Printf("Connect lost: %v", err)
		}
		client := mqtt.NewClient(opts)
		token := client.Connect()
		token.Wait()
		if token.Error() != nil {
			fmt.Println("Connect Fail", token.Error())
			continue
		}
		//client.Subscribe("popdeng/leaderboard", 1, nil)
		select{}
		for {
			//token = client.Publish("popdeng/leaderboard", 0, false, `[{"country_code":"WW","clicks":0,"cps":0},{"country_code":"KP","clicks":6969696969,"cps":6969696969},{"country_code":"IS","clicks":6969696969,"cps":6969696969},{"country_code":"WW","clicks":0,"cps":0},{"country_code":"HK","clicks":69,"cps":1},{"country_code":"AT","clicks":69,"cps":2},{"country_code":"CN","clicks":69,"cps":3},{"country_code":"KW","clicks":69,"cps":4},{"country_code":"ZZ","clicks":0,"cps":0},{"country_code":"BE","clicks":69,"cps":5},{"country_code":"YE","clicks":69,"cps":6},{"country_code":"ZZ","clicks":0,"cps":0},{"country_code":"LR","clicks":69,"cps":7},{"country_code":"AQ","clicks":69,"cps":8},{"country_code":"RO","clicks":69,"cps":9},{"country_code":"IS","clicks":69,"cps":10},{"country_code":"NO","clicks":69,"cps":11},{"country_code":"AO","clicks":69,"cps":12},{"country_code":"ZZ","clicks":0,"cps":0},{"country_code":"TH","clicks":1000,"cps":1}]`)
			token = client.Publish("popdeng/clicks", 0, false, "RHKLKWVMTGVHG8CCC!.8h4n") // iceland
			token.Wait()
			if token.Error() != nil {
				fmt.Println("send click error", token.Error())
				break
			}
			time.Sleep(time.Millisecond)
		}
		client.Disconnect(0)
	}
}

func main() {
	// adjust threads here for better cps / connection
	for i := 0; i < 50; i++ {
		go Connect()
	}
	select {}
}