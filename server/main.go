package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type DNSRecord struct {
	Name    string		`json:"id,omitempty"`
	Image   string		`json:"name,omitempty"`
	IPs     []string	`json:"ips,omitempty"`
	TTL     int				`json:"ttl,omitempty"`
	Aliases []string	`json:"aliases,omitempty"`
}

func GetDNSRecords(w http.ResponseWriter, req *http.Request) {
	res, err := http.Get("http://dnsdock.devtools.vm/services")
	if err != nil {
		panic(err.Error())
	}

	//var records []DNSRecord
	var records map[string]interface{}
	json.NewDecoder(res.Body).Decode(&records)

	log.Printf("Records returned: %s", records)

	//params := mux.Vars(req)
	//for _, item := range people {
	//	if item.ID == params["id"] {
	//		json.NewEncoder(w).Encode(item)
	//		return
	//	}
	//}
	//json.NewEncoder(w).Encode(&Person{})
}


func main() {
	router := mux.NewRouter()
	router.HandleFunc("/dnsrecords", GetDNSRecords).Methods("GET")
	log.Fatal(http.ListenAndServe(":9090", router))
}
