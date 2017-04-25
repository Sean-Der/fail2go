package fail2go

import (
	"database/sql"
	"encoding/json"
	"strings"

	"github.com/kisielk/og-rek"
	_ "github.com/mattn/go-sqlite3"
)

func (conn *Conn) GlobalStatus() ([]string, error) {
	fail2BanOutput, err := conn.fail2banRequest([]string{"status"})
	if err != nil {
		return nil, err
	}

	jails := fail2BanOutput.([]interface{})[1].(ogórek.Tuple)[1]
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

type Ban struct {
	Jail, IP  string
	TimeOfBan int
	Data      BanData
}

type BanData struct {
	Matches  []string
	Failures int
}

func (conn *Conn) GlobalBans() (results []Ban, err error) {
	DBFile, err := conn.GlobalDBFile()
	if err != nil {
		return nil, err
	}

	dbConn, err := sql.Open("sqlite3", DBFile)
	if err != nil {
		return nil, err
	}
	rows, err := dbConn.Query("select jail, ip, timeofban, data from bans")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var ban Ban
		var data string
		rows.Scan(&ban.Jail, &ban.IP, &ban.TimeOfBan, &data)
		json.Unmarshal([]byte(data), &ban.Data)
		results = append(results, ban)
	}
	rows.Close()

	return results, nil
}
