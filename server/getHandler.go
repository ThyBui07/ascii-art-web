package server

import (
	"html/template"
	"net/http"
	"strconv"
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
	leng := strconv.Itoa(len(b))
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachement")
	w.Header().Set("Content-Length", leng)

	w.Write(b)
	return
}
