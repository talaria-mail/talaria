package main

import (
	"fmt"

	"code.nfsmith.ca/talaria/pkg/imap"
	"code.nfsmith.ca/talaria/pkg/submission"
)

func main() {
	var sub submission.Server
	sub.Addr = "localhost:2525"

	var imapServer imap.Server
	imapServer = "localhost:8143"

	fmt.Println(sub.ListenAndServe())
}
