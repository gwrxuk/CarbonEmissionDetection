package routers

import (
	"carbon/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()

	// Item routes
	router.HandleFunc("/items", handlers.GetItems).Methods("GET")
	router.HandleFunc("/items/{id}", handlers.GetItem).Methods("GET")
	router.HandleFunc("/items", handlers.CreateItem).Methods("POST")
	router.HandleFunc("/items/{id}", handlers.UpdateItem).Methods("PUT")
	router.HandleFunc("/items/{id}", handlers.DeleteItem).Methods("DELETE")
	router.HandleFunc("/prediction", handlers.Prediction).Methods("POST")
	router.HandleFunc("/", serveIndex)
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	return router
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/index.html")
}

