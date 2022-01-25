package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"github.com/gorilla/mux"
)

func SetMemoryLimit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	memory := vars["memory"] // memory
	log.Println("memory variable" + memory)

	cmd, err := exec.Command("/bin/sh", "limit.sh").Output()
	if err != nil {
		fmt.Printf("error %s", err)
	}
	output := string(cmd)
	log.Println("output" + output)
}

func SetUploadLimit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	upload := vars["upload"] // memory
	log.Println("upload variable" + upload)

	cmd, err := exec.Command("/bin/sh", "upload.sh").Output()
	if err != nil {
		fmt.Printf("error %s", err)
	}
	output := string(cmd)
	log.Println("output" + output)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	//specify status code
	w.WriteHeader(http.StatusOK)

	//update response writer
	fmt.Fprintf(w, "API is up and running")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/health-check", HealthCheck).Methods("GET")
	router.HandleFunc("/memory-limit/{memory}", SetMemoryLimit).Methods("POST")
	router.HandleFunc("/upload-limit/{upload}", SetMemoryLimit).Methods("POST")
	http.Handle("/", router)
	http.ListenAndServe(":8080", router)
}
