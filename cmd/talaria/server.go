package main

import (
	"github.com/spf13/cobra"
	"net/http"
)

//go:generate go-bindata -prefix "../../frontend/dist/" -fs ../../frontend/dist/...

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Talaria server",
	Run:   runServerCmd,
}

func runServerCmd(cmd *cobra.Command, args []string) {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(AssetFile()))
	http.ListenAndServe(":8080", mux)
}
