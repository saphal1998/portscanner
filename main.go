package main

import (
	"flag"
	"fmt"
	"math"
	"net"
	"sync"
	"time"
)

type Configuration struct {
	concurrentPortCount uint
	timeout             time.Duration
}

func scanPort(host string, port uint16, timeout time.Duration, openPortChan chan uint16) {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, port), timeout)
	if err != nil {
		openPortChan <- 0
		return
	}
	defer conn.Close()
	openPortChan <- port
}

func openPortProcessor(openPortChannel chan uint16) {
	for port := range openPortChannel {
		if port != 0 {
			fmt.Printf("%d port is open\n", port)
		}
	}
}

func main() {
	host := flag.String("host", "127.0.0.1", "hostname to scan")
	timeout := flag.Uint("timeout", 100, "timeout for connection establishment")
	concurrentPortCount := flag.Uint("concurrent port count", 1000, "number of ports to look at concurrently")
	flag.Parse()

	conf := Configuration{
		timeout:             time.Duration(*timeout) * time.Millisecond,
		concurrentPortCount: *concurrentPortCount,
	}

	fmt.Printf("Scanning host %s\n", *host)

	openPortChannel := make(chan uint16, conf.concurrentPortCount)
	limiter := make(chan int, conf.concurrentPortCount)

	go openPortProcessor(openPortChannel)

	var wg sync.WaitGroup

	// First 1024 ports are reserved anyway
	for i := 1025; i < math.MaxUint16; i++ {
		limiter <- 1
		wg.Add(1)
		go func(host string, port uint16, timeout time.Duration, openPortChannel chan uint16) {
			scanPort(host, port, conf.timeout, openPortChannel)
			<-limiter
			wg.Done()
		}(*host, uint16(i), conf.timeout, openPortChannel)
	}

	wg.Wait()
}
