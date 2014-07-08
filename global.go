package fail2go

import (
	"strings"
)

func (conn *Conn) GlobalStatus() ([]string, error) {
	fail2BanOutput, err := conn.fail2banRequest([]string{"status"})
	if err != nil {
		return nil, err
	}

	jails := fail2BanOutput.([]interface{})[1].([]interface{})[1].([]interface{})[1]
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

	return output.([]interface{})[1].(string), nil
}
