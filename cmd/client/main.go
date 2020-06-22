package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"os"
	"simple-adder-grpc-golang/pkg/api"
	"strconv"
)

func main() {
	x, _ := strconv.Atoi(os.Args[1])
	y, _ := strconv.Atoi(os.Args[2])

	x_req := int64(x)
	y_req := int64(y)


	conn, err := grpc.Dial("127.0.0.1:8000", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	client := api.NewAddClient(conn)

	resp, err := client.Add(context.Background(),  &api.AddRequest{X: x_req, Y: y_req})
	fmt.Println(resp)

}