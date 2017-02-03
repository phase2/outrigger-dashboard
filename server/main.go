package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/fsouza/go-dockerclient"
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
	if client, err := docker.NewClientFromEnv(); err == nil {

		if containers, err := client.ListContainers(docker.ListContainersOptions{All: false}); err == nil {
			json.NewEncoder(w).Encode(containers)
		} else {
			panic(err)
		}
	} else {
		panic(err)
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
	log.Fatal(http.ListenAndServe(":80", router))
}
