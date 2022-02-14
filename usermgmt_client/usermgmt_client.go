package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/rennerp/go-grpc-tutorial/usermgmt"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("Did not connect: %v", err)

	}

	defer conn.Close()
	c := pb.NewUserManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var new_users = make(map[string]int32)
	new_users["Allice"] = 43
	new_users["Bob"] = 19

	for name, age := range new_users {
		response, error := c.CreateNewUser(ctx, &pb.NewUser{Name: name, Age: age})
		if error != nil {
			log.Fatalf("Could not create users %v", error)
		}

		log.Printf(`User Details:
NAME: %s
AGE: %d
ID: %d`, response.GetName(), response.GetAge(), response.GetId())
	}

	params := &pb.GetUsersParams{}
	response, err := c.GetUsers(ctx, params)

	if err != nil {
		log.Fatalf("Could not retrieve users: %v", err)
	}

	log.Print("\nUSER LIST: \n")
	fmt.Printf("r.getUsers(): %v \n", response.GetUsers())
}
