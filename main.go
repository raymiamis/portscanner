package main

import (
	"flag"
	"fmt"
	"net"
	"strconv"
	"time"
)

func scanPort(protocol, hostname string, port int, timeout time.Duration) bool {
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, timeout)
	if err != nil {
		return false
	}
	_ = conn.Close()
	return true
}

func main() {
	host := flag.String("host", "", "target host")
	startPort := flag.Int("start", 1, "start port (1–65535), standard 1")
	endPort := flag.Int("end", 1024, "end port (1–65535), standard 1024")
	timeout := flag.Int("timeout", 200, "timeout in ms for each port, standard 200ms")

	flag.Parse()

	if *host == "" {
		fmt.Println("error: target host is necessary. example: portscanner -host scanme.nmap.org")
		flag.Usage()
		return
	}
	if *startPort < 1 || *endPort > 65535 || *startPort > *endPort {
		fmt.Println("error: invalid start/end port combination. example: -start 20 -end 80")
		flag.Usage()
		return
	}

	fmt.Printf("initializing scan of %s from port %d to %d (timeout %dms)...\n", *host, *startPort, *endPort, *timeout)
	estTime := (*endPort + 1 - *startPort) * *timeout / 1000
	fmt.Printf("estimated time: %ds\n", estTime)

	for port := *startPort; port <= *endPort; port++ {
		if scanPort("tcp", *host, port, time.Duration(*timeout)*time.Millisecond) {
			fmt.Printf("--> port %d is open\n", port)
		}
	}

	fmt.Printf("all set! scan completed.")
}
