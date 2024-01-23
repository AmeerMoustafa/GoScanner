package main

import (
	"log"
	"net"
	"strconv"
	"time"
)

var ip_to_scan = "192.168.1.1"
var min_port = 1
var max_port = 1024

func main() {
	active_threads := 0
	done_channel := make(chan bool)

	for port := min_port; port <= max_port; port++ {
		go testTCPConnection(ip_to_scan, port, done_channel) // go threads
		active_threads++
	}

	for active_threads > 0 {
		<-done_channel
		active_threads--
	}
}

func testTCPConnection(ip string, port int, done_channel chan bool) {
	_, err := net.DialTimeout("tcp", ip+":"+strconv.Itoa(port), time.Second*10)
	if err == nil {
		log.Printf("Host %s has open port: %d\n", ip, port)
	}
	done_channel <- true
}