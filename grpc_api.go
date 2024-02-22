package main

import (
	"context"
	"math/rand"
	"net"

	"github.com/alijabbar034/go-microservice/proto"
	"google.golang.org/grpc"
)

type GpcPriceFetcher struct {
	svc PriceFetcher
	proto.UnimplementedPriceFetcherServer
}

func makeGrpcServer(sv PriceFetcher, listenAdd string) error {

	grpcServer := NewGpcPriceFetcherServer(sv)
	ln, err := net.Listen("tcp", listenAdd)
	if err != nil {
		return err
	}
	opt := []grpc.ServerOption{}
	server := grpc.NewServer(opt...)
	proto.RegisterPriceFetcherServer(server, grpcServer)
	return server.Serve(ln)

}

func NewGpcPriceFetcherServer(svc PriceFetcher) *GpcPriceFetcher {
	return &GpcPriceFetcher{
		svc: svc,
	}
}

func (s *GpcPriceFetcher) FetchPrice(ctx context.Context, req *proto.PriceRequest) (*proto.PriceResponse, error) {

	reqId := rand.Intn(100000)
	ctx = context.WithValue(ctx, "reqID", reqId)
	price, err := s.svc.FetchPrice(ctx, req.GetTicker())
	if err != nil {
		return nil, err
	}

	reqP := &proto.PriceResponse{
		Price:  float32(price),
		Ticker: req.GetTicker(),
	}

	return reqP, nil

}
