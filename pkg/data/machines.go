package data

import (
	"fmt"
	"log"
)

type Machine struct {
	Name       string `yaml:"name"`
	Host       string `yaml:"host"`
	User       string `yaml:"user"`
	Password   string `yaml:"password"`
	Port       int    `yaml:"port"`
	PrivateKey string `yaml:"private_key"`
	PublicKey  string `yaml:"public_key"`
	Type       string `yaml:"type"`
}

func (m *Machine) Add() {
	var data Data
	data.Read()

	data.Machines = append(data.Machines, *m)
	data.Write()

}

func (m *Machine) List() {
	var data Data
	data.Read()

	for i, val := range data.Machines {
		fmt.Printf("Index: %d Name: %s\n", i, val.Name)
	}
}

func (m *Machine) Del() {
	var data Data
	data.Read()
	index := m.GetIndex()
	if len(data.Machines) == (index + 1) {
		data.Machines = data.Machines[:index]
	} else {
		data.Machines = append(data.Machines[:index], data.Machines[index+1:]...)
	}

	data.Write()

}

func (m *Machine) Update() {
	var data Data
	data.Read()
	index := m.GetIndex()
	data.Machines[index] = *m
	data.Write()

}

func (m *Machine) GetIndex() (index int) {
	var data Data
	data.Read()
	index = -1
	for i, val := range data.Machines {
		if val.Name == m.Name {
			index = i
			break
		}
	}
	if index == -1 {
		log.Fatal("Not found group:", m.Name)
	}
	return
}
