package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/gorilla/mux"
)

// Get all dnsdock records and return as JSON
func GetDNSRecords(w http.ResponseWriter, req *http.Request) {
	res, err := http.Get("http://dnsdock.devtools.vm/services")
	if err != nil {
		panic(err.Error())
	}

	var records map[string]interface{}
	json.NewDecoder(res.Body).Decode(&records)

	json.NewEncoder(w).Encode(records)
}

// Get all running containers and return as JSON
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

// Listen to container lifecycle events so we can notify the dashboard client to refresh
func SetupDockerEventListener() {
	client, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	// Listen to only these events
	filters := filters.NewArgs()
	filters.Add("event", "start")
	filters.Add("event", "die")
	filters.Add("event", "pause")
	filters.Add("event", "unpause")

	log.Print("Registering Docker Event Listener....")
	messages, errs := client.Events(context.Background(), types.EventsOptions{Filters: filters})

loop:
	for {
		select {
		case err := <-errs:
			if err != nil && err != io.EOF {
				log.Fatal(err)
			}
			break loop
		case e := <-messages:
			// Ping the dashboard here via websocket
			log.Printf("Docker Event: %s", e.Action)
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
