package fail2go

import (
	"bytes"
	"errors"
	"github.com/kisielk/og-rek"
	"net"
)

func (conn *Conn) fail2banRequest(input []string) (interface{}, error) {
	c, err := net.Dial("unix", conn.Fail2banSocket)

	if err != nil {
		return nil, errors.New("Failed to contact fail2ban socket")
	}

	p := &bytes.Buffer{}
	ogórek.NewEncoder(p).Encode(input)
	c.Write(p.Bytes())
	c.Write([]byte("<F2B_END_COMMAND>"))

	buf := make([]byte, 0)
	tmpBuf := make([]byte, 1)
	for {
		bufRead, _ := c.Read(tmpBuf)

		if bufRead != 0 {
			buf = append(buf, tmpBuf...)
		} else {
			buf = buf[:len(buf)-17]
			break
		}

	}

	dec := ogórek.NewDecoder(bytes.NewBuffer(buf))
	fail2banOutput, err := dec.Decode()

	if fail2banOutput != nil && err == nil {
		fail2banOutput = fail2banOutput.([]interface{})[1]

		switch fail2banOutput.(type) {
		case ogórek.Call:
			Call := fail2banOutput.(ogórek.Call)
			return nil, errors.New(Call.Callable.Name + ": " + Call.Args[0].(string))
		}
	}

	return fail2banOutput, err
}
