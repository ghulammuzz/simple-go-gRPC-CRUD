package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"

	pb "learn-grpc/protobuff"

	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
)

const dbPath = "crud.db"

type server struct {
	pb.UnimplementedCrudServiceServer
	db *sql.DB
}

func (s *server) GetAgeById(ctx context.Context, req *pb.ReadRequest) (*pb.AgeResponse, error) {
	var age int32

	err := s.db.QueryRow("SELECT age FROM people WHERE id=?", req.Id).Scan(&age)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Person not found with ID: %d", req.Id)
		}
		log.Printf("Error getting age: %v", err)
		return nil, err
	}

	return &pb.AgeResponse{Age: age}, nil
}

func (s *server) CreatePerson(ctx context.Context, req *pb.Person) (*pb.Person, error) {
	result, err := s.db.Exec("INSERT INTO people (name, age) VALUES (?, ?)", req.Name, req.Age)
	if err != nil {
		log.Printf("Error creating person: %v", err)
		return nil, err
	}

	id, _ := result.LastInsertId()
	return &pb.Person{Id: id, Name: req.Name, Age: req.Age}, nil
}

func (s *server) ReadPerson(ctx context.Context, req *pb.ReadRequest) (*pb.Person, error) {
	var person pb.Person
	err := s.db.QueryRow("SELECT id, name, age FROM people WHERE id=?", req.Id).
		Scan(&person.Id, &person.Name, &person.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Person not found with ID: %d", req.Id)
		}
		log.Printf("Error reading person: %v", err)
		return nil, err
	}

	return &person, nil
}

func (s *server) UpdatePerson(ctx context.Context, req *pb.Person) (*pb.Person, error) {
	_, err := s.db.Exec("UPDATE people SET name=?, age=? WHERE id=?", req.Name, req.Age, req.Id)
	if err != nil {
		log.Printf("Error updating person: %v", err)
		return nil, err
	}

	return req, nil
}

func (s *server) DeletePerson(ctx context.Context, req *pb.ReadRequest) (*pb.DeleteResponse, error) {
	result, err := s.db.Exec("DELETE FROM people WHERE id=?", req.Id)
	if err != nil {
		log.Printf("Error deleting person: %v", err)
		return nil, err
	}

	rowsAffected, _ := result.RowsAffected()
	return &pb.DeleteResponse{Success: rowsAffected > 0}, nil
}

func main() {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	if err := createTableIfNotExists(db); err != nil {
		log.Fatalf("Error creating table: %v", err)
	}

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterCrudServiceServer(s, &server{db: db})

	log.Println("Server is listening on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func createTableIfNotExists(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS people (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		age INTEGER NOT NULL
	)`)
	return err
}
