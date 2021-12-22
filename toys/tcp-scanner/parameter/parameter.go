package parameter

import "flag"

var (
	Host = flag.String("host", "127.0.0.1", "IP address of a host")
	StartPort = flag.Int("start", 1, "the start of a range of ports")
	EndPort = flag.Int("end", 100, "the end of a range of ports")
)

func init() {
	flag.Parse()
}