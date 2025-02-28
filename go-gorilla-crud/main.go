package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type DB struct {
	Users []User
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	counter := 0
	db := DB{}

	r := mux.NewRouter()

	r.HandleFunc("/user", createUser(counter, &db)).Methods("POST")
	r.HandleFunc("/user/{id}", getUser(&db)).Methods("GET")
	r.HandleFunc("/user/{id}", nil).Methods("PUT")
	r.HandleFunc("/user/{id}", nil).Methods("DELETE")

	log.Println("Server started on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func createUser(counter int, db *DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var user User

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			fmt.Println(err)
		}

		user.ID = counter
		counter += 1

		db.Users = append(db.Users, user)
		log.Printf("User created successfully: %v\n", user)

		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "User created successfully")
	}
}

func getUser(db *DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		strId := vars["id"]

		id, err := strconv.Atoi(strId)
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}

		for i := range db.Users {
			if db.Users[i].ID == id {
				json.NewEncoder(w).Encode(db.Users[i])
				return
			}
		}

		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "User not found")
	}
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
