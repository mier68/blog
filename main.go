package main

import (
	"log"
	"net/http"
	"time"

	"github.com/goprogramming/blog/internal/routers"
)

func main() {
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		log.Fatalf("connect error:%v", err)
	}
}
