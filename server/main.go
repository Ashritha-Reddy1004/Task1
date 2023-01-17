package main

//importing required packages
import (
	"context"
	"fmt"
	"log"
	"main/proto_pb"
	"net"
	"os"
	"os/signal"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

// To handle errors
func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type server struct {
	proto_pb.UnimplementedUserServiceServer
}

// Server side user items
type user_item struct {
	Id    primitive.ObjectID `bson:"_id,omitempty"`
	Name  string             `bson:"name"`
	Email string             `bson:"email"`
	Phone int64              `bson:"phone"`
}

// server side activity items
type activity_item struct {
	Id           primitive.ObjectID `bson:"_id,omitempty"`
	ActivityType string             `bson:"activity_type"`
	Duration     int32              `bson:"duration"`
	Label        string             `bson:"label"`
	Timestamp    string             `bson:"timestamp"`
	Name         string             `bson:"name"`
}

// Function to push the user [server side]
func PushUser(ctx context.Context, item user_item) string {
	Name := item.Name
	filter := bson.M{
		"Name": Name,
	}

	var result_data []user_item
	cursor, err := user_collection.Find(context.TODO(), filter)
	handleError(err)

	cursor.All(context.Background(), &result_data)

	if len(result_data) != 0 {
		result := "User already exists"
		return result
	}

	user_collection.InsertOne(ctx, item)
	result := "User added"
	return result
}

// Function to push user activity [server side]
func PushActivity(ctx context.Context, item activity_item) string {
	activity_collection.InsertOne(ctx, item)
	result := "Activity inserted"
	return result
}

// Function to add users [server side]
func (*server) AddUser(ctx context.Context, req *proto_pb.UserReq) (*proto_pb.UserRes, error) {
	fmt.Println(req)
	name := req.GetUser().GetName()
	email := req.GetUser().GetEmail()
	phone := req.GetUser().GetPhone()

	newUserItem := user_item{
		Name:  name,
		Email: email,
		Phone: phone,
	}
	dbres := PushUser(ctx, newUserItem)
	result := fmt.Sprintf("%v", dbres)

	userAddResponse := proto_pb.UserRes{
		Result: result,
	}
	return &userAddResponse, nil
}

// Function to add user activity [server side]
func (*server) AddActivity(ctx context.Context, req *proto_pb.ActivityReq) (*proto_pb.ActivityRes, error) {
	fmt.Println(req)
	activity_type := req.GetActivity().GetActivityType()
	duration := req.GetActivity().GetDuration()
	label := req.GetActivity().GetLabel()
	timestamp := req.GetActivity().GetTimeStamp()
	name := req.GetActivity().GetName()
	newActivityItem := activity_item{
		ActivityType: activity_type,
		Duration:     duration,
		Label:        label,
		Timestamp:    timestamp,
		Name:         name,
	}
	dbres := PushActivity(ctx, newActivityItem)
	result := fmt.Sprintf("%v", dbres)

	AddActivityRes := proto_pb.ActivityRes{
		Result: result,
	}

	return &AddActivityRes, nil
}

// Function to verify whether the activity is done or not [server side]
func (*server) IsDone(ctx context.Context, req *proto_pb.IsDoneReq) (*proto_pb.IsDoneRes, error) {
	name := req.Name
	ActivityType := req.ActivityType

	fmt.Println("ActivityIsDone() is working:", name, ActivityType)

	result := &proto_pb.IsDoneRes{
		Result: true,
	}

	return result, nil
}

// Function to update user records
func (*server) UpdateUser(ctx context.Context, req *proto_pb.UpdateReq) (*proto_pb.UpdateRes, error) {
	fmt.Println(req)
	email := req.GetUser().GetEmail()
	name := req.GetUser().GetName()
	phone := req.GetUser().GetPhone()

	filter := bson.M{
		"Name": name,
	}

	update := bson.D{{"$set", bson.D{{"email", email}, {"name", name}, {"phone", phone}}}}
	_, err := user_collection.UpdateOne(context.Background(), filter, update)

	handleError(err)
	result := "User details updated"

	updateUserResponse := proto_pb.UpdateRes{
		Result: result,
	}
	return &updateUserResponse, nil
}

// Function to remove user records
func (*server) RemoveUser(ctx context.Context, req *proto_pb.RemoveUserReq) (*proto_pb.RemoveUserRes, error) {
	fmt.Println(req)
	Name := req.GetName()
	filter := bson.M{
		"Name": Name,
	}
	r, err := user_collection.DeleteOne(context.TODO(), filter)
	handleError(err)
	var result string
	if r.DeletedCount == 0 {
		result = "User record does not exist"
	} else {
		result = "User record deleted"
	}
	removeUserResponse := proto_pb.RemoveUserRes{
		Result: result,
	}
	return &removeUserResponse, nil
}

// Function to print the activity details
func (*server) PrintActivity(ctx context.Context, req *proto_pb.PrintActivityReq) (*proto_pb.PrintActivityRes, error) {
	fmt.Println(req)
	name := req.GetName()
	filter := bson.M{
		"name": name,
	}
	var result_data []activity_item
	cursor, err := activity_collection.Find(context.TODO(), filter)
	handleError(err)
	cursor.All(context.Background(), &result_data)
	if len(result_data) == 0 {
		getActivityResponse := proto_pb.PrintActivityRes{
			Status:   false,
			Activity: nil,
		}
		return &getActivityResponse, nil
	} else {
		getActivityResponse := proto_pb.PrintActivityRes{
			Status: true,
			Activity: &proto_pb.Activity{
				ActivityType: result_data[0].ActivityType,
				TimeStamp:    result_data[0].Timestamp,
				Duration:     result_data[0].Duration,
				Label:        result_data[0].Label,
			},
		}
		return &getActivityResponse, nil
	}
}

// Function to verify whether the activity to check it is valid or not [server side]
func (*server) IsValid(ctx context.Context, req *proto_pb.IsValidReq) (*proto_pb.IsValidRes, error) {
	fmt.Println(req)
	name := req.GetName()
	activity_type := req.GetActivityType()
	filter := bson.M{
		"name":          name,
		"activity_type": activity_type,
	}
	var result_data []activity_item
	cursor, err := activity_collection.Find(context.Background(), filter)
	handleError(err)
	cursor.All(context.Background(), &result_data)
	var result string
	if len(result_data) == 0 {
		result = "Details does not exist"
	} else {
		if result_data[0].Duration > 2 {
			result = "Valid Activity"
		} else {
			result = "Invalid Activity"
		}
	}

	IsValidRes := proto_pb.IsValidRes{
		Result: result,
	}
	return &IsValidRes, nil
}

// To get the environment variables
func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")

	handleError(err)

	return os.Getenv(key)
}

