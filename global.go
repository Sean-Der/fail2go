package fail2go

import (
	"strings"
)

func (conn *Conn) GlobalStatus() ([]string, error) {
	fail2BanOutput, err := conn.fail2banRequest([]string{"status"})
	if err != nil {
		return nil, err
	}

	jails := fail2BanOutput.([]interface{})[1].([]interface{})[1]
	output := make([]string, 0)
	for _, v := range strings.Split(jails.(string), ",") {
		output = append(output, strings.TrimSpace(v))
	}

	return output, nil
}

func (conn *Conn) GlobalPing() (string, error) {
	output, err := conn.fail2banRequest([]string{"ping"})
	if err != nil {
		return "", err
	}

	return output.(string), nil
}

func (conn *Conn) GlobalDBFile() (string, error) {
	output, err := conn.fail2banRequest([]string{"get", "dbfile"})
	if err != nil {
		return "", err
	}

	return output.(string), nil
}

func (conn *Conn) GlobalSetDBFile(dbfile string) (string, error) {
	output, err := conn.fail2banRequest([]string{"set", "dbfile", dbfile})
	if err != nil {
		return "", err
	}

	return output.(string), nil
}
