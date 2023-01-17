/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"log"
	"main/proto_pb"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "user",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

var adduserCmd = &cobra.Command{
	Use:   "adduser",
	Short: "To add the user details",
	Long: `To update existing user details.
	Inputs:
		email, name, phone-number
	Example:
		client updateuser --email=<email> <name> <phone-number>`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
		fmt.Println("The start time is :")
		fmt.Println(GetTimeStamp())
		handleError(err)
		fmt.Println("User session started")
		//Closing the connection
		defer conn.Close()
		a := proto_pb.NewUserServiceClient(conn)
		phone, err := strconv.ParseInt(args[2], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		AddUser(a, args[0], args[1], phone)
	},
}

var updateuserCmd = &cobra.Command{
	Use:   "updateuser",
	Short: "Update user details.",
	Long: `To update existing user details.
Inputs:
	email, name, phone-number
Example:
	client updateuser --email=<email> <name> <phone-number>
	`,
	Run: func(cmd *cobra.Command, args []string) {
		name, err := cmd.Flags().GetString("Name")
		if err != nil {
			log.Fatal(err)
		}
		phone, err := strconv.ParseInt(args[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
		handleError(err)
		defer conn.Close()

		a := proto_pb.NewUserServiceClient(conn)

		UpdateUser(a, name, args[0], phone)
	},
}

var addactCmd = &cobra.Command{
	Use:   "addactivity",
	Short: "To add an activity of user",
	Long: `To add a specific activity of a user.
Activity types : (Play, Eat, Sleep, Study).
Inputs:
	name, activitytype, duration, label
	
Example:
	client addact <name> <activitytype> <duration> <label>
	`,
	Run: func(cmd *cobra.Command, args []string) {

		conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
		handleError(err)
		defer conn.Close()
		a := proto_pb.NewUserServiceClient(conn)
		duration, err := strconv.ParseInt(args[1], 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		AddActivity(a, args[0], int32(duration), args[2], args[3])
	},
}

var isvalidCmd = &cobra.Command{
	Use:   "isvalid",
	Short: "To check whether the user activity is valid or not.",
	Long: `To checkout whether the user activiti is valid or not.
	
Inputs:
	name, activity-type
Example:
	client isvalid --name=<name> <activity-type>
	`,
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
		handleError(err)
		defer conn.Close()
		a := proto_pb.NewUserServiceClient(conn)
		name, err2 := cmd.Flags().GetString("name")
		handleError(err2)
		activityType := args[0]
		IsValid(a, name, activityType)
	},
}

var isdoneCmd = &cobra.Command{
	Use:   "isdone",
	Short: "To check whether the user activity is done or not.",
	Long: `To checkout whether the user activiti is done or not.
	
Inputs:
	name, activity-type
Example:
	client isdone --name=<name> <activity-type>
	`,
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
		handleError(err)
		defer conn.Close()
		a := proto_pb.NewUserServiceClient(conn)
		name, err2 := cmd.Flags().GetString("name")
		handleError(err2)
		activityType := args[0]
		IsDone(a, name, activityType)
	},
}

var remuserCmd = &cobra.Command{
	Use:   "remuser",
	Short: "To delete an existing user.",
	Long: `To delete an existing user by taking email.
	
Input:
	name
Example:
	client deluser --name=<name>
`,
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
		handleError(err)
		defer conn.Close()
		a := proto_pb.NewUserServiceClient(conn)
		name, err := cmd.Flags().GetString("name")
		handleError(err)
		RemoveUser(a, name)
	},
}

var printuserCmd = &cobra.Command{
	Use:   "printuser",
	Short: "To get the user details",
	Long: `To get the user details (name, email, phone) by taking email.
Input: 
	name
Example:
	client getuser --name=<name>`,
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
		handleError(err)
		defer conn.Close()
		a := proto_pb.NewUserServiceClient(conn)
		name, err2 := cmd.Flags().GetString("name")
		handleError(err2)
		PrintUser(a, name)
	},
}

var printactCmd = &cobra.Command{
	Use:   "printactivity",
	Short: "To get activity data of a user",
	Long: `To get activity data of a specific user.
Inputs:
	name
Example:
	client getact --name=<name>
	`,
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
		handleError(err)
		defer conn.Close()
		a := proto_pb.NewUserServiceClient(conn)
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			log.Fatal(err)
		}
		PrintActivity(a, name)

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.user.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.AddCommand(adduserCmd)
	rootCmd.AddCommand(updateuserCmd)
	rootCmd.AddCommand(addactCmd)
	rootCmd.AddCommand(isvalidCmd)
	rootCmd.AddCommand(isdoneCmd)
	rootCmd.AddCommand(remuserCmd)
	rootCmd.AddCommand(printuserCmd)
	rootCmd.AddCommand(printactCmd)
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
