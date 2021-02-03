package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, p.store.GetPlayerScore(player))
}

func main() {
	server := PlayerServer{}
	handler := http.HandlerFunc(server.ServeHTTP)
	if err := http.ListenAndServe(":4000", handler); err != nil {
		log.Fatalf("could not listen on port 4000 %v", err)
	}
}

// func PlayerServer(w http.ResponseWriter, r *http.Request) {
// 	player := strings.TrimPrefix(r.URL.Path, "/players/")
// 	score := GetPlayerScore(player)
// 	fmt.Fprint(w, score)
// }

func GetPlayerScore(player string) string {
	if player == "Pepper" {
		return "20"
	}

	if player == "Floyd" {
		return "10"
	}

	return ""
}
