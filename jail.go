package fail2go

func (conn *Fail2goConn) JailStatus(jail string) (map[string]interface{}, error) {
	fail2banInput := []string{"status", jail}

	fail2banOutput, err := conn.fail2banRequest(fail2banInput)
	if err != nil {
		return nil, err
	}

	//TODO use reflection to assert data structures and give proper errors
	action := fail2banOutput.([]interface{})[1].([]interface{})[1].([]interface{})[1]
	filter := fail2banOutput.([]interface{})[1].([]interface{})[0].([]interface{})[1]

	output := make(map[string]interface{})

	output["currentlyFailed"] = filter.([]interface{})[0].([]interface{})[1]
	output["totalFailed"] = filter.([]interface{})[1].([]interface{})[1]
	output["fileList"] = filter.([]interface{})[2].([]interface{})[1]

	output["currentlyBanned"] = action.([]interface{})[0].([]interface{})[1]
	output["totalBanned"] = action.([]interface{})[1].([]interface{})[1]
	output["IPList"] = action.([]interface{})[2].([]interface{})[1]

	return output, nil
}

func (conn *Fail2goConn) JailFailRegex(jail string) (map[string][]interface{}, error) {
	fail2banInput := []string{"get", jail, "failregex"}

	fail2banOutput, err := conn.fail2banRequest(fail2banInput)
	if err != nil {
		return nil, err
	}

	return map[string][]interface{}{
		"failregex": fail2banOutput.([]interface{})[1].([]interface{}),
	}, nil
}

func (conn *Fail2goConn) JailBanIP(jail string, ip string) (map[string]string, error) {
	fail2banInput := []string{"set", jail, "banip", ip}

	fail2banOutput, err := conn.fail2banRequest(fail2banInput)
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"bannedIP": fail2banOutput.([]interface{})[1].(string),
	}, nil
}

func (conn *Fail2goConn) JailUnbanIP(jail string, ip string) (map[string]string, error) {
	fail2banInput := []string{"set", jail, "unbanip", ip}

	fail2banOutput, err := conn.fail2banRequest(fail2banInput)
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"unbannedIP": fail2banOutput.([]interface{})[1].(string),
	}, nil
}
