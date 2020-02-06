package main

import (
	// "fmt"
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

func dpkg_all(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("/root/data/dpkg/dpkg.json")
	if err != nil {
		log.Println("Error fetching file")
		return
	}
	var packages []Package
	err = json.Unmarshal(data, &packages)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(packages)
}

// func dpkg_one(w http.ResponseWriter, r *http.Request) {
// 	data, err := ioutil.ReadFile("/root/data/dpkg/dpkg.json")
// 	if err != nil {
// 		log.Println("Error fetching file")
// 		return
// 	}
// 	var packages []Package
// 	err = json.Unmarshal(data, &packages)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	title := r.URL.Path[len("/package/"):]
// 	for _, pack := range packages {
// 		if title == pack.Package {
// 			json.NewEncoder(w).Encode(pack)
// 			return
// 		}
// 	}
// }

func main() {
	path := "/root/go/go_services/reaktor-app/static/build"
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(path))))
	http.HandleFunc("/api/dpkg/", dpkg_all)
	// http.HandleFunc("/package/", dpkg_one)
	log.Println("Server running on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}