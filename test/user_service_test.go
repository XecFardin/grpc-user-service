package test

import (
	"context"
	"reflect"
	"testing"

	pb "github.com/XecFardin/grpc-user-service/proto"
	"github.com/XecFardin/grpc-user-service/server"
)

func TestGetUserByID(t *testing.T) {
	server := &server.UserServiceServer{}

	// Test Case 1: Existing User ID
	req1 := &pb.UserIDRequest{Id: 1}
	resp1, err := server.GetUserByID(context.Background(), req1)
	if err != nil {
		t.Errorf("Error fetching user by ID 1: %v", err)
	}
	expectedUser1 := &pb.User{Id: 1, Name: "Abdulla", City: "LA", Phone: 1234567890, Height: 5.8, Married: true}
	if !reflect.DeepEqual(resp1.User, expectedUser1) {
		t.Errorf("Expected user %+v, but got %+v", expectedUser1, resp1.User)
	}

	// Test Case 2: Non-existing User ID
	req2 := &pb.UserIDRequest{Id: 10}
	resp2, err := server.GetUserByID(context.Background(), req2)
	if err == nil || resp2 != nil {
		t.Errorf("Expected error for non-existing user ID 10, but got response: %+v", resp2)
	}
	expectedErr := "User not found"
	if err != nil && err.Error() != expectedErr {
		t.Errorf("Expected error message '%s', but got '%s'", expectedErr, err.Error())
	}
}

func TestGetUsersByIDs(t *testing.T) {
	server := &server.UserServiceServer{}

	// Test Case 1: Existing User IDs
	req1 := &pb.UserIDsRequest{Ids: []int32{1, 3}}
	resp1, err := server.GetUsersByIDs(context.Background(), req1)
	if err != nil {
		t.Errorf("Error fetching users by IDs 1 and 3: %v", err)
	}
	expectedUsers1 := []*pb.User{
		{Id: 1, Name: "Abdulla", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
		{Id: 3, Name: "Ibtisaam", City: "SF", Phone: 1234567892, Height: 5.7, Married: true},
	}
	if !reflect.DeepEqual(resp1.Users, expectedUsers1) {
		t.Errorf("Expected users %+v, but got %+v", expectedUsers1, resp1.Users)
	}

	// Test Case 2: Non-existing User IDs
	req2 := &pb.UserIDsRequest{Ids: []int32{10, 20}}
	resp2, err := server.GetUsersByIDs(context.Background(), req2)
	if err != nil {
		t.Errorf("Error fetching users by non-existing IDs: %v", err)
	}
	if len(resp2.Users) != 0 {
		t.Errorf("Expected empty user list for non-existing IDs, but got %+v", resp2.Users)
	}
}

func TestSearchUsers(t *testing.T) {
	server := &server.UserServiceServer{}
	// Test Case 1: Matching City
	req1 := &pb.SearchRequest{City: "LA"}
	resp1, err := server.SearchUsers(context.Background(), req1)
	if err != nil {
		t.Errorf("Error searching users by city: %v", err)
	}
	expectedUsers1 := []*pb.User{
		{Id: 1, Name: "Abdulla", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
		{Id: 4, Name: "Adesh", City: "LA", Phone: 1234567893, Height: 5.6, Married: false},
	}
	if !reflect.DeepEqual(resp1.Users, expectedUsers1) {
		t.Errorf("Expected users with City 'LA' %+v, but got %+v", expectedUsers1, resp1.Users)
	}

	// Test Case 2: Matching Phone Number
	req2 := &pb.SearchRequest{Phone: 1234567891}
	resp2, err := server.SearchUsers(context.Background(), req2)
	if err != nil {
		t.Errorf("Error searching users by phone number: %v", err)
	}
	expectedUsers2 := []*pb.User{
		{Id: 2, Name: "Safwan", City: "NY", Phone: 1234567891, Height: 5.9, Married: false},
	}
	if !reflect.DeepEqual(resp2.Users, expectedUsers2) {
		t.Errorf("Expected user with Phone number 1234567891 %+v, but got %+v", expectedUsers2, resp2.Users)
	}

	// Test Case 4: Multiple Criteria - City and Phone
	req4 := &pb.SearchRequest{City: "LA", Phone: 1234567890}
	resp4, err := server.SearchUsers(context.Background(), req4)
	if err != nil {
		t.Errorf("Error searching users by city and phone number: %v", err)
	}
	expectedUsers4 := []*pb.User{
		{Id: 1, Name: "Abdulla", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
	}
	if !reflect.DeepEqual(resp4.Users, expectedUsers4) {
		t.Errorf("Expected user with City 'LA' and Phone number 1234567890 %+v, but got %+v", expectedUsers4, resp4.Users)
	}

	// Test Case 6: No Matching Users
	req6 := &pb.SearchRequest{City: "NonExistingCity"}
	resp6, err := server.SearchUsers(context.Background(), req6)
	if err != nil {
		t.Errorf("Error searching users with no matching criteria: %v", err)
	}
	if len(resp6.Users) != 0 {
		t.Errorf("Expected no users for non-existing city, but got %+v", resp6.Users)
	}

	// Test Case 7: Empty Search Request
	req7 := &pb.SearchRequest{}
	resp7, err := server.SearchUsers(context.Background(), req7)
	if err != nil {
		t.Errorf("Error searching all users: %v", err)
	}
	expectedUsers7 := []*pb.User{
		{Id: 1, Name: "Abdulla", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
		{Id: 2, Name: "Safwan", City: "NY", Phone: 1234567891, Height: 5.9, Married: false},
		{Id: 3, Name: "Ibtisaam", City: "SF", Phone: 1234567892, Height: 5.7, Married: true},
		{Id: 4, Name: "Adesh", City: "LA", Phone: 1234567893, Height: 5.6, Married: false},
	}
	if !reflect.DeepEqual(resp7.Users, expectedUsers7) {
		t.Errorf("Expected all users %+v, but got %+v", expectedUsers7, resp7.Users)
	}

	// Test Case 8: Edge Case - Zero Phone Number
	req8 := &pb.SearchRequest{Phone: 0}
	resp8, err := server.SearchUsers(context.Background(), req8)
	if err != nil {
		t.Errorf("Error searching users with zero phone number: %v", err)
	}
	expectedUsers8 := []*pb.User{
		{Id: 1, Name: "Abdulla", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
		{Id: 2, Name: "Safwan", City: "NY", Phone: 1234567891, Height: 5.9, Married: false},
		{Id: 3, Name: "Ibtisaam", City: "SF", Phone: 1234567892, Height: 5.7, Married: true},
		{Id: 4, Name: "Adesh", City: "LA", Phone: 1234567893, Height: 5.6, Married: false},
	}
	if !reflect.DeepEqual(resp8.Users, expectedUsers8) {
		t.Errorf("Expected all users %+v with zero phone number, but got %+v", expectedUsers8, resp8.Users)
	}
}
