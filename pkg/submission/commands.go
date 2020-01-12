package submission

import (
	"github.com/alecthomas/participle"
	"github.com/alecthomas/participle/lexer"
	"github.com/alecthomas/participle/lexer/regex"
)

const grammar = `
	# Command keywords
	HELO = (?i)HELO
	EHLO = (?i)EHLO
	MAILFROM: = (?)MAIL FROM:

	SP = \x20
	CRLF = \r\n
	String = \S+
`

var (
	Lexer  lexer.Definition
	Parser *participle.Parser
)

func init() {
	var err error
	Lexer, err = regex.New(grammar)
	if err != nil {
		panic(err)
	}

	Parser, err = participle.Build(
		&Request{},
		participle.Lexer(Lexer),
	)
	if err != nil {
		panic(err)
	}
}

type Request struct {
	HELORequest *HELORequest `  @@`
	EHLORequest *EHLORequest `| @@`
}

type HELORequest struct {
	Domain string `HELO SP @String CRLF`
}

type EHLORequest struct {
	Domain string `EHLO SP @String CRLF`
}

type MAILRequest struct {
	ReversePath string    `MAILFROM @String`
	Parameters  []*string `[`
}
