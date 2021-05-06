package data

import (
	"fmt"
	"log"
)

type Group struct {
	Name     string   `yaml:"name"`
	Machines []string `yaml:"machine"`
}

func (g *Group) Add() {
	var data Data
	data.Read()
	IsDuplicateName(g.Name)
	data.Groups = append(data.Groups, *g)
	data.Write()
}

func (g *Group) List() {
	var data Data
	data.Read()
	for i, val := range data.Groups {
		fmt.Printf("Index: %d Name: %s\n", i, val.Name)
	}
}

func (g *Group) Del() {
	var data Data
	data.Read()
	var index int
	index = -1
	for i, val := range data.Groups {
		if val.Name == g.Name {
			index = i
			break
		}
	}
	if index == -1 {
		log.Fatal("Not found group:", g.Name)
	}
	if len(data.Groups) == (index + 1) {
		data.Groups = data.Groups[:index]
	} else {
		data.Groups = append(data.Groups[:index], data.Groups[index+1:]...)
	}
	data.Write()
}

func (g *Group) Update(newName string) {
	var data Data
	data.Read()

	var index int
	index = -1
	for i, val := range data.Groups {
		if val.Name == g.Name {
			index = i
			break
		}
	}
	if index == -1 {
		log.Fatal("Not found group:", g.Name)
	}
	g.Name = newName
	IsDuplicateName(g.Name)
	data.Groups[index] = *g
	data.Write()
	fmt.Println("Change name successed.")
}

func (g *Group) IsNotIn() (notIn []string) {
	var data Data
	data.Read()
	if len(g.Machines) == 0 {
		return notIn
	}
	res := GetMachineNameMap()

	for _, machine := range g.Machines {
		if _, ok := res[machine]; !ok {
			notIn = append(notIn, machine)
		}
	}
	return notIn
}
