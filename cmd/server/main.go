package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	internalhttp "github.com/elena9bulat-wq/go-rest-userprofile/internal/http"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := internalhttp.NewRouter()

	server := &http.Server{
		Addr:              ":" + port,
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
	}

	fmt.Printf("Server running on http://localhost:%s\n", port)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}
