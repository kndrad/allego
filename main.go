package main

import (
	"allego/device"
	"allego/email"
	"allego/resource"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	clientID, secret := os.Getenv("CLIENT_ID"), os.Getenv("SECRET_ID")
	client := http.DefaultClient

	resp, err := device.Token(client, clientID, secret)
	if err != nil {
		log.Fatal(err)
	}

	cart, err := resource.Cart(client, resp.AccessToken)
	if err != nil {
		log.Fatal(err)
	}
	orders := cart.FindOrders("Pies Demonek wspiera WOŚP! - zdjęcia")

	details := make(chan *resource.Details, len(orders))
	var wg sync.WaitGroup
	for _, order := range orders {
		wg.Add(1)
		go func(id string) {
			defer wg.Done()
			fmt.Println(id)
			orderDetails, err := resource.OrderDetails(id, client, resp.AccessToken)
			if err != nil {
				log.Fatal(err)
			}
			details <- orderDetails
		}(order.ID)
	}
	wg.Wait()
	close(details)

	for d := range details {
		fmt.Println(email.Extract(d.Message))
	}

}
