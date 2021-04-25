package ssh

type Exec struct {
	Message chan string
}

func (e *Exec) Excute(s *Ssh, command string) (err error) {

	return
}
