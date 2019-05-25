package main
import (
  "net/http"
  "os"
)
func ping(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("pong"))
}
func main() {
  port := getEnv("PORT", "80")
  http.Handle("/", http.FileServer(http.Dir("./http")))
  http.HandleFunc("/ping", ping)
  if err := http.ListenAndServe(":"+port, nil); err != nil {
    panic(err)
  }
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
