package main

import "net/http"
import "os"
import "fmt"

func main () {
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":80", nil)
}

func Hello (w http.ResponseWriter, r * http.Request) {
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")
	fmt.Fprintf(w,"<h1>Hello Full Cycle! I am %s, I am %s</h1>", name, age)
}