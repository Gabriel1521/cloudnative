
package main

import (
	"fmt"
	"io"
	"net/http"
	"hw2/controller"
)


func main(){
	mux := http.NewServerMux()
	mux.HandleFunc("/getHeader",controller.GetHeader)
	logger.Print("Listening server address localhost:80")
	if err := http.ListenAndServe("localhost:80", nil); err != nil {
		logger.Fatal("start server failer")
	}
}
