package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

func helloName(w http.ResponseWriter, r *http.Request) {
	log.Printf("Serving request: %s", r.URL.Path)
	host, _ := os.Hostname()
	fmt.Fprintf(w, "Hola Mundo!\n")
	fmt.Fprintf(w, "Version: 1.0.0\n")
	fmt.Fprintf(w, "Hostname: %s\n", host)


	address, _ := net.InterfaceAddrs()

	for _, a := range address {
		inet, ok := a.(*net.IPNet)

		if ok && !inet.IP.IsLoopback() {
			if inet.IP.To4() != nil {
				fmt.Fprintf(w, "IP Address: %s\n", inet.IP.String())
			}
		}
	}
}



func main()  {
	port := "8080"

	server := http.NewServeMux()
	server.HandleFunc("/", helloName)
	log.Printf("Server listening on port %s", port)

	err := http.ListenAndServe(":"+port, server)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}


