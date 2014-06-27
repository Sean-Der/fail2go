package fail2go

import (
	"strings"
)

func (conn *Conn) GlobalStatus() ([]string, error) {
	fail2banInput := []string{"status"}

	fail2BanOutput, err := conn.fail2banRequest(fail2banInput)
	if err != nil {
		return nil, err
	}

	//TODO use reflection to assert data structures and give proper errors
	jails := fail2BanOutput.([]interface{})[1].([]interface{})[1].([]interface{})[1]
	output := make([]string, 0)
	for _, v := range strings.Split(jails.(string), ",") {
		output = append(output, strings.TrimSpace(v))
	}

	return output, nil
}

func (conn *Conn) GlobalPing() (string, error) {
	fail2banInput := []string{"ping"}

	output, err := conn.fail2banRequest(fail2banInput)
	if err != nil {
		return "", err
	}

	//TODO use reflection to assert data structures and give proper errors
	return output.([]interface{})[1].(string), nil
}
