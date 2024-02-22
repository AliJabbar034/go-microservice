package main

import (
	"context"
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/alijabbar034/go-microservice/types"
)

type JsonApiServer struct {
	svc        PriceFetcher
	listenAddr string
}
type ApiFunc func(context.Context, http.ResponseWriter, *http.Request) error

func (s *JsonApiServer) Run() {
	http.HandleFunc("/", makeApiFun(s.handleFetchPrice))
	http.ListenAndServe(s.listenAddr, nil)
}

func NewJsonApiServer(listenAddr string, svc PriceFetcher) *JsonApiServer {

	return &JsonApiServer{
		svc:        svc,
		listenAddr: listenAddr,
	}
}

func makeApiFun(fn ApiFunc) http.HandlerFunc {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "reqID", rand.Intn(10000))
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(ctx, w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *JsonApiServer) handleFetchPrice(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	ticker := r.URL.Query().Get("ticker")
	price, err := s.svc.FetchPrice(ctx, ticker)
	if err != nil {
		return err
	}

	priceRes := types.PriceResponse{
		Price:  price,
		Ticker: ticker,
	}
	return writeJson(w, http.StatusOK, &priceRes)

}

func writeJson(w http.ResponseWriter, s int, v any) error {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(s)
	return json.NewEncoder(w).Encode(v)
}
