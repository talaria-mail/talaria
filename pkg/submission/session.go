package submission

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

const (
	LINELIMIT = 1 << 10 // 1 KiB limit to size of commands
	BODYLIMIT = 1 << 20 // 10 MiB limit to size of message body
)

const (
	stateStart int = 0
	stateEhlo      = iota
	stateAuthenticated
	stateMail
	stateRcpt
	stateData
)

type Session struct {
	// Raw connection
	conn io.ReadWriteCloser

	// Limit reader wrapped on top of raw connection
	// to migirate DOS and message sizing limits
	lmt io.LimitedReader

	// Buffered reader on top of limit reader to
	// peak at commands
	buf *bufio.Reader

	authenticated bool
	state         int

	// Message state
	from       string
	recipients []string
	msg        []byte
}

func (s *Session) Run() {
	defer s.conn.Close()
	s.Greet()

	for {
		s.lmt.N = LINELIMIT
		cmd, err := s.buf.Peek(4)
		if err != nil {
			// TODO: recover from session panic
			panic(err)
		}

		// Must accept commands in a case insensitive fashion
		// https://tools.ietf.org/html/rfc5321#section-2.4
		cmdstr := strings.ToUpper(string(cmd))

		switch cmdstr {
		case "EHLO":
			s.Ehlo()

		case "NOOP":
			// https://tools.ietf.org/html/rfc5321#section-4.1.1.9
			s.Noop()

		case "QUIT":
			// https://tools.ietf.org/html/rfc5321#section-4.1.1.10
			s.Quit()

			// Note that we're quiting whether or not the QUIT command was
			// well formed.
			return

		default:
			s.BadRequest()
		}
	}
}

func (s *Session) Greet() {
	// FIXME: Inject correct hostname
	const hostname = "localhost"
	fmt.Fprintf(s.conn, "220 %s ESMTP talaria\r\n", hostname)
}

func (s *Session) Ehlo() {
	if s.state != stateStart {
		fmt.Fprintf(s.conn, "503 Talaria only accepts EHLO at the start of a session")
		return
	}

	cmd, err := s.buf.ReadString('\n')
	switch {
	case err == io.EOF:
		fmt.Fprintf(s.conn, "500 Who there, line is too long\r\n")

	case err != nil:
		fmt.Fprintf(s.conn, "500 I don't understand\r\n")

	case !strings.HasSuffix(cmd, "\r\n"):
		fmt.Fprintf(s.conn, "501 EHLO must end with CRLF")

	default:
		// FIXME: inject the right hostname
		const hostname = "localhost"
		fmt.Fprintf(s.conn, "250-%s Gimme you mail\r\n", hostname)
		fmt.Fprintf(s.conn, "250 AUTH PLAIN\r\n")
		s.state = stateEhlo
	}
}

func (s *Session) Noop() {
	cmd, err := s.buf.ReadString('\n')
	switch {
	case err == io.EOF:
		fmt.Fprintf(s.conn, "500 Whoa there, line is too long\r\n")

	case err != nil:
		fmt.Fprintf(s.conn, "500 I don't even know how you failed here\r\n")

	case !strings.HasSuffix(cmd, "\r\n"):
		fmt.Fprintf(s.conn, "501 You messed up a quit? (shakes head)\r\n")

	default:
		fmt.Fprintf(s.conn, "250 Ok\r\n")
	}
}

func (s *Session) Quit() {
	cmd, err := s.buf.ReadString('\n')
	switch {
	case err == io.EOF:
		fmt.Fprintf(s.conn, "500 Whoa there, line is too long\r\n")

	case err != nil:
		fmt.Fprintf(s.conn, "500 I don't even know how you failed here\r\n")

	case strings.ToUpper(cmd) != "QUIT\r\n":
		fmt.Fprintf(s.conn, "501 You messed up a quit? (shakes head)\r\n")

	default:
		fmt.Fprintf(s.conn, "221 Bye\r\n")
	}
}

func (s *Session) BadRequest() {
}
