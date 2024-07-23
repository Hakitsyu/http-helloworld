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
	count int32 = 0
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(struct {
			Count        int32  `json:"count"`
			InstanceName string `json:"instanceName"`
		}{
			Count:        atomic.LoadInt32(&count),
			InstanceName: os.Getenv("INSTANCE_NAME"),
		}); err != nil {
			log.Println(err)
		}

		atomic.AddInt32(&count, 1)
	})

	log.Fatal(http.ListenAndServe(":8080", router))
}
