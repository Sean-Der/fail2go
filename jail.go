package fail2go

import (
	"errors"
	"github.com/kisielk/og-rek"
	"strconv"
)

func (conn *Conn) JailStatus(jail string) (int64, int64, []string, int64, int64, []string, error) {
	fail2banOutput, err := conn.fail2banRequest([]string{"status", jail})
	if err != nil {
		return 0, 0, nil, 0, 0, nil, err
	}

	action := fail2banOutput.([]interface{})[1].([]interface{})[1].([]interface{})[1]
	filter := fail2banOutput.([]interface{})[1].([]interface{})[0].([]interface{})[1]

	return filter.([]interface{})[0].([]interface{})[1].(int64),
		filter.([]interface{})[1].([]interface{})[1].(int64),
		interfaceSliceToStringSlice(filter.([]interface{})[2].([]interface{})[1].([]interface{})),
		action.([]interface{})[0].([]interface{})[1].(int64),
		action.([]interface{})[1].([]interface{})[1].(int64),
		interfaceSliceToStringSlice(action.([]interface{})[2].([]interface{})[1].([]interface{})),
		nil
}

func (conn *Conn) JailFailRegex(jail string) ([]string, error) {
	fail2banOutput, err := conn.fail2banRequest([]string{"get", jail, "failregex"})
	if err != nil {
		return nil, err
	}
	return interfaceSliceToStringSlice(fail2banOutput.([]interface{})[1].([]interface{})), nil
}

func (conn *Conn) JailAddFailRegex(jail string, regex string) ([]string, error) {
	fail2banOutput, err := conn.fail2banRequest([]string{"set", jail, "addfailregex", regex})
	if err != nil {
		return nil, err
	}
	fail2banOutput = fail2banOutput.([]interface{})[1]

	switch fail2banOutput.(type) {
	case ogórek.Call:
		return nil, errors.New(fail2banOutput.(ogórek.Call).Args[0].(string))
	case []interface{}:
		return interfaceSliceToStringSlice(fail2banOutput.([]interface{})), nil
	}
	return nil, errors.New("Unexpected output from fail2ban")
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
	return fail2banOutput.([]interface{})[1].(string), nil
}

func (conn *Conn) JailUnbanIP(jail string, ip string) (string, error) {
	fail2banOutput, err := conn.fail2banRequest([]string{"set", jail, "unbanip", ip})
	if err != nil {
		return "", err
	}
	return fail2banOutput.([]interface{})[1].(string), nil
}

func (conn *Conn) JailFindTime(jail string) (int64, error) {
	fail2banOutput, err := conn.fail2banRequest([]string{"get", jail, "findtime"})
	if err != nil {
		return -1, err
	}
	return fail2banOutput.([]interface{})[1].(int64), nil
}

func (conn *Conn) JailSetFindTime(jail string, findTime int) (int64, error) {
	fail2banOutput, err := conn.fail2banRequest([]string{"set", jail, "findtime", strconv.Itoa(findTime)})
	if err != nil {
		return -1, err
	}
	return fail2banOutput.([]interface{})[1].(int64), nil
}

func (conn *Conn) JailMaxRetry(jail string) (int64, error) {
	fail2banOutput, err := conn.fail2banRequest([]string{"get", jail, "maxretry"})
	if err != nil {
		return -1, err
	}
	return fail2banOutput.([]interface{})[1].(int64), nil
}

func (conn *Conn) JailSetMaxRetry(jail string, maxRetry int) (int64, error) {
	fail2banOutput, err := conn.fail2banRequest([]string{"set", jail, "maxretry", strconv.Itoa(maxRetry)})
	if err != nil {
		return -1, err
	}
	return fail2banOutput.([]interface{})[1].(int64), nil
}

func (conn *Conn) JailUseDNS(jail string) (string, error) {
	fail2banOutput, err := conn.fail2banRequest([]string{"get", jail, "usedns"})
	if err != nil {
		return "", err
	}
	return fail2banOutput.([]interface{})[1].(string), nil
}

func (conn *Conn) JailSetUseDNS(jail string, useDNS string) (string, error) {
	fail2banOutput, err := conn.fail2banRequest([]string{"set", jail, "usedns", useDNS})
	if err != nil {
		return "", err
	}
	return fail2banOutput.([]interface{})[1].(string), nil
}
