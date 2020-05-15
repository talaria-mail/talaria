package main

import (
	"log"
	"net"

	"github.com/nsmith5/talaria/pkg/auth"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start talaria server",
	Run:   serverCmdRun,
}

func serverCmdRun(cmd *cobra.Command, args []string) {
	as := auth.NewServer()
	gs := grpc.NewServer()
	auth.RegisterAuthServer(gs, as)

	ln, err := net.Listen("tcp", ":5002")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("GRPC binding to :5002....")
	log.Fatal(gs.Serve(ln))
}
