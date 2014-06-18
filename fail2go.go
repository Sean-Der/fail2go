package fail2go

type Fail2goConn struct {
	Fail2banSocket string
}

func Newfail2goConn(fail2banSocket string) *Fail2goConn {
	return &Fail2goConn{fail2banSocket}
}
