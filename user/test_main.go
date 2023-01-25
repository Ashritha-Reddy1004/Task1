package main

import (
	"fmt"
	"main/activity_pb"
	"testing"

	"google.golang.org/grpc"
)




func check_bool(t *testing.T, got bool, want bool) {
	if got != want {
		s := fmt.Sprint("got", got, ", wanted", want)
		t.Errorf(s)
	}
}

func connectToDB() (proto_pb.UserServiceClient, *grpc.ClientConn) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	handleError(err)
	a := proto_pb.NewUserServiceClient(conn)
	return a, conn
}

func TestUserAdd(t *testing.T) {
	a, conn := connectToServer()
	defer conn.Close()

	got := AddUser(a, "test1", "test1@gmail.com", 73867912566)
	want := "User already exist"
	check_string(t, got, want)
	got = AddUser(a, "Test2", "test2@gmail.com", 11111111)
	want = "User already exist"
	check_string(t, got, want)
  got = AddUser(a, "Test3", "test3@outllok.com", 187851214564561)
	want = "User already exist"
	check_string(t, got, want)
}

func TestActivityAdd(t *testing.T) {
	a, conn := connectToServer()
	defer conn.Close()

	got := AddActivity(a, "test1@gmail.com", "Sleep", 7, "label4")
	want := "User activity added successfully"
	check_string(t, got, want)
}

func TestprintUser(t *testing.T) {
	a, conn := connectToServer()
	defer conn.Close()

	got := PrintUser(a, "test1@outlook.com")
	want := true
	check_bool(t, got, want)
	got = PrintUser(a, "unknownuser@gmail.com")
	want = false
	check_bool(t, got, want)
  got = PrintUser(a, "anonymous@gmail.com")
	want = false
	check_bool(t, got, want)
}

func TestPrintActivity(t *testing.T) {
	a, conn := connectToServer()
	defer conn.Close()

	got := PrintActivity(a, "test1@gmail.com")
	want := true
  check_bool(t, got, want)
  got := PrintActivity(a, "test2@gmail.com")
	want := true
  check_bool(t, got, want)
  got := PrintActivity(a, "anonymous@gmail.com")
	want := false
  check_bool(t, got, want)
}

func TestUpdateUser(t *testing.T) {
	a, conn := connectToServer()
	defer conn.Close()
	name := "ashritha"
	email := "ashrithareddy1004@gmail.com"
	phone := 7894561230
	got := UpdateUser(a, email, name, int64(phone))
	want := "User details updated"
	check_string(t, got, want)
  name := "ankith"
	email := "ankithgmail.com"
	phone := 78782145484
	got := UpdateUser(a, email, name, int64(phone))
	want := "User details updated"
	check_string(t, got, want)
  name := "manasa"
	email := "manasa1506@gmail.com"
	phone := 965478213
	got := UpdateUser(a, email, name, int64(phone))
	want := "User details updated"
	check_string(t, got, want)
}

func TestRemoveUser(t *testing.T) {
	a, conn := connectToServer()
	defer conn.Close()

	got := RemoveUser(a, "test1@gmail.com")
	want := "User deleted successfully"
	check_string(t, got, want)
  got := RemoveUser(a, "test2@gmail.com")
	want := "User deleted successfully"
	check_string(t, got, want)
   got := RemoveUser(a, "test3@gmail.com")
	want := "User deleted successfully"
	check_string(t, got, want)
}
}

func TestIsValid(t *testing.T) {
	a, conn := connectToServer()
	defer conn.Close()

	got := ActivityIsValid(a, "test2@gmail.com", "Play")
	want := "Activity is InValid"
	check_string(t, got, want)
  got = ActivityIsValid(a, "test1@gmail.com", "Study")
	want = "Activity is Valid"
	check_string(t, got, want)
  ot = ActivityIsValid(a, "test3@gmail.com", "Study")
	want = "Activity is InValid"
	check_string(t, got, want)
}


func check_string(t *testing.T, got string, want string) {
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestsDone(t *testing.T) {
	a, conn := connectToServer()
	defer conn.Close()

	got := IsDone(a, "test1@gmail.com", "Read")
	want := "Activity is not yet completed"
	check_string(t, got, want)

	got = ActivityIsDone(a, "test2@gmail.com", "Study")
	want = "Activity is Done"
	check_string(t, got, want)
  
  got = ActivityIsDone(a, "test3@gmail.com", "Sleep")
	want = "Activity is Done"
	check_string(t, got, want)
}
