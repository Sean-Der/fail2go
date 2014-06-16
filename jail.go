package fail2go

import ()

func JailStatus(jail string) (map[string]interface{}, error) {
	fail2banInput := make([]string, 2)
	fail2banInput[0] = "status"
	fail2banInput[1] = jail

	fail2banOutput, err := fail2banRequest(fail2banInput)
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
	output["ipList"] = action.([]interface{})[2].([]interface{})[1]

	return output, nil
}

func JailFailRegex(jail string) (map[string][]interface{}, error) {
	fail2banInput := make([]string, 3)
	fail2banInput[0] = "get"
	fail2banInput[1] = jail
	fail2banInput[2] = "failregex"

	fail2banOutput, err := fail2banRequest(fail2banInput)
	if err != nil {
		return nil, err
	}

	return map[string][]interface{}{
		"failregex": fail2banOutput.([]interface{})[1].([]interface{}),
	}, nil
}
