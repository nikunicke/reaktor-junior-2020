package main

import (
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type Package struct {
	Package			string
	Depends			[][]string
	Reverse_Depends	[]string
	Description		string
}

func get_dpkg_data(w http.ResponseWriter, r *http.Request) {
	var packages []Package

	data, err := ioutil.ReadFile("/root/data/dpkg/dpkg.json")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if err = json.Unmarshal(data, &packages); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(packages)
}

func main() {
	path := "/root/go/go_services/reaktor-app/static/build"
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(path))))
	http.HandleFunc("/api/dpkg/", get_dpkg_data)
	log.Println("Server running on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}