package ssh

import (
	"log"
	"os"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

func (s *Ssh) Terminal() (err error) {
	defer s.Close()
	session, err := s.Client.NewSession()
	if err != nil {
		log.Printf("cannot open new session: %v", err)
		return err
	}
	defer session.Close()

	go func() {
		<-s.Ctx.Done()
		s.Client.Close()
	}()

	fd := int(os.Stdin.Fd())
	state, err := terminal.MakeRaw(fd)
	if err != nil {
		log.Printf("terminal make raw: %s", err)
		return err
	}
	defer terminal.Restore(fd, state)

	w, h, err := terminal.GetSize(fd)
	if err != nil {
		log.Printf("terminal get size: %s", err)
		return err
	}

	modes := ssh.TerminalModes{
		ssh.ECHO:          1,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}

	term := os.Getenv("TERM")
	if term == "" {
		term = "xterm-256color"
	}
	if err := session.RequestPty(term, h, w, modes); err != nil {
		log.Printf("session xterm: %s", err)
		return err
	}

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin

	if err = session.Shell(); err != nil {
		log.Printf("session shell: %s", err)
		return err
	}

	if err = session.Wait(); err != nil {
		if e, ok := err.(*ssh.ExitError); ok {
			switch e.ExitStatus() {
			case 130:
				return err
			}
		}
		log.Printf("ssh: %s", err)
	}
	return err
}
