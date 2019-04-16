package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// App manages the server
type App struct {

	// storage.go::Storage manages the actual internal storage
	*Storage

	// ws-hub.go:: Hub manages the websocket connections
	*Hub

	// router to manage the routes
	router *mux.Router
}

// Initializes the app with the required data
func initApp() *App {

	h := newHub()

	// Run the goroutine managing the websocket connections
	go h.run()

	a := &App{
		Hub:     h,
		Storage: &Storage{hashmap: make(map[string]string)},
	}

	// Create the routes for the app
	a.createRoutes()

	return a
}

// Create routes for the server
func (a *App) createRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/api", a.postHandler).Methods("POST")
	r.HandleFunc("/api/{key}", a.getHandler).Methods("GET")
	r.HandleFunc("/api/watch", a.wsHandler)
	a.router = r
}

func (a *App) postHandler(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	// Fit the request body inside a pair, i.e, a (string, string) struct
	var entry Pair
	if err := decoder.Decode(&entry); err != nil {
		// if the format is incorrect
		sendError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Insert the data into the storage
	a.Storage.putValue(&entry)

	// Broadcast message to all listening websockets
	a.Hub.broadcast <- &entry

	// Send response back the response
	sendResponse(w, http.StatusCreated, &map[string]string{"Key": entry.Key, "Value": entry.Value})
}

func (a *App) getHandler(w http.ResponseWriter, r *http.Request) {
	// gets key form url
	key := mux.Vars(r)["key"]

	// get value from storage
	value, ok := a.Storage.getValue(key)

	// if value not present
	if !ok {
		sendError(w, http.StatusNotFound, "Key not found")
		return
	}

	// send back the response
	sendResponse(w, http.StatusOK, &map[string]string{"Key": key, "Value": value})

}

// Handle incoming websocket connection requests
func (a *App) wsHandler(w http.ResponseWriter, r *http.Request) {
	serveWs(a.Hub, w, r)
}
