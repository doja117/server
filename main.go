package main

import (
	"net/http"
)

func main() {
	m := http.NewServeMux()
	imgDir := http.FileServer(http.Dir("./assets/logo.png"))
	m.Handle("./assets/logo.png", imgDir)
	m.Handle("/", http.FileServer(http.Dir(".")))
	m.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("OK"))
	})
	corsMux := middleWareCors(m)
	s := http.Server{}
	s.Handler = corsMux
	s.Addr = "localhost:8080"
	s.ListenAndServe()
}
func middleWareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
