package ssh

import (
	"context"
	"fmt"
	"log"
	"time"

	"golang.org/x/crypto/ssh"

	"github.com/agocan/oak/pkg/data"
)

type Ssh struct {
	Client *ssh.Client
	Ctx    context.Context
	Cancel context.CancelFunc
	Name   string
}

func NewSsh(machine *data.Machine) (s *Ssh) {

	config := &ssh.ClientConfig{
		User:            machine.User,
		Timeout:         5 * time.Second,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	if machine.Type == "password" {
		config.Auth = []ssh.AuthMethod{
			ssh.Password(machine.Password),
		}
	} else {

	}

	hostport := fmt.Sprintf("%s:%d", machine.Host, machine.Port)
	client, err := ssh.Dial("tcp", hostport, config)
	if err != nil {
		log.Fatal("connect ssh err: ", err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	s = &Ssh{Client: client, Ctx: ctx, Cancel: cancel, Name: machine.Name}
	return
}

func (s *Ssh) Close() {
	s.Client.Close()
}
