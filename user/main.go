package main

//Import required packages
import (
	"context"
	"fmt"
	"log"
	"main/proto_pb"
	"time"

	"google.golang.org/grpc"
)

// Main function to call rest of the functions
func main() {
	//Connection establishment with the database
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	fmt.Println("The start time is :")
	fmt.Println(GetTimeStamp())
	handleError(err)
	fmt.Println("User session started")
	//Closing the connection
	defer conn.Close()

	//a := proto_pb.NewUserServiceClient(conn)
	//calling the created functions
	//AddUser(a)
	// UpdateUser(a)
	// AddActivity(a)
	// IsValid(a)
	// IsDone(a)

}

// Function to handle client side errors
func handleError(err error) {
	if err != nil {
		fmt.Println("Error raised :", err)
		log.Fatal(err)
	}
}

// Function to add users[client side]
func AddUser(a proto_pb.UserServiceClient, name string, email string, phone int64) {
	AddUserReq := proto_pb.UserReq{
		User: &proto_pb.User{
			Name:  name,
			Email: email,
			Phone: phone,
		},
	}
	result, err := a.AddUser(context.Background(), &AddUserReq)
	handleError(err)
	fmt.Println(result)
}

// Function to update the user details [client side]
func UpdateUser(a proto_pb.UserServiceClient, name string, email string, phone int64) {
	UpdateUserReq := proto_pb.UpdateReq{
		User: &proto_pb.User{
			Name:  name,
			Email: email,
			Phone: phone,
		},
	}
	result, err := a.UpdateUser(context.Background(), &UpdateUserReq)
	handleError(err)
	fmt.Println(result)
}

// Function to get the timestamp
func GetTimeStamp() string {
	Time := time.Now()
	TimeStamp := Time.Format("09-01-2022 23:59:59 Sunday")
	return TimeStamp
}

// Function to add the user activity[client side]
func AddActivity(a proto_pb.UserServiceClient, activitytype string, duration int32, label string, name string) {
	AddActivityReq := proto_pb.ActivityReq{
		Activity: &proto_pb.Activity{
			ActivityType: activitytype,
			TimeStamp:    GetTimeStamp(),
			Duration:     1,
			Label:        label,
			Name:         name,
		},
	}

	result, err := a.AddActivity(context.Background(), &AddActivityReq)
	handleError(err)
	fmt.Println(result)
}

// Function to validate the user activity[client side]
func IsValid(a proto_pb.UserServiceClient, name string, activitytype string) {
	IsValidReq := proto_pb.IsValidReq{
		Name:         name,
		ActivityType: activitytype,
	}
	result, err := a.IsValid(context.Background(), &IsValidReq)
	handleError(err)
	fmt.Println(result)
}

// Function to check whether the given activity is completed or not
func IsDone(a proto_pb.UserServiceClient, name string, activitytype string) {
	IsDoneReq := proto_pb.IsDoneReq{
		Name:         name,
		ActivityType: activitytype,
	}
	result, err := a.IsDone(context.Background(), &IsDoneReq)
	handleError(err)
	fmt.Println(result)
}

func PrintActivity(a proto_pb.UserServiceClient, name string) {
	getActivityRequest := proto_pb.PrintActivityReq{
		Name: name,
	}
	res, err := a.PrintActivity(context.Background(), &getActivityRequest)
	handleError(err)
	fmt.Println(res)
}

func PrintUser(a proto_pb.UserServiceClient, name string) {
	getUserRequest := proto_pb.PrintUserReq{
		Name: name,
	}
	res, err := a.PrintUser(context.Background(), &getUserRequest)
	handleError(err)
	fmt.Println(res)
}

func RemoveUser(a proto_pb.UserServiceClient, name string) {
	removeUserRequest := proto_pb.RemoveUserReq{
		Name: name,
	}
	res, err := a.RemoveUser(context.Background(), &removeUserRequest)
	handleError(err)
	fmt.Println(res)
}
