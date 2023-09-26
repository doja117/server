package main

import (
	"log"
	"net/http"
)

type apiConfig struct {
	fileServerHits int
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return cfg.fileServerHits
}
func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type ", "text/plain;charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}

func main() {

	filePathRoot := "."
	port := "8080"

	mux := http.NewServeMux()
	mux.Handle("/app/", http.StripPrefix("/app", http.FileServer(http.Dir(filePathRoot))))
	mux.HandleFunc("/healthz", handlerReadiness)
	mux.HandleFunc("/reset", middlewareMetricsInc)
	corsMux := middlewareCors(mux)
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: corsMux,
	}
	log.Fatal(srv.ListenAndServe())
}
