package data

import (
	"fmt"
	"log"
	"strings"
)

func (g *Group) AddMachine() {
	res := g.IsNotIn()
	if len(res) != 0 {
		resStr := strings.Join(res, ",")
		log.Fatal(fmt.Sprintf(" %s is not in machines.", resStr))
	}

	g.UpdateMachine()
	fmt.Println("Add machine successed.")
}

func (g *Group) ListMachine() {
	group := GetGroup(g.Name)
	for i, val := range group.Machines {
		fmt.Printf("Index: %d Name: %s\n", i, val)
	}
}

func (g *Group) DelMachine(slice []string) {
	g = GetGroup(g.Name)
	res := g.IsNotIn()
	if len(res) != 0 {
		resStr := strings.Join(res, ",")
		log.Fatal(fmt.Sprintf(" %s is not in machines.", resStr))
	}
	for _, v := range slice {
		g.Machines = removeSliceElement(g.Machines, v)
	}
	g.UpdateMachine()
	fmt.Println("Remove machine successed.")
}

func (g *Group) UpdateMachine() {
	var data Data
	data.Read()
	g.Machines = removeDuplicateElement(g.Machines)

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
	data.Groups[index] = *g
	data.Write()
}
