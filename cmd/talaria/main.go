package main

import (
	"fmt"

	"code.nfsmith.ca/talaria/pkg/submission"
)

func main() {
	var sub submission.Server
	sub.Addr = "localhost:2525"

	fmt.Println(sub.ListenAndServe())
}
