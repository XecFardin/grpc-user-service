# gRPC User Service

This is a Golang gRPC service that manages user details and includes search functionality.

## Building and Running the Application

### Prerequisites

- Golang 1.18 or higher
- Docker (optional, for containerization)

### Steps

1. Clone the repository:

    ```sh
    git clone https://github.com/yourusername/grpc-user-service.git
    cd grpc-user-service
    ```

2. Build the application:

    ```sh
    go build -o main .
    ```

3. Run the application:

    ```sh
    ./main
    ```

4. To run the application using Docker:

    ```sh
      sudo docker compose up --build -d
      sudo docker logs grpc-user-service -f
    --output : Server is listening on port 50051
    --to stop
    sudo docker compose down
    ```

## Accessing the gRPC Service Endpoints

The service listens on port `50051`. You can use any gRPC client to interact with the service.

## Downloading the grpcurl in linux
wget https://github.com/fullstorydev/grpcurl/releases/download/v1.8.7/grpcurl_1.8.7_linux_x86_64.tar.gz

tar -xzf grpcurl_1.8.7_linux_x86_64.tar.gz

sudo mv grpcurl /usr/local/bin/

grpcurl --version

### Example using `grpcurl` in terminal

1. Get User by ID:

    ```sh
    grpcurl -plaintext -d '{"id": 1}' localhost:50051 user.UserService/GetUserByID
    ```

2. Get Users by IDs:

    ```sh
    grpcurl -plaintext -d '{"ids": [1, 2]}' localhost:50051 user.UserService/GetUsersByIDs
    ```

3. Search Users:

    ```sh
    grpcurl -plaintext -d '{"city": "LA"}' localhost:50051 user.UserService/SearchUsers
    ```
note: If you are using different go versions kindly do changes in you dockerfile 