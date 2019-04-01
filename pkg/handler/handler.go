package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/fredeastside/luasgo/pkg/luas"
	"github.com/pkg/errors"
)

// Handler - serve HTTP requests
type Handler struct {
	http.Handler
}

// NewHandler - constrcutor for Handler struct
func NewHandler() *Handler {
	return &Handler{}
}

// ErrorResponse - error response presenter
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// NewErrorResponse - constrcutor for ErrorResponse struct
func NewErrorResponse(code int, message string) *ErrorResponse {
	return &ErrorResponse{code, message}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	urlArr := strings.Split(r.URL.Path, "/")
	switch urlArr[1] {
	case "stops":
		if len(urlArr) == 2 {
			stopsHandler(w, r)

			return
		}

		stopHandler(urlArr[2], w, r)
	case "fares":
		faresHandler(w, r)
	default:
		http.NotFound(w, r)
	}
}

func stopsHandler(w http.ResponseWriter, r *http.Request) {
	stops, err := luas.GetStops()
	if err != nil {
		writeBadResponseJSON(w)

		return
	}

	writeJSON(w, http.StatusOK, stops)
}

func stopHandler(stop string, w http.ResponseWriter, r *http.Request) {
	data, err := luas.GetStop(stop)
	if err != nil {
		writeBadResponseJSON(w)

		return
	}
	writeJSON(w, http.StatusOK, data)
}

func faresHandler(w http.ResponseWriter, r *http.Request) {
	from := r.URL.Query().Get("from")
	if from == "" {
		writeBadResponseJSON(w)

		return
	}

	to := r.URL.Query().Get("to")
	if to == "" {
		writeBadResponseJSON(w)

		return
	}

	fares, err := luas.GetFares(from, to, isChildren(r.URL.Query().Get("children")))
	if err != nil {
		writeBadResponseJSON(w)

		return
	}
	writeJSON(w, http.StatusOK, fares)
}

func isChildren(param string) bool {
	if param == "1" || param == "true" {
		return true
	}

	return false
}

func writeBadResponseJSON(w http.ResponseWriter) {
	writeJSON(w, http.StatusBadRequest, NewErrorResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest)))
}

func writeJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	json, err := json.Marshal(data)
	if err != nil {
		log.Fatal(errors.Wrap(err, "can not marshal data to json."))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(json)
}
