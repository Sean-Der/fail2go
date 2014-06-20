package fail2go

func (conn *Fail2goConn) JailStatus(jail string) (int64, int64, []interface{}, int64, int64, []interface{}, error) {
	fail2banInput := []string{"status", jail}

	fail2banOutput, err := conn.fail2banRequest(fail2banInput)
	if err != nil {
		return 0, 0, nil, 0, 0, nil, err
	}

	//TODO use reflection to assert data structures and give proper errors
	action := fail2banOutput.([]interface{})[1].([]interface{})[1].([]interface{})[1]
	filter := fail2banOutput.([]interface{})[1].([]interface{})[0].([]interface{})[1]

	return filter.([]interface{})[0].([]interface{})[1].(int64),
		filter.([]interface{})[1].([]interface{})[1].(int64),
		filter.([]interface{})[2].([]interface{})[1].([]interface{}),
		action.([]interface{})[0].([]interface{})[1].(int64),
		action.([]interface{})[1].([]interface{})[1].(int64),
		action.([]interface{})[2].([]interface{})[1].([]interface{}),
		nil
}

func (conn *Fail2goConn) JailFailRegex(jail string) ([]interface{}, error) {
	fail2banInput := []string{"get", jail, "failregex"}

	fail2banOutput, err := conn.fail2banRequest(fail2banInput)
	if err != nil {
		return nil, err
	}
	return fail2banOutput.([]interface{})[1].([]interface{}), nil
}

func (conn *Fail2goConn) JailBanIP(jail string, ip string) (string, error) {
	fail2banInput := []string{"set", jail, "banip", ip}

	fail2banOutput, err := conn.fail2banRequest(fail2banInput)
	if err != nil {
		return "", err
	}
	return fail2banOutput.([]interface{})[1].(string), nil
}

func (conn *Fail2goConn) JailUnbanIP(jail string, ip string) (string, error) {
	fail2banInput := []string{"set", jail, "unbanip", ip}

	fail2banOutput, err := conn.fail2banRequest(fail2banInput)
	if err != nil {
		return "", err
	}
	return fail2banOutput.([]interface{})[1].(string), nil
}
