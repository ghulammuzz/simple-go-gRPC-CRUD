<!-- Repository Name: Learn-gRPC-CRUD -->

## Description
This repository contains a simple CRUD (Create, Read, Update, Delete) gRPC (Google Remote Procedure Call) service implemented in Go, using SQLite as the database. The service provides basic operations on a `people` table, allowing users to interact with person records by performing CRUD operations.

## Contents:

### main.go:

- Implements the gRPC server with CRUD operations.
- Defines protobuf service and messages using the `protobuff` package.
- Establishes a connection to an SQLite database (`crud.db`).
- Provides methods for handling gRPC requests such as `GetAgeById`, `CreatePerson`, `ReadPerson`, `UpdatePerson`, and `DeletePerson`.

### crud.db:

- SQLite database file where person records are stored.

### client.go
- Implements a Go client for the gRPC CRUD service.
- Utilizes the gRPC library and `learn-grpc/protobuff` package for communication with the server.
- Includes functions for performing CRUD operations:
    - `createPerson`: Creates a new person record.
    - `readPerson`: Reads details of a person by their ID.
    - `updatePerson`: Updates details of a person.
    - `deletePerson`: Deletes a person by their ID.
    - `getAgeById`: Retrieves the age of a person by their ID.

## Usage:

1. Clone the Repository:

    ```bash
    git clone https://github.com/ghulammuzz/simple-go-gRPC-CRUD.git
    cd simple-go-gRPC-CRUD
    ```
2. Run the gRPC Server:

    - Make sure you have Go installed on your machine.
    ```bash
    go run main.go
    ```

3. Run the gRPC Client (Other Terminal)

    - Execute the client application to interact with the gRPC server.
    ```bash
    go run client/client.go
    ```

4. gRPC Methods:
    - The service exposes the following gRPC methods:
        - `GetAgeById`: Get the age of a person by their ID.
        - `CreatePerson`: Create a new person record.
        - `ReadPerson`: Read details of a person by their ID.
        - `UpdatePerson`: Update details of a person.
        - `DeletePerson`: Delete a person by their ID.
    - Uncomment the specific method calls (e.g., createPerson, readPerson) in the main function of client.go to test various CRUD operations.
## Dependencies:

- Go
- SQLite3
- github.com/mattn/go-sqlite3 (SQLite3 driver for Go)
- google.golang.org/grpc (gRPC package)

### Notes:

1. The database schema includes a `people` table with columns for `ID`, `name`, and `age`.
2. The service listens on port `50051` by default.
3. The `createTableIfNotExists` function ensures that the required table exists in the SQLite database.

### Caution:

- This example uses an SQLite database and is primarily intended for learning purposes. In a production environment, consider using a more robust database solution.
- Feel free to explore and modify the code to suit your needs or extend the functionality of the gRPC CRUD service.