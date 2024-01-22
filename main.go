package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

func main() {
	host := flag.String("host", "127.0.0.1", "hostname to scan")
	port := flag.Int("port", 8080, "port to check")

	flag.Parse()

	fmt.Printf("Host %s, Port: %d\n", *host, *port)

	_, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *host, *port))
	if err != nil {
		fmt.Println("Port is closed")
		os.Exit(1)
	}

	fmt.Printf("Port %d is open", *port)
	os.Exit(0)

}
