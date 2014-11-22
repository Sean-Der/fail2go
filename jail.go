package fail2go

import (
	"errors"
	"strconv"
)

func (conn *Conn) JailStatus(jail string) (currentlyFailed int64, totalFailed int64, fileList []string, currentlyBanned int64, totalBanned int64, IPList []string, err error) {
	fail2banOutput, err := conn.fail2banRequest([]string{"status", jail})
	if err != nil {
		return
	}

	action := fail2banOutput.([]interface{})[1].([]interface{})[1]
	filter := fail2banOutput.([]interface{})[0].([]interface{})[1]

	currentlyFailed = filter.([]interface{})[0].([]interface{})[1].(int64)
	totalFailed = filter.([]interface{})[1].([]interface{})[1].(int64)
	fileList = interfaceSliceToStringSlice(filter.([]interface{})[2].([]interface{})[1].([]interface{}))
	currentlyBanned = action.([]interface{})[0].([]interface{})[1].(int64)
	totalBanned = action.([]interface{})[1].([]interface{})[1].(int64)
	IPList = interfaceSliceToStringSlice(action.([]interface{})[2].([]interface{})[1].([]interface{}))
	return
}

func (conn *Conn) JailFailRegex(jail string) ([]string, error) {
	fail2banOutput, err := conn.fail2banRequest([]string{"get", jail, "failregex"})
	if err != nil {
		return nil, err
	}
	return interfaceSliceToStringSlice(fail2banOutput.([]interface{})), nil
}

func (conn *Conn) JailAddFailRegex(jail string, regex string) ([]string, error) {
	fail2banOutput, err := conn.fail2banRequest([]string{"set", jail, "addfailregex", regex})
	if err != nil {
		return nil, err
	}
	return interfaceSliceToStringSlice(fail2banOutput.([]interface{})), nil
}

func (conn *Conn) JailDeleteFailRegex(jail string, regex string) (interface{}, error) {
	currentRegexes, _ := conn.JailFailRegex(jail)
	regexIndex := stringInSliceIndex(regex, currentRegexes)
	if regexIndex == -1 {
		return nil, errors.New("Regex is not in jail")
	}

	fail2banOutput, err := conn.fail2banRequest([]string{"set", jail, "delfailregex", strconv.Itoa(regexIndex)})
	if err != nil {
		return nil, err
	}
	return fail2banOutput, nil
}

func (conn *Conn) JailBanIP(jail string, ip string) (string, error) {
	fail2banOutput, err := conn.fail2banRequest([]string{"set", jail, "banip", ip})
	if err != nil {
		return "", err
	}
	return fail2banOutput.(string), nil
}

func (conn *Conn) JailUnbanIP(jail string, ip string) (string, error) {
	fail2banOutput, err := conn.fail2banRequest([]string{"set", jail, "unbanip", ip})
	if err != nil {
		return "", err
	}
	return fail2banOutput.(string), nil
}

func (conn *Conn) JailFindTime(jail string) (int64, error) {
	fail2banOutput, err := conn.fail2banRequest([]string{"get", jail, "findtime"})
	if err != nil {
		return -1, err
	}
	return fail2banOutput.(int64), nil
}

func (conn *Conn) JailSetFindTime(jail string, findTime int) (int64, error) {
	fail2banOutput, err := conn.fail2banRequest([]string{"set", jail, "findtime", strconv.Itoa(findTime)})
	if err != nil {
		return -1, err
	}
	return fail2banOutput.(int64), nil
}

func (conn *Conn) JailMaxRetry(jail string) (int64, error) {
	fail2banOutput, err := conn.fail2banRequest([]string{"get", jail, "maxretry"})
	if err != nil {
		return -1, err
	}
	return fail2banOutput.(int64), nil
}

func (conn *Conn) JailSetMaxRetry(jail string, maxRetry int) (int64, error) {
	fail2banOutput, err := conn.fail2banRequest([]string{"set", jail, "maxretry", strconv.Itoa(maxRetry)})
	if err != nil {
		return -1, err
	}
	return fail2banOutput.(int64), nil
}

func (conn *Conn) JailUseDNS(jail string) (string, error) {
	fail2banOutput, err := conn.fail2banRequest([]string{"get", jail, "usedns"})
	if err != nil {
		return "", err
	}
	return fail2banOutput.(string), nil
}

func (conn *Conn) JailSetUseDNS(jail string, useDNS string) (string, error) {
	fail2banOutput, err := conn.fail2banRequest([]string{"set", jail, "usedns", useDNS})
	if err != nil {
		return "", err
	}
	return fail2banOutput.(string), nil
}

func (conn *Conn) JailActions(jail string) ([]string, error) {
	fail2banOutput, err := conn.fail2banRequest([]string{"get", jail, "actions"})
	if err != nil {
		return nil, err
	}
	return interfaceSliceToStringSlice(fail2banOutput.([]interface{})), nil
}
