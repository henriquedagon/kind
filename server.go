package main

import (
	"net/http"
	"os"
	"fmt"
	"io/ioutil"
	"log"
)

func main () {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/configmap", ConfigMap)
	http.ListenAndServe(":80", nil)
}

func Hello (w http.ResponseWriter, r * http.Request) {
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")
	fmt.Fprintf(w,"<h1>Hello Full Cycle! I am %s, I am %s</h1>", name, age)
}

func ConfigMap (w http.ResponseWriter, r * http.Request) {
	data, err := ioutil.ReadFile("/go/myfamily/myfamily.txt")
	if err != nil {
		log.Fatalf("Error reading file: ", err)
	}
	fmt.Fprintf(w,"<h1>My family: %s</h1>", string(data))
}