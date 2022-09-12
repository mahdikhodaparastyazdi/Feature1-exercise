package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
)

type Order struct {
	Orderid int    `json:"order_id"`
	Price   int    `json:"price"`
	Title   string `json:"title"`
}

var orders []Order

func createOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var order Order
	_ = json.NewDecoder(r.Body).Decode(&order)
	payload, _ := json.Marshal(order)

	orders = append(orders, order)

	if err := redisClient.Publish(ctx, "send-user-data", payload).Err(); err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(order)
}

var ctx = context.Background()

var redisClient = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
	DB:   0,
})

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/order", createOrder).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", r))

}
