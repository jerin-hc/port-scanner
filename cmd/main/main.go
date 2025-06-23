package main

import (
	"flag"
	"github.com/jerin-hc/port-scanner/scanner"
	"strconv"
	"strings"
)

var ip *string = flag.String("ip", "", "traget IP address")
var portRange *string = flag.String("range", "0:1024", "Port range")

func main() {
	flag.Parse()
	if *ip == "" {
		panic("Target IP required")
	}
	ports := strings.Split(*portRange, ":")
	start, err := strconv.Atoi(ports[0])
	if err != nil {
		panic("invalid starting port")
	}
	end, err := strconv.Atoi(ports[1])
	if err != nil {
		panic("invalid ending port")
	}

	a := scanner.ApplicationContext{
		Ip:        *ip,
		PortStart: start,
		PortEnd:   end,
	}
	a.Start()
}
