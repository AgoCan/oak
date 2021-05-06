package ssh

import (
	"fmt"
	"log"
	"sync"

	"github.com/agocan/oak/pkg/data"
	"github.com/agocan/oak/pkg/utils"
)

var wg sync.WaitGroup

func (s *Ssh) Excute(command string) (err error) {

	session, err := s.Client.NewSession()
	logFileName := s.Name + ".log"
	if err != nil {
		utils.WriteDirectlyLog(logFileName, err.Error())
		log.Fatal("创建ssh session 失败", err)
	}
	defer session.Close()
	combo, err := session.CombinedOutput(command)
	if err != nil {
		utils.WriteDirectlyLog(logFileName, err.Error())
		utils.WriteDirectlyLog(logFileName, string(combo))
		log.Fatal("远程执行cmd 失败", err)
	}
	utils.WriteDirectlyLog(logFileName, string(combo))

	log.Printf("%s: \"%s\" is succeed.", s.Name, command)
	return
}

func GroupExec(group *data.Group, command string) {
	for _, k := range group.Machines {

		machine := data.GetMachine(k)
		if machine == nil {
			logMsg := fmt.Sprintf("%s machine not found.\n", group.Name)
			log.Fatal(logMsg)
		}
		s := NewSsh(machine)
		wg.Add(1)
		go func() {
			defer wg.Done()
			s.Excute(command)
		}()

	}
	wg.Wait()
}
