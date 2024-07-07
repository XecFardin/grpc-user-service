package server

import (
	"context"
	"fmt"

	pb "github.com/XecFardin/grpc-user-service/proto"
)

type UserServiceServer struct {
	pb.UnimplementedUserServiceServer
}

var users = []pb.User{
	{Id: 1, Name: "Abdulla", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
	{Id: 2, Name: "Safwan", City: "NY", Phone: 1234567891, Height: 5.9, Married: false},
	{Id: 3, Name: "Ibtisaam", City: "SF", Phone: 1234567892, Height: 5.7, Married: true},
	{Id: 4, Name: "Adesh", City: "LA", Phone: 1234567893, Height: 5.6, Married: false},
}

func (s *UserServiceServer) GetUserByID(ctx context.Context, req *pb.UserIDRequest) (*pb.UserResponse, error) {
	for _, user := range users {
		if user.Id == req.Id {
			return &pb.UserResponse{User: &user}, nil
		}
	}
	return nil, fmt.Errorf("User not found")
}

func (s *UserServiceServer) GetUsersByIDs(ctx context.Context, req *pb.UserIDsRequest) (*pb.UsersResponse, error) {
	var result []*pb.User
	for _, id := range req.Ids {
		for _, user := range users {
			if user.Id == id {
				u := user // Copy user to avoid pointer reference issues
				result = append(result, &u)
			}
		}
	}
	return &pb.UsersResponse{Users: result}, nil
}

func (s *UserServiceServer) SearchUsers(ctx context.Context, req *pb.SearchRequest) (*pb.UsersResponse, error) {
	fmt.Printf("Received SearchUsers request: %+v\n", req)
	var result []*pb.User
	for _, user := range users {
		matches := true
		if req.City != "" && user.City != req.City {
			matches = false
		}
		if req.Phone != 0 && user.Phone != req.Phone {
			matches = false
		}
		if req.Married != nil && user.Married != req.Married.Value {
			matches = false
		}
		if matches {
			u := user // Copy user to avoid pointer reference issues
			result = append(result, &u)
		}
	}
	return &pb.UsersResponse{Users: result}, nil
}
