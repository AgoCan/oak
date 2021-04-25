package operator

import (
	"flag"
	"fmt"
	"testing"

	"github.com/agocan/oak/config"
)

var configFile = flag.String("c", "../config/config.yaml", "config fime path.")

func TestData(t *testing.T) {
	config.Init(*configFile)

	var data Data
	data.Read()
	fmt.Println(data)
	fmt.Println(data.IsOk())
	// 	data.Write()
}
