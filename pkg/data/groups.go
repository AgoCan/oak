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

func (g *Group) Del() {
	var data Data
	data.Read()
	var index int
	for i, val := range data.Machines {
		if val.Name == g.Name {
			index = i
			break
		} else {
			log.Fatal("Not found machine:", g.Name)
		}
	}
	if len(data.Groups) == (index + 1) {
		data.Groups = data.Groups[:index]
	} else {
		data.Groups = append(data.Groups[:index], data.Groups[index+1:]...)
	}
	data.Write()
}

func (g *Group) Update() {
	var data Data
	data.Read()
	var index int
	for i, val := range data.Machines {
		if val.Name == g.Name {
			index = i
			break
		} else {
			log.Fatal("Not found machine:", g.Name)
		}
	}
	data.Groups[index] = *g
	data.Write()
}

func (g *Group) AddMachine() {
	var data Data
	data.Read()
	g.Update()
	data.Write()
}
func (g *Group) ListMachine() {
	var data Data
	data.Read()
}
func (g *Group) DelMachine() {
	var data Data
	data.Read()
}
func (g *Group) UpdateMachine() {
	var data Data
	data.Read()
}
