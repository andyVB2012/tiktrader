package main

import (
	"context"
	"encoding/json"
	"main/mytypes"
	"math/rand"
	"net/http"
)

type JSONAPIServer struct {
	listenAddr string
	svc        PriceFetcher
}

type APIFunc func(context.Context, http.ResponseWriter, *http.Request) error

func NewJSONAPIServer(addr string, svc PriceFetcher) *JSONAPIServer {
	return &JSONAPIServer{
		listenAddr: addr,
		svc:        svc,
	}
}

func (s *JSONAPIServer) Run() {
	// start http server
	http.HandleFunc("/", makeHTTPAPIFunc(s.handleFetchPrice))

	http.ListenAndServe(s.listenAddr, nil)

}
func makeHTTPAPIFunc(apiFn APIFunc) http.HandlerFunc {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "requestID", rand.Intn(10000000000))
	return func(w http.ResponseWriter, r *http.Request) {
		if err := apiFn(ctx, w, r); err != nil {
			// handle error
			writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		}

	}
}

func (s *JSONAPIServer) handleFetchPrice(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	// fetch price from service
	ticker := r.URL.Query().Get("ticker")
	price, err := s.svc.FetchPrice(ctx, ticker)
	if err != nil {
		return err
	}
	// write response
	priceResp := mytypes.PriceResponse{
		Ticker: ticker,
		Price:  price,
	}
	return writeJSON(w, http.StatusOK, &priceResp)

}

func writeJSON(w http.ResponseWriter, s int, v any) error {
	// write json response
	w.WriteHeader(s)
	return json.NewEncoder(w).Encode(v)

}
