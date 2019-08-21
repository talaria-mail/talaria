package imap

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"runtime/debug"
	"strings"
)

const (
	// 1 KiB limit to the size of and command
	LINELIMIT = 1 << 10
)

// State of the IMAP submission session
type State int

const (
	NotAuthenticated State = 0
	Authenticated          = iota
	Selected
)

func ParseLine(r *bufio.Reader) (tag, cmd string, args []string, err error) {
	var line string

	line, err = r.ReadString('\n')
	if err != nil {
		return
	}

	if !strings.HasSuffix(line, "\r\n") {
		err = errors.New(`Line doesn't end in \r\n`)
		return
	}

	line = strings.TrimSuffix(line, "\r\n")
	parts := strings.Split(line, " ")
	if len(parts) == 1 {
		err = errors.New(`Line needs a tag and command`)
		return
	}

	return parts[0], strings.ToUpper(parts[1]), parts[2:], nil
}

// Session is a single TCP connection containing mail submission dialog.
//
// NOTE: There are /no/ network related entities here. The session / server
// boundary is an abstraction boundry for network related things. Network
// specific details (e.g. Read and write timeouts) are all operated at the
// server level and all protocol specific things happen at the session
// level.
type Session struct {
	// Raw connection
	conn io.ReadWriteCloser

	// Limit reader wrapped on top of raw connection
	// to migirate DOS and message sizing limits
	lmt io.LimitedReader

	// Buffered reader on top of limit reader to
	// peak at commands
	buf *bufio.Reader

	hostname string
	state    State
}

// Run initiates a mail submission session.
func (s *Session) Run() {
	// Any panics during a session should be recovered we don't bring down the
	// whole server because of one naughty connection.
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Session panicked:", string(debug.Stack()))
		}
	}()
	defer s.conn.Close()

	for {
		s.lmt.N = LINELIMIT
		tag, cmd, args, err := ParseLine(s.buf)
		if err != nil {
			// TODO: be better here.
			return
		}

		switch cmd {
		case "CAPABILITY":
			s.Capability(tag, args)

		case "NOOP":
			s.Noop(tag, args)

		case "LOGOUT":
			s.Logout(tag, args)
			return

		case "STARTTLS":
			s.NotImplemented(tag, args)

		case "AUTHENTICATE":
			s.Authenticate(tag, args)

		case "LOGIN":
			s.Login(tag, args)

		case "SELECT":
			s.Select(tag, args)

		case "EXAMINE":
			s.Examine(tag, args)

		case "CREATE":
			s.Create(tag, args)

		case "DELETE":
			s.Delete(tag, args)

		case "RENAME":
			s.Rename(tag, args)

		case "SUBSCRIBE":
			s.Subscribe(tag, args)

		case "UNSUBSCRIBE":
			s.Unsubscribe(tag, args)

		case "LIST":
			s.List(tag, args)

		case "LSUB":
			s.LSub(tag, args)

		case "STATUS":
			s.Status(tag, args)

		case "APPEND":
			s.Append(tag, args)

		case "CHECK":
			s.Check(tag, args)

		case "CLOSE":
			s.Close(tag, args)

		case "EXPUNGE":
			s.Expunge(tag, args)

		case "SEARCH":
			s.Search(tag, args)

		case "FETCH":
			s.Fetch(tag, args)

		case "STORE":
			s.Store(tag, args)

		case "COPY":
			s.Copy(tag, args)

		case "UID":
			s.UID(tag, args)

		default:
			s.BadCommand(tag, args)
		}
	}
}

func (s *Session) Capability(tag string, args []string) {
	fmt.Fprint(s.conn, "* CAPABILITY IMAP4rev1 AUTH=PLAIN\r\n")
	fmt.Fprintf(s.conn, "%s OK CAPABILITY completed\r\n", tag)
}

func (s *Session) Noop(tag string, args []string) {
	// TODO: Return status updates here
	fmt.Fprintf(s.conn, "%s OK NOOP completed\r\n", tag)
}

func (s *Session) Logout(tag string, args []string) {
	fmt.Fprint(s.conn, "* BYE IMAP4rev1 talaria logging off\r\n")
	fmt.Fprintf(s.conn, "%s OK LOGOUT completed\r\n")
}

func (s *Session) Authenticate(tag string, args []string) {
	if s.state != NotAuthenticated {
		fmt.Fprintf("%s BAD AUTHENTICATE already logged in\r\n")
		return
	}

	// TODO: Actual auth!
}
