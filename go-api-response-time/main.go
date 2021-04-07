package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Client struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
    Email string `json:"email"`
    Phone string `json:"phone"`
    Domain string `json:"domain"`
    Description string `json:"description"`
    Tags string `json:"tags"`
    CreatedAt string `json:"created_at"`
    UpdatedAt string `json:"updated_at"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:secret@tcp(127.0.0.1:3306)/api_response_time")
	if err != nil {
		panic(err.Error())
	}

	query, err := db.Query("SELECT * FROM clients")
	if err != nil {
		panic(err.Error())
	}
	defer query.Close()

	var clients []*Client
	for query.Next() {
		var client Client
		err := query.Scan(&client.ID, &client.Name, &client.Email, &client.Phone, &client.Domain, &client.Description, &client.Tags, &client.CreatedAt, &client.UpdatedAt)
		if err != nil {
			panic(err.Error())
		}
		clients = append(clients, &client)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(clients)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/v1/clients/", indexHandler)
	http.ListenAndServe(":"+port, mux)
}
