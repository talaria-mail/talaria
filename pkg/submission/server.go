package submission

import (
	"bufio"
	"io"
	"net"
)

type Server struct {
	Addr string
}

func (s *Server) Serve(l net.Listener) error {
	for {
		conn, err := l.Accept()
		if err != nil {
			return err
		}

		session := s.NewSession(conn)
		go session.Run()
	}
}

func (s *Server) ListenAndServe() error {
	l, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return err
	}

	return s.Serve(l)
}

func (s *Server) NewSession(conn io.ReadWriteCloser) Session {
	hostname, _, err := net.SplitHostPort(s.Addr)
	if err != nil {
		// Shouldn't have gotten this far with an invalid address. Lets just
		// freak out
		panic(err)
	}

	session := Session{
		conn:     conn,
		lmt:      io.LimitedReader{R: conn, N: LINELIMIT},
		hostname: hostname,
	}

	session.buf = bufio.NewReader(&session.lmt)
	return session
}
