package main

import (
	"log"
	"net"
	"strconv"
	"sync"
	"time"
)

var ip_to_scan = "192.168.1.1"
var min_port = 1
var max_port = 1024
var wg sync.WaitGroup

func main() {

	for port := min_port; port <= max_port; port++ {
		wg.Add(1)
		go testTCPConnection(ip_to_scan, port) // go threads
	}

	wg.Wait()
	
}

func testTCPConnection(ip string, port int) {
	defer wg.Done()
	_, err := net.DialTimeout("tcp", ip+":"+strconv.Itoa(port), time.Second*10)
	if err == nil {
		log.Printf("Host %s has open port: %d\n", ip, port)
	}
	
}