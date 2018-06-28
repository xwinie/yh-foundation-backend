package cores

import (
	"flag"
)

func InitFlag() *string {

	configFile := flag.String("conf", "conf/app.conf", "set configuration `file`")
	flag.Parse()
	return configFile
}
