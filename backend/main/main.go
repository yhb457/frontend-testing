package main

import (
	"fmt"
	"khu-capstone-18-backend/database"
	"khu-capstone-18-backend/handler"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

func main() {
	if err := database.ConnectDB(); err != nil {
		fmt.Println("DB CONNECTION ERR:", err)
		return
	}

	if err := database.TestDB(); err != nil {
		fmt.Println("DB PING ERR:", err)
		return
	}

	r := mux.NewRouter()
	r.HandleFunc("/auth/signup", handler.SignUpHandler).Methods("POST")
	r.HandleFunc("/auth/login", handler.LoginHandler).Methods("POST")
	r.HandleFunc("/auth/logout", handler.LogoutHandler).Methods("POST")
	r.HandleFunc("/competitions", handler.CompetitionHandler).Methods("GET")
	r.HandleFunc("/competition", handler.PostCompetitionHandler).Methods("POST")

	// CORS 미들웨어 설정
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
	})

	handler := c.Handler(r)

	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
