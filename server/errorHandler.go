package server

import "net/http"

func errorInternalServer(w http.ResponseWriter) {
	http.Error(w, "Error 500 Internal Server", 500)
}

func errorBadRequest(w http.ResponseWriter, err error) {
	http.Error(w, "Error 400 Bad Request\n"+err.Error(), 400)
}

func errorNotFound(w http.ResponseWriter) {
	http.Error(w, " Error 404 Not found", 404)
}
