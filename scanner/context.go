package scanner

import (
	"fmt"
	"net"
	"sync"
	"time"
)

type Result struct {
	Port   int
	IsOpen bool
}

type ApplicationContext struct {
	Ip        string
	PortStart int
	PortEnd   int
	Report    []Result
}

const BUF = 3000

var pipe = make(chan Result, BUF)

func (a *ApplicationContext) Start() {
	fmt.Printf("Scaning started, traget ip: %s, port[%d:%d]\n\n", a.Ip, a.PortStart, a.PortEnd)
	var wg sync.WaitGroup

	wg.Add(a.PortEnd - a.PortStart + 1)
	for i := a.PortStart; i <= a.PortEnd; i++ {
		go scan(a.Ip, i)
	}

	go listen(a, &wg)
	wg.Wait()

	c := 1
	for _, v := range a.Report {
		if v.IsOpen {
			fmt.Printf("[%d] port %d open\n", c, v.Port)
			c++
		}
	}
}

func scan(ip string, port int) {
	addr := fmt.Sprintf("%s:%d", ip, port)

	conn, err := net.DialTimeout("tcp", addr, 2*time.Second)
	r := Result{
		Port:   port,
		IsOpen: err == nil,
	}
	pipe <- r
	if err == nil {
		defer conn.Close()
	}
}

func listen(a *ApplicationContext, wg *sync.WaitGroup) {
	for i := a.PortStart; i <= a.PortEnd; {
		r := <-pipe
		a.Report = append(a.Report, r)
		i++
		wg.Done()
	}
}
