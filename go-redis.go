package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type Student struct {
	Name    string `json:"name"`
	RollNum int    `json:"rollNum"`
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	ctx := context.Background()

	pong, err := client.Ping(ctx).Result()
	fmt.Println(pong, err)

	stud, err := json.Marshal(Student{Name: "Varun", RollNum: 1})
	if err != nil {
		fmt.Println(err)
	}

	err = client.Set(ctx, "student", stud, 0).Err()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Value set succesfully")
	}

	val, err := client.Get(ctx, "student").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)

}
