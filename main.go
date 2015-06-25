package main

import (
	//"fmt"
	"log"
	"net/http"
	"os"
)

var logger = log.New(os.Stdout, "logger: ", log.Lshortfile)

func main(){
	logger.Println("server is starting at ....8989")
	http.HandleFunc("/", index)
	err := http.ListenAndServe(":8989", nil)
	if err != nil {
		logger.Println(err)
	}
}

func index(w http.ResponseWriter, req *http.Request){
	//fmt.Fprint(w, "Hello World")
	http.ServeFile(w, req, "./index.html")
}
