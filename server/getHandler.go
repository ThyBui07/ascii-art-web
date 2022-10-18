package server

import (
	"html/template"
	"net/http"
)

//create a variable from server package
var Tpl *template.Template

func GetRequest(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorNotFound(w)
		return
	}
	if Tpl == nil {
		errorInternalServer(w)
		return
	}
	Tpl.Execute(w, struct{ Submit bool }{false})
}

func GetDownloadFile(w http.ResponseWriter, r *http.Request) {
	b := []byte(FinalText)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(b)
	return
}
