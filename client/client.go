package main

import (
	"context"
	"log"
	"time"

	pb "learn-grpc/protobuff"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const address = "localhost:50051"

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewCrudServiceClient(conn)

	// createPerson(client)
	age := getAgeById(client, 2)

	log.Println(age)
	// readPerson(client)
	// updatePerson(client)
	// deletePerson(client)
}
func getAgeById(client pb.CrudServiceClient, id int64) int32 {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	readRequest := &pb.ReadRequest{Id: id}

	response, err := client.GetAgeById(ctx, readRequest)
	if err != nil {
		log.Fatalf("Error getting age: %v", err)
	}

	if response != nil {
		log.Printf("Age of person with ID %d: %d", id, response.Age)
	} else {
		log.Printf("Person with ID %d not found", id)
	}

	return response.Age
}

func createPerson(client pb.CrudServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	person := &pb.Person{Name: "John Doe", Age: 30}

	response, err := client.CreatePerson(ctx, person)
	if err != nil {
		log.Fatalf("Error creating person: %v", err)
	}

	log.Printf("Created Person: %v", response)
}

func readPerson(client pb.CrudServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	readRequest := &pb.ReadRequest{Id: 1}

	response, err := client.ReadPerson(ctx, readRequest)
	if err != nil {
		log.Fatalf("Error reading person: %v", err)
	}

	log.Printf("Read Person: %v", response)
}

func updatePerson(client pb.CrudServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	person := &pb.Person{Id: 1, Name: "Updated John Doe", Age: 35}

	response, err := client.UpdatePerson(ctx, person)
	if err != nil {
		log.Fatalf("Error updating person: %v", err)
	}

	log.Printf("Updated Person: %v", response)
}
func deletePerson(client pb.CrudServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	deleteRequest := &pb.ReadRequest{Id: 1}

	response, err := client.DeletePerson(ctx, deleteRequest)
	if err != nil {
		log.Fatalf("Error deleting person: %v", err)
	}

	log.Printf("Delete Response: %v", response)
}
