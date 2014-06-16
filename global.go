package fail2go

import (
	"strings"
)

func GlobalStatus() ([]string, error) {
	fail2banInput := make([]string, 1)
	fail2banInput[0] = "status"

	output, err := fail2banRequest(fail2banInput)
	if err != nil {
		return nil, err
	}

	//TODO use reflection to assert data structures and give proper errors
	jails := output.([]interface{})[1].([]interface{})[1].([]interface{})[1]
	return strings.Split(jails.(string), ","), nil
}

func GlobalPing() (string, error) {
	fail2banInput := make([]string, 1)
	fail2banInput[0] = "ping"

	output, err := fail2banRequest(fail2banInput)
	if err != nil {
		return "", err
	}

	//TODO use reflection to assert data structures and give proper errors
	return output.([]interface{})[1].(string), nil
}
