package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"sync/atomic"

	"github.com/gorilla/mux"
)

var (
	configuration *Configuration
	count         int32 = 0
)

type Configuration struct {
	Port         string
	InstanceName string
}

func main() {
	loadConfiguration()

	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(struct {
			Count        int32  `json:"count"`
			InstanceName string `json:"instanceName"`
		}{
			Count:        atomic.LoadInt32(&count),
			InstanceName: configuration.InstanceName,
		}); err != nil {
			log.Println(err)
		}

		atomic.AddInt32(&count, 1)
	})

	log.Fatal(http.ListenAndServe(":"+configuration.Port, router))
}

func loadConfiguration() {
	const defaultPort = "80"

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = defaultPort
	}

	configuration = &Configuration{
		Port:         port,
		InstanceName: os.Getenv("APP_INSTANCE_NAME"),
	}
}
