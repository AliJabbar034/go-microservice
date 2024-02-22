package main

import (
	"context"
	"fmt"
	"log"

	"github.com/alijabbar034/go-microservice/client"
)

func main() {
	go func() {
		client := client.NewClient("http://localhost:8080")
		pric, err := client.FetchPrice(context.Background(), "AAPL")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(pric)
		return
	}()
	svc := NewLoginService(NewMetricsService(&priceFetcher{}))
	server := NewJsonApiServer(":8080", svc)
	server.Run()
	price, err := svc.FetchPrice(context.Background(), "MSFT")
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
		return
	}
	fmt.Println(price)
}
