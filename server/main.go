package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/gorilla/mux"
)

func GetDNSRecords(w http.ResponseWriter, req *http.Request) {
	res, err := http.Get("http://dnsdock.devtools.vm/services")
	if err != nil {
		panic(err.Error())
	}

	//var records []DNSRecord
	var records map[string]interface{}
	json.NewDecoder(res.Body).Decode(&records)

	json.NewEncoder(w).Encode(records)
}

func GetContainers(w http.ResponseWriter, req *http.Request) {
	client, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	if containers, err := client.ContainerList(context.Background(), types.ContainerListOptions{All: false}); err == nil {
		json.NewEncoder(w).Encode(containers)
	} else {
		panic(err)
	}
}

func SetupDockerEventListener() {
	client, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	log.Print("Registering Docker Event Listener....")
	messages, errs := client.Events(context.Background(), types.EventsOptions{})

	loop:
		for {
			select {
				case err := <-errs:
					if err != nil && err != io.EOF {
						log.Fatal(err)
					}
					break loop
				case e := <-messages:
					log.Print(e)
 			}
		}
}


func Redirect(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, "/static/", 301)
}

func main() {
	router := mux.NewRouter()

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("/app/"))))

	router.HandleFunc("/", Redirect).Methods("GET")
	router.HandleFunc("/api/dnsrecords", GetDNSRecords).Methods("GET")
	router.HandleFunc("/api/containers", GetContainers).Methods("GET")

	// Start a goroutine for the http server event loop so we can
	// still register for docker events on the main goroutine
	go func() {
		port := ":80"
		log.Printf("Starting HTTP server listening on port %s", port)
		log.Fatal(http.ListenAndServe(port, router))
	}()

	SetupDockerEventListener()
}
