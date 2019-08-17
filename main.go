package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net"
	"strings"
)

func main() {
	l, err := net.Listen("tcp", "127.0.0.1:2525")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	fmt.Fprintf(conn, "220 localhost ESMTP talaria\r\n")

	var linelimit int64 = 1 << 10 // KiB
	// Limit reader
	l := io.LimitedReader{
		R: conn,
		N: linelimit,
	}
	buf := bufio.NewReader(&l)

	for {
		l.N = linelimit
		cmd, err := readCmd(buf)
		if err != nil {
			fmt.Println("Failed to read cmd:", err)
			return
		}

		switch cmd.Type {
		case "EHLO":
			handleEhlo(conn)
		case "AUTH":
			handleAuth(conn)
		case "MAIL":
			handleMail(conn)
		case "RCPT":
			handleRcpt(conn)
		case "DATA":
			handleData(conn)
		default:
			fmt.Printf("Received command: %v\r\n", cmd)
		}
	}
}

type Cmd struct {
	Type string
	Args []string
}

func readCmd(rw *bufio.Reader) (*Cmd, error) {
	str, err := rw.ReadString('\n')
	if err != nil {
		return nil, err
	}
	if !strings.HasSuffix(str, "\r\n") {
		return nil, errors.New("Invalid command, must end in \\r\\n")
	}
	str = strings.TrimSuffix(str, "\r\n")

	parts := strings.Split(str, " ")
	if len(parts) == 0 {
		return nil, errors.New("Wat")
	}
	return &Cmd{parts[0], parts[1:]}, nil
}

func handleEhlo(w io.Writer) error {
	io.WriteString(w, "250-localhost gimme the mail\r\n")
	io.WriteString(w, "250 AUTH PLAIN\r\n")
	return nil
}

func handleAuth(w io.Writer) error {
	io.WriteString(w, "235 2.7.0 Authentication successful\r\n")
	return nil
}

func handleMail(w io.Writer) {
	io.WriteString(w, "250 Ok\r\n")
}

func handleRcpt(w io.Writer) {
	io.WriteString(w, "250 Ok\r\n")
}

func handleData(w io.Writer) {
	io.WriteString(w, "354 do your thing\r\n")
}
