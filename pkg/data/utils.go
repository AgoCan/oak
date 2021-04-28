package data

func GetMachineNameMap() (res map[string]int) {
	var data Data
	data.Read()
	res = make(map[string]int, len(data.Machines))
	for _, v := range data.Machines {
		res[v.Name] = 0
	}
	return res
}

func removeDuplicateElement(slice []string) []string {
	result := make([]string, 0, len(slice))
	temp := map[string]struct{}{}
	for _, item := range slice {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

func GetGroup(name string) (group *Group) {
	var data Data
	data.Read()
	for _, val := range data.Groups {
		if val.Name == name {
			return &val
		}
	}
	return
}

func removeSliceElement(slice []string, element string) []string {

	for i, val := range slice {
		if val == element {
			slice = append(slice[:i], slice[i+1:]...)
			break
		}
	}
	return slice
}

func GetMachine(name string) (machine *Machine) {
	var data Data
	data.Read()
	for _, val := range data.Machines {
		if val.Name == name {
			return &val
		}
	}
	return
}
