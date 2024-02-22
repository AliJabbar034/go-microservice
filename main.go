package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/alijabbar034/go-microservice/client"
	"github.com/alijabbar034/go-microservice/proto"
)

func main() {
	const (
		jsonAdd string = ":8080"
		grpcAdd string = ":9000"
	)
	ctx := context.Background()
	go func() {
		client := client.NewClient("http://localhost:8080")
		pric, err := client.FetchPrice(context.Background(), "AAPL")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(pric)
		return
	}()
	grpcClient, err := client.NewGrpcClient(grpcAdd)
	if err != nil {
		log.Fatal(err)
	}
	go func() {

		for {
			time.Sleep(4 * time.Second)
			res, err := grpcClient.FetchPrice(ctx, &proto.PriceRequest{
				Ticker: "MSFT",
			})
			if err != nil {
				log.Fatal(err)

			}
			fmt.Println("Fetching price grpc", res)
		}
	}()
	svc := NewLoginService(NewMetricsService(&priceFetcher{}))
	go makeGrpcServer(svc, grpcAdd)
	server := NewJsonApiServer(jsonAdd, svc)
	server.Run()

	price, err := svc.FetchPrice(context.Background(), "MSFT")
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
		return
	}
	fmt.Println(price)
}
