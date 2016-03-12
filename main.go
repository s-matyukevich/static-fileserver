package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	var server, port, dir string
	flag.StringVar(&server, "s", "0.0.0.0", "Server")
	flag.StringVar(&port, "p", "80", "Port")
	flag.StringVar(&dir, "d", "./public", "Directory")
	flag.Parse()
	fmt.Printf("Starting file server...\n")
	panic(http.ListenAndServe(server+":"+port, http.FileServer(http.Dir(dir))))
}