// Print user details
func (*server) PrintUser(ctx context.Context, req *proto_pb.PrintUserReq) (*proto_pb.PrintUserRes, error) {
	fmt.Println(req)
	Name := req.GetName()
	filter := bson.M{
		"Name": Name,
	}
	var result_data []user_item
	cursor, err := user_collection.Find(context.TODO(), filter)
	handleError(err)
	cursor.All(context.Background(), &result_data)
	if len(result_data) == 0 {
		getUserResponse := proto_pb.PrintUserRes{
			Status: false,
			User:   nil,
		}
		return &getUserResponse, nil
	} else {
		getUserResponse := proto_pb.PrintUserRes{
			Status: true,
			User: &proto_pb.User{
				Email: result_data[0].Email,
				Name:  result_data[0].Name,
				Phone: result_data[0].Phone,
			},
		}
		return &getUserResponse, nil
	}
}

// main function
var user_collection *mongo.Collection
var activity_collection *mongo.Collection

func main() {
	godotenv.Load(".env")

	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)

	proto_pb.RegisterUserServiceServer(s, &server{})

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	handleError(err)

	go func() {
		if err := s.Serve(lis); err != nil {
			handleError(err)
		}
	}()
	//To establish remote connection
	mongo_uri := goDotEnvVariable("MONGODB_URI")
	client, err := mongo.NewClient(options.Client().ApplyURI(mongo_uri))
	handleError(err)

	fmt.Println("MongoDB Session Started")

	err = client.Connect(context.TODO())
	handleError(err)

	user_collection = client.Database("useractivity").Collection("userdata")

	activity_collection = client.Database("UserActivity").Collection("UserActivityData")

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

	fmt.Println("MongoDB Session Ended")
	if err := client.Disconnect(context.TODO()); err != nil {
		handleError(err)
	}

	s.Stop()

}
