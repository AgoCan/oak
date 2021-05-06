package utils

import (
	"io"
	"os"

	"github.com/agocan/oak/config"
)

func WriteDirectlyLog(fileName, msg string) error {
	if !IsExist(config.ExecLogPath) {
		return CreateDir(config.ExecLogPath)
	}
	var (
		err error
		f   *os.File
	)

	f, err = os.OpenFile(config.ExecLogPath+fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	_, err = io.WriteString(f, "\n"+msg)

	defer f.Close()
	return err
}

func CreateDir(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	os.Chmod(path, os.ModePerm)
	return nil
}

func IsExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}
