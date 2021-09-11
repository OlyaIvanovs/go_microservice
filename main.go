package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
)

var opt struct {
	Port int `default:"9090"`
}

func main() {
	err := envconfig.Process("Microservice", &opt)
	if err != nil {
		log.Printf("Failed to parse command line arguments: %s", err.Error())
	}

	r := mux.NewRouter()

	routes := []struct {
		route   string
		handler func(http.ResponseWriter, *http.Request)
		method  string
	}{
		{
			route:   "/",
			method:  "GET",
			handler: handleGetHello,
		}, {
			route:   "/goodbye",
			method:  "GET",
			handler: handleGetGoodbye,
		}}

	for _, p := range routes {
		r.HandleFunc(p.route, p.handler).Methods(p.method)
	}

	port := strconv.Itoa(opt.Port)
	log.Println(port)

	http.ListenAndServe(":"+port, r)
}

func handleGetHello(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello world!")
	d, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Oops", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Hello %s", d)
}

func handleGetGoodbye(w http.ResponseWriter, r *http.Request) {
	log.Println("Goodbye world!")
}
