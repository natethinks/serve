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

	var ip net.IP
	host, _ := os.Hostname()
	addrs, _ := net.LookupIP(host)
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			ip = ipv4
		}
	}

	fmt.Printf("Serving current directory %v on http://%v:8080", dir, ip)
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir(dir))))
}
