package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
)

var (
	indexTemplate *template.Template
)

func getRoot(w http.ResponseWriter, _ *http.Request) {

	err := indexTemplate.Execute(w, nil)

	if err != nil {
		w.WriteHeader(500)
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			return
		}
	}
}

func main() {

	indexTemplate = template.Must(template.New("index.html").ParseFiles("templates/index.html"))

	http.HandleFunc("/", getRoot)

	err := http.ListenAndServe("0.0.0.0:8080", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("[+] Server closed")
	} else if err != nil {
		fmt.Printf("[-] %s, \n", err.Error())
	}
}
