package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jasonsites/gosk-grpc/internal/protos"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("starting test client...")

	address := "localhost:50051"
	cc, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("error connecting to grpc server: %v\n", err)
	}

	defer func() {
		cc.Close()
		fmt.Println("test client closed")
	}()

	client := protos.NewDomainClient(cc)

	req := protos.DomainRequest{
		Prop: &protos.RequestProperty{Id: 1},
	}

	c, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.Action(c, &req)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	} else {
		fmt.Printf("response: %v\n", res)
	}
}
