package fail2go

import (
	"strings"
)

func (conn *Fail2goConn) GlobalStatus() ([]string, error) {
	fail2banInput := []string{"status"}

	output, err := conn.fail2banRequest(fail2banInput)
	if err != nil {
		return nil, err
	}

	//TODO use reflection to assert data structures and give proper errors
	jails := output.([]interface{})[1].([]interface{})[1].([]interface{})[1]
	return strings.Split(jails.(string), ","), nil
}

func (conn *Fail2goConn) GlobalPing() (string, error) {
	fail2banInput := []string{"ping"}

	output, err := conn.fail2banRequest(fail2banInput)
	if err != nil {
		return "", err
	}

	//TODO use reflection to assert data structures and give proper errors
	return output.([]interface{})[1].(string), nil
}
