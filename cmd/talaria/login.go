package main

import (
	"context"
	"log"

	"github.com/nsmith5/talaria/pkg/auth"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to a server",
	Run:   loginCmdRun,
}

func loginCmdRun(cmd *cobra.Command, args []string) {
	conn, err := grpc.Dial("127.0.0.1:5002", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ac := auth.NewAuthClient(conn)
	req := &auth.LoginRequest{
		Username: "root",
		Password: "passwerd",
	}
	lm, err := ac.Login(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(lm)
}
