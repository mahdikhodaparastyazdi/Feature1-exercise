package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type Order struct {
	Orderid int    `json:"order_id"`
	Price   int    `json:"price"`
	Title   string `json:"title"`
}

var ctx = context.Background()

var redisClient = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})

func main() {
	subscriber := redisClient.Subscribe(ctx, "send-user-data")

	order := Order{}
	for {
		msg, err := subscriber.ReceiveMessage(ctx)
		if err != nil {
			panic(err)
		}

		if err := json.Unmarshal([]byte(msg.Payload), &order); err != nil {
			panic(err)
		}
		fmt.Println("Received message from " + msg.Channel + " channel.")
		fmt.Printf("%+v\n", order)
		// ...
	}
}
