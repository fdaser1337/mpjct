package main
import (
	"fmt"
	"net/http"
)

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}
func login(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "login endpoint")
}
func main() {
	http.HandleFunc("/health", health)
	http.HandleFunc("/login", login)

	fmt.Println("Auth service running on :8080")
	http.ListenAndServe(":8080", nil)
}