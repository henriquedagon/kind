package main

import (
	"net/http"
	"os"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

var startedAt = time.Now()

func main () {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/secret", Secret)
	http.HandleFunc("/configmap", ConfigMap)
	http.HandleFunc("/healthz", Healthz)
	http.ListenAndServe(":80", nil)
}

func Hello (w http.ResponseWriter, r * http.Request) {
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")
	fmt.Fprintf(w,"<h1>Hello Full Cycle! I am %s, I am %s</h1>", name, age)
}

func Secret (w http.ResponseWriter, r * http.Request) {
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	fmt.Fprintf(w,"username: %s\n password: %s", user, password)
}

func ConfigMap (w http.ResponseWriter, r * http.Request) {
	data, err := ioutil.ReadFile("/go/myfamily/myfamily.txt")
	if err != nil {
		log.Fatalf("Error reading file: ", err)
	}
	fmt.Fprintf(w,"<h1>My family: %s</h1>", string(data))
}

func Healthz (w http.ResponseWriter, r * http.Request) {
	duration := time.Since(startedAt)

	if duration.Seconds() < 10 {
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("Duration: %v", duration.Seconds())))
	} else {
		w.WriteHeader(200)
		w.Write([]byte("ok"))
	}
}

