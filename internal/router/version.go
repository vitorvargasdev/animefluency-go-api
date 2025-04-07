package router

import (
	"fmt"
	"net/http"
	"os"
)

func getVersion(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	version := os.Getenv("VERSION")

	if _, err := w.Write([]byte(version)); err != nil {
		fmt.Println(err)
	}
}
