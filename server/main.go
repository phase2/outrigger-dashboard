package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Clients connected to the websocket
var clients = make(map[*websocket.Conn]bool)

// Channel for broadcast of Container changes
var broadcast = make(chan []types.Container)

// Get all dnsdock records and return as JSON
func GetDNSRecords(w http.ResponseWriter, req *http.Request) {
	res, err := http.Get("http://dnsdock.outrigger.vm/services")
	if err != nil {
		fmt.Println("Error getting DNS records")
		panic(err.Error())
	}

	var records map[string]interface{}
	json.NewDecoder(res.Body).Decode(&records)
	json.NewEncoder(w).Encode(records)
}

// Create a Docker client from the current environment
func GetDockerClient() (*client.Client) {
	client, err := client.NewEnvClient()
	if err != nil {
		fmt.Println("Error creating new Docker client")
		panic(err)
	}
	return client
}

// Query Docker for all *running* containers
func GetContainers() []types.Container {
	client := GetDockerClient()

	if containers, err := client.ContainerList(context.Background(), types.ContainerListOptions{All: false}); err == nil {
		return containers
	} else {
		fmt.Println("Error finding all running containers")
		panic(err)
	}
}

// Get all running containers and encode response as JSON
func GetContainersJson(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(GetContainers())
}

// Get the Container object for the given id
func GetContainer(id string) types.ContainerJSON {
	client := GetDockerClient()
	container, err := client.ContainerInspect(context.Background(), id)
	if err != nil {
		fmt.Printf("Error inspecting container: %s", id)
		panic(err)
	}
	return container
}

// Get JSON for deep inspections of a container
func GetContainerJson(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	json.NewEncoder(w).Encode(GetContainer(id))
}

// WebSocket to relay container events to client
func ContainerWebSocket(w http.ResponseWriter, req *http.Request) {
	ws, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		fmt.Printf("Error upgrading websocket: %s", err.Error())
		return
	}

	defer ws.Close()

	// Register the client
	clients[ws] = true

	for {
		containers := <-broadcast

		for client := range clients {
			err := client.WriteJSON(containers)
			if err != nil {
				fmt.Printf("Error writing JSON: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}

}

// Listen to container lifecycle events so we can notify the dashboard client to refresh
func SetupDockerEventListener() {
	client, err := client.NewEnvClient()
	if err != nil {
		fmt.Println("Error creating Docker client")
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
				fmt.Println("Error returned from Docker event listener")
				log.Fatal(err)
			}
			break loop
		case e := <-messages:
			log.Printf("Received event '%s' from '%s'", e.Status, e.From);
			// Broadcast the event to listeners
			broadcast <- GetContainers()
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
	router.HandleFunc("/api/container/{id:[a-z0-9]+}", GetContainerJson).Methods("GET")
	router.HandleFunc("/api/containers", GetContainersJson).Methods("GET")
	router.HandleFunc("/api/containers/ws", ContainerWebSocket).Methods("GET")

	// Start a goroutine for the http server event loop so we can
	// still register for docker events on the main goroutine
	go func() {
		port := ":80"
		log.Printf("Starting HTTP server listening on port %s", port)
		log.Fatal(http.ListenAndServe(port, router))
	}()

	SetupDockerEventListener()
}
