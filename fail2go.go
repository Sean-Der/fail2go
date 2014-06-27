package fail2go

type Conn struct {
	Fail2banSocket string
}

func Newfail2goConn(fail2banSocket string) *Conn {
	return &Conn{fail2banSocket}
}
