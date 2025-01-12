package main

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"context"
	"encoding/json"
)

var ctx = context.Background()

type Node struct {
	Name string `json:"name"`
	IP string `json:"ip"`
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
	pong, err:= client.Ping(ctx).Result()
	fmt.Println(pong,err)

	// Writing the Node data into Redis
	// Creating a list of nodes to be appended.
	Node_List := "Node_Details"
	json_body, err := json.Marshal(Node{
		Name: "ubuntuserver",
		IP: "192.168.64.11",
	})
	if err != nil {
		fmt.Println(err)
	}

	_, err = client.RPush(ctx, Node_List, json_body).Result()
	if err != nil {
		fmt.Println(err)
	}

	// Fetch the details of List of Node Details.
	res, err := client.LRange(ctx, Node_List, 0, -1).Result()
	if err != nil {
		fmt.Println(err)
	}
	for _,item := range res {
		var node_dict Node
		err := json.Unmarshal([]byte(item),&node_dict)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(node_dict)
	}
}
