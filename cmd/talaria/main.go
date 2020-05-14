package main

import (
	"fmt"

	"github.com/nsmith5/talaria/pkg/submission"
)

func main() {
	var sub submission.Server
	sub.Addr = "0.0.0.0:2525"

	fmt.Println("Binding submission server to port :2525....")
	sub.ListenAndServe()
}
