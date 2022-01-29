package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

const listenAddr = ":8000"

const lineBreak = "==============================================="

func main() {
	http.HandleFunc("/", handler)
	fmt.Printf("Starting listening %s\n", listenAddr)
	if err := http.ListenAndServe(listenAddr, nil); err != nil {
		panic(err)
	}
}

func handler(writer http.ResponseWriter, request *http.Request) {
	requestDump, err := httputil.DumpRequest(request, true)
	if err != nil {
		return
	}
	fmt.Println(lineBreak)
	fmt.Println(string(requestDump))
	fmt.Println(lineBreak)
}
