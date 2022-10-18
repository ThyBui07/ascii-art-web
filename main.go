package main

import (
	"ascii/server"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

var tmpl *template.Template

func init() {
	_, err := os.Stat("template/index.html")
	if err == nil {
		tmpl = template.Must(template.ParseFiles("template/index.html"))
		fmt.Println("hello from inti")
	} else {
		tmpl = nil
		fmt.Println("hello nil")
	}
	//read the template

}

func main() {
	// Http. handle(pattern, handler Http.handler) => http.Fileserver(root) return handler => dir : root directory
	// call variable from server package and add value
	server.Tpl = tmpl
	//handle css from static directory
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	//get request
	http.HandleFunc("/", server.GetRequest)
	//post request
	http.HandleFunc("/ascii-art", server.PostRequest)
	//get request download
	http.HandleFunc("/download", server.GetDownloadFile)
	//open port- listen
	log.Fatal(http.ListenAndServe(":8080", nil))
}
