package main
import (
  "net/http"
  "os"
)
func ping(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("pong"))
}
func main() {
  port := os.Getenv("PORT")
  http.Handle("/", http.FileServer(http.Dir("./http")))
  http.HandleFunc("/ping", ping)
  if err := http.ListenAndServe(":"+port, nil); err != nil {
    panic(err)
  }
}
