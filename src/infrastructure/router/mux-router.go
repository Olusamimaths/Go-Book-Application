package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type muxRouter struct{}

var (
	muxRouterInstance = mux.NewRouter()
)

func NewMuxRouter() Router {
	return &muxRouter{}
}

func (m *muxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxRouterInstance.HandleFunc(uri, f).Methods("GET")
}

func (m *muxRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxRouterInstance.HandleFunc(uri, f).Methods("POST")
}

func (m *muxRouter) SERVE(port string) {
	fmt.Printf("Mux HTTP server is running on port %v", port)
	http.ListenAndServe(port, muxRouterInstance)
}