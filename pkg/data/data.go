package data

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"

	"github.com/agocan/oak/config"
)

type Data struct {
	Groups   []Group   `yaml:"groups"`
	Machines []Machine `yaml:"machines"`
}

func (d *Data) Read() {
	data, _ := ioutil.ReadFile(config.StorageParh)
	err := yaml.Unmarshal(data, &d)
	if err != nil {
		log.Fatal(err)
	}
}

func (d *Data) Write() {
	out, err := yaml.Marshal(&d)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(config.StorageParh, out, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

// IsOk 判断groups内的数据和machines数据是否能匹配上
func (d *Data) IsOk() bool {
	var machines map[string]string
	machines = make(map[string]string, 10)
	for _, mac := range d.Machines {
		machines[mac.Name] = ""
	}

	for _, group := range d.Groups {
		for _, macName := range group.Machines {
			if _, ok := machines[macName]; !ok {
				return false
			}
		}
	}
	return true
}

func (d *Data) GetGroup(index int) (group Group) {
	return d.Groups[index]
}

func (d *Data) GetMachine(index int) (machine Machine) {
	return d.Machines[index]
}
