package config

import (
	"fmt"
	"log"
	"os"

	homedir "github.com/mitchellh/go-homedir"
)

func GetOakDefaultDir() string {
	homeDirector, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)
	}

	oakDirector := fmt.Sprintf("%s/.oak/", homeDirector)
	if _, err := os.Stat(oakDirector); os.IsNotExist(err) {
		err = os.Mkdir(oakDirector, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}

	return oakDirector
}
