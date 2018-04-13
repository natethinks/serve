package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	ip := GetOutboundIP()

	fmt.Printf("Starting Server\nCurrent directory: %v\nLocal address: http://localhost:8080\nExternal address: http://%v:8080", dir, ip)
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir(dir))))
}

// Get the ip address of the current machine
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	local := conn.LocalAddr().(*net.UDPAddr)

	return local.IP
}
