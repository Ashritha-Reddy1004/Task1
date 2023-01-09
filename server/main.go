package main

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

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type server struct {
	proto_pb.UnimplementedUserServiceServer
}

type user_item struct {
	Id    primitive.ObjectID `bson:"_id,omitempty"`
	Name  string             `bson:"name"`
	Email string             `bson:"email"`
	Phone int64              `bson:"phone"`
}

type activity_item struct {
	Id           primitive.ObjectID `bson:"_id,omitempty"`
	ActivityType string             `bson:"activity_type"`
	Duration     int32              `bson:"duration"`
	Label        string             `bson:"label"`
	Timestamp    string             `bson:"timestamp"`
	Email        string             `bson:"email"`
}

func PushUser(ctx context.Context, item user_item) string {
	Name := item.Name
	filter := bson.M{
		"Name": Name,
	}

	var result_data []user_item
	cursor, err := collection.Find(context.TODO(), filter)
	handleError(err)

	cursor.All(context.Background(), &result_data)

	if len(result_data) != 0 {
		result := "User already exists"
		return result
	}

	collection.InsertOne(ctx, item)
	result := "User added"
	return result
}

func PushActivity(ctx context.Context, item activity_item) string {
	collection.InsertOne(ctx, item)
	result := "Activity inserted"
	return result
}

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

func (*server) AddActivity(ctx context.Context, req *proto_pb.ActivityReq) (*proto_pb.ActivityRes, error) {
	fmt.Println(req)
	activity_type := req.GetActivity().GetActivityType()
	duration := req.GetActivity().GetDuration()
	label := req.GetActivity().GetLabel()
	timestamp := req.GetActivity().GetTimeStamp()
	// email := req.GetActivity().GetEmail()
	newActivityItem := activity_item{
		ActivityType: activity_type,
		Duration:     duration,
		Label:        label,
		Timestamp:    timestamp,
		// Email:        email,
	}
	dbres := PushActivity(ctx, newActivityItem)
	result := fmt.Sprintf("%v", dbres)

	AddActivityRes := proto_pb.ActivityRes{
		Result: result,
	}

	return &AddActivityRes, nil

}
func (*server) IsDone(ctx context.Context, req *proto_pb.IsDoneReq) (*proto_pb.IsDoneRes, error) {
	name := req.Name
	ActivityType := req.ActivityType

	fmt.Println("ActivityIsDone() is working:", name, ActivityType)

	result := &proto_pb.IsDoneRes{
		Result: true,
	}

	return result, nil
}

func (*server) IsValid(ctx context.Context, req *proto_pb.IsValidReq) (*proto_pb.IsValidRes, error) {
	fmt.Println(req)
	name := req.GetName()
	activity_type := req.GetActivityType()
	filter := bson.M{
		"name":          name,
		"activity_type": activity_type,
	}
	var result_data []activity_item
	cursor, err := collection.Find(context.Background(), filter)
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

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")

	handleError(err)

	return os.Getenv(key)
}

var collection *mongo.Collection

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

	mongo_uri := goDotEnvVariable("MONGODB_URI")
	client, err := mongo.NewClient(options.Client().ApplyURI(mongo_uri))
	handleError(err)

	fmt.Println("MongoDB Session Started")

	err = client.Connect(context.TODO())
	handleError(err)

	collection = client.Database("useractivity").Collection("useractivitydata")

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

	fmt.Println("MongoDB Session Ended")
	if err := client.Disconnect(context.TODO()); err != nil {
		handleError(err)
	}

	s.Stop()

}
