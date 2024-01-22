package main

import (
	"flag"
	"fmt"
	"math"
	"net"
)

func scanPort(host string, port int) bool {
	_, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		return false
	}
	return true
}

func main() {
	host := flag.String("host", "127.0.0.1", "hostname to scan")
	flag.Parse()

	fmt.Printf("Scanning host %s\n", *host)

	for i := 0; i < math.MaxUint16; i++ {
		result := scanPort(*host, i)

		if result {
			fmt.Printf("Port %d is open\n", i)
		}
	}

}
