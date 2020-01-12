package submission

import (
	"fmt"
	"strings"
	"testing"
)

func TestParseHELO(t *testing.T) {
	reqs := make(chan *Request, 10)
	r := strings.NewReader("HELO example.com\r\nEHLO example.com\r\n")
	err := Parser.Parse(r, reqs)
	if err != nil {
		t.Error(err)
	}

	for req := range reqs {
		fmt.Printf("%#v\n", req)
	}
}
