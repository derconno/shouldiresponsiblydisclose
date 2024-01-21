package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

var (
	indexTemplate *template.Template

	data []struct {
		Name string
		Src  string
	}
)

func getRoot(w http.ResponseWriter, _ *http.Request) {

	err := indexTemplate.Execute(w, data)

	if err != nil {
		w.WriteHeader(500)
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			return
		}
	}
}

func main() {
	setup()

	http.HandleFunc("/", getRoot)
	http.Handle("/submit", http.RedirectHandler("https://github.com/derconno/shouldiresponsiblydisclose/issues", http.StatusFound))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	err := http.ListenAndServe("0.0.0.0:8080", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("[+] Server closed")
	} else if err != nil {
		fmt.Printf("[-] %s, \n", err.Error())
	}
}

func setup() {
	indexTemplate = template.Must(template.New("index.html").ParseFiles("templates/index.html"))

	rawData, err := os.ReadFile("data.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(rawData, &data)
	if err != nil {
		panic(err)
	}
}
