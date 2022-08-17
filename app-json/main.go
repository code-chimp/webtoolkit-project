package main

import (
	"github.com/code-chimp/webtoolkit"
	"log"
	"net/http"
)

type RequestPayload struct {
	Action  string `json:"action"`
	Message string `json:"message"`
}
type ResponsePayload struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode,omitempty"`
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("."))))
	mux.HandleFunc("/receive-post", receivePost)
	mux.HandleFunc("/remote-service", remoteService)
	mux.HandleFunc("/simulated-service", simulatedService)

	log.Println("Starting service...")

	err := http.ListenAndServe(":8081", mux)
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func receivePost(w http.ResponseWriter, r *http.Request) {
	var payload RequestPayload
	var t webtoolkit.Tools

	err := t.ReadJSON(w, r, &payload)
	if err != nil {
		_ = t.ErrorJSON(w, err)
		return
	}

	response := ResponsePayload{
		Message: "hit the handler okay",
	}

	err = t.WriteJSON(w, http.StatusAccepted, response)
	if err != nil {
		_ = t.ErrorJSON(w, err)
		return
	}
}

func remoteService(w http.ResponseWriter, r *http.Request) {
	var payload RequestPayload
	var t webtoolkit.Tools

	err := t.ReadJSON(w, r, &payload)
	if err != nil {
		_ = t.ErrorJSON(w, err)
		return
	}

	_, statusCode, err := t.PushJSONToRemote("http://localhost:8081/simulated-service", payload)
	if err != nil {
		_ = t.ErrorJSON(w, err)
		return
	}

	response := ResponsePayload{
		Message:    "hit the handler okay",
		StatusCode: statusCode,
	}

	err = t.WriteJSON(w, http.StatusAccepted, response)
	if err != nil {
		_ = t.ErrorJSON(w, err)
		return
	}
}

func simulatedService(w http.ResponseWriter, r *http.Request) {
	payload := ResponsePayload{
		Message: "ok",
	}

	var t webtoolkit.Tools

	_ = t.WriteJSON(w, http.StatusOK, payload)
}
