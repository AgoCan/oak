package ssh

import (
	"bytes"
	"fmt"
	"io"
	"log"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/pkg/sftp"
)

func (s *Ssh) Copy() {
	defer s.Close()
	client, err := sftp.NewClient(s.Client)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	home, err := homedir.Dir()
	w := client.Walk(home)
	for w.Step() {
		if w.Err() != nil {
			continue
		}
		log.Println(w.Path())
	}

	f, err := client.Open("hello.txt")
	buf := new(bytes.Buffer)
	buf.ReadFrom(io.LimitReader(f, 100000))
	msg := buf.String()
	fmt.Println(msg)
	// f, err := client.Create("hello.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// if _, err := f.Write([]byte("Hello world!")); err != nil {
	// 	log.Fatal(err)
	// }
	// f.Close()

	// check it's there
	fi, err := client.Lstat("hello.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(fi)
}
