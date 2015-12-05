package main

import "net/http"

func serve(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "www/"+r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", serve)
	http.ListenAndServe(":8001", nil)
}
