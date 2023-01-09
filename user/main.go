package main

import (
	"context"
	"fmt"
	"log"
	"main/proto_pb"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	fmt.Println("The start time is :")
	fmt.Println(GetTimeStamp())
	handleError(err)
	fmt.Println("User session started")
	defer conn.Close()

	a := proto_pb.NewUserServiceClient(conn)
	AddUser(a)
	UpdateUser(a)
	AddActivity(a)
	IsValid(a)
	IsDone(a)

}
func handleError(err error) {
	if err != nil {
		fmt.Println("Error raised :", err)
		log.Fatal(err)
	}
}
func AddUser(a proto_pb.UserServiceClient) {
	AddUserReq := proto_pb.UserReq{
		User: &proto_pb.User{
			Name:  "ashritha11",
			Email: "ashritha1@gmail.com",
			Phone: 123456789,
		},
	}
	result, err := a.AddUser(context.Background(), &AddUserReq)
	handleError(err)
	fmt.Println(result)
}
func UpdateUser(a proto_pb.UserServiceClient) {
	UpdateUserReq := proto_pb.UpdateReq{
		User: &proto_pb.User{
			Name:  "ashritha",
			Email: "ashritha@gmail.com",
			Phone: 123456789,
		},
	}
	result, err := a.UpdateUser(context.Background(), &UpdateUserReq)
	handleError(err)
	fmt.Println(result)
}
func GetTimeStamp() string {
	Time := time.Now()
	TimeStamp := Time.Format("09-01-2022 23:59:59 Sunday")
	return TimeStamp
}
func AddActivity(a proto_pb.UserServiceClient) {
	AddActivityReq := proto_pb.ActivityReq{
		Activity: &proto_pb.Activity{
			ActivityType: "Play",
			TimeStamp:    GetTimeStamp(),
			Duration:     1,
			Label:        "Playing gives joy to the heart",
			Name:         "Ashritha",
		},
	}

	result, err := a.AddActivity(context.Background(), &AddActivityReq)
	handleError(err)
	fmt.Println(result)
}
func IsValid(a proto_pb.UserServiceClient) {
	IsValidReq := proto_pb.IsValidReq{
		Name:         "Ashritha",
		ActivityType: "Play",
	}
	result, err := a.IsValid(context.Background(), &IsValidReq)
	handleError(err)
	fmt.Println(result)
}
func IsDone(a proto_pb.UserServiceClient) {
	IsDoneReq := proto_pb.IsDoneReq{
		Name:         "Ashritha",
		ActivityType: "Play",
	}
	result, err := a.IsDone(context.Background(), &IsDoneReq)
	handleError(err)
	fmt.Println(result)
}
