package main


import (
	"fmt"
	"net/http"
	"log"
	"github.com/gorilla/mux"
)

// Router - Returns the gorilla mux router
func Router() *mux.Router {
	router := mux.NewRouter()
	// router.HandleFunc("/api/my-endpoint", Handler).Methods("GET","OPTIONS")
	return router
}

func main() {
	r := Router()

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("client/build/static"))))

	r.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "client/build/index.html")
	})

	fmt.Println("App running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}