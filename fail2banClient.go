package fail2go

import (
	"bytes"
	"errors"
	"net"

	"github.com/kisielk/og-rek"
)

const END_COMMAND = "<F2B_END_COMMAND>"

func (conn *Conn) fail2banRequest(input []string) (interface{}, error) {
	c, err := net.Dial("unix", conn.Fail2banSocket)

	if err != nil {
		return nil, errors.New("Failed to contact fail2ban socket")
	}

	p := &bytes.Buffer{}
	ogórek.NewEncoder(p).Encode(input)
	c.Write(p.Bytes())
	c.Write([]byte(END_COMMAND))

	buf := make([]byte, 0)
	tmpBuf := make([]byte, 1)
	for {
		bufRead, err := c.Read(tmpBuf)

		if err != nil {
			return nil, errors.New("Failed to contact fail2ban socket")
		} else if bufRead != 0 {
			buf = append(buf, tmpBuf...)
			if bytes.HasSuffix(buf, []byte(END_COMMAND)) {
				c.Close()
				break
			}
		} else {
			break
		}

	}
	buf = buf[:len(buf)-len(END_COMMAND)]

	dec := ogórek.NewDecoder(bytes.NewBuffer(buf))
	fail2banOutput, err := dec.Decode()

	if fail2banOutput != nil && err == nil {
		fail2banOutput = fail2banOutput.(ogórek.Tuple)[1]

		switch fail2banOutput.(type) {
		case ogórek.Call:
			Call := fail2banOutput.(ogórek.Call)
			return nil, errors.New(Call.Callable.Name + ": " + Call.Args[0].(string))
		}
	}

	return fail2banOutput, err
}
