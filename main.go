package oak

import (
	"flag"

	"github.com/agocan/oak/config"
)

var configFile = flag.String("c", "../config/config.yaml", "config fime path.")

func main() {

	config.Init(*configFile)
}
