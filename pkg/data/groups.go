package operator

import (
	"fmt"
)

type Group struct {
	Name     string   `yaml:"name"`
	Machines []string `yaml:"machine"`
}

func (g *Group) Add() {
	var data Data
	data.Read()
	data.Groups = append(data.Groups, *g)
	data.Write()
}

func (g *Group) List() {
	var data Data
	data.Read()
	for i, val := range data.Groups {
		fmt.Printf("Index: %d Name: %s", i, val.Name)
	}
}

func (g *Group) Del(index int) {
	var data Data
	data.Read()
	if len(data.Groups) == (index + 1) {
		data.Groups = data.Groups[:index]
	} else {
		data.Groups = append(data.Groups[:index], data.Groups[index+1:]...)
	}
	data.Write()
}

func (g *Group) Update(index int) {
	var data Data
	data.Read()
	data.Groups[index] = *g
	data.Write()
}

func (g *Group) AddMachine(index int, Name string) {
	var data Data
	data.Read()
	//data.Groups = append(data.Groups, *g)
	data.Write()
}
func (g *Group) ListMachine(index int, Name string) {
	var data Data
	data.Read()
}
func (g *Group) DelMachine(index int, Name string) {
	var data Data
	data.Read()
}
func (g *Group) UpdateMachine(index int, Name string) {
	var data Data
	data.Read()
}
