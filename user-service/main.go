package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {

	r := chi.NewRouter()

	r.Get("/", AllUsers)
	r.Post("/create", CreateUser)
	r.Delete("/delete/{id}", Deleteuser)
	r.Put("/upadte/{id}", Updateuser)
	r.Get("/login", Login)
	r.Post("/login", Login)
	r.Get("/register", Register)
	r.Post("/register", Register)

	log.Fatal(http.ListenAndServe(":8000", nil))

}

type User struct {
	Id       string
	Name     string
	Username string
	Password string
	Roles    []Role
}

type Role struct {
	Id   string
	Name string
}

func AllUsers(w http.ResponseWriter, r *http.Request) {

	users := []User{{Id: "324", Name: "fff", Username: "dddd", Password: "ff", Roles: []Role{}}}

	json.NewEncoder(w).Encode(users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {

	user := User{Id: "324", Name: "fff", Username: "dddd", Password: "ff", Roles: []Role{}}
	json.NewDecoder(r.Body).Decode(&user)

	json.NewEncoder(w).Encode(user)

}

func Deleteuser(w http.ResponseWriter, r *http.Request) {
	// var user User
	var id = chi.URLParam(r, "id")
	fmt.Println(id)
	json.NewEncoder(w).Encode("user deleeted successfully")
}

func Updateuser(w http.ResponseWriter, r *http.Request) {
	var id = chi.URLParam(r, "id")
	fmt.Println(id)
	json.NewEncoder(w).Encode("user updated successfully")
}

func Login(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		w.Write([]byte("login post called"))
	} else {
		w.Write([]byte("login get called"))

	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		w.Write([]byte("register post called"))
	} else {
		w.Write([]byte("register get called"))
	}
}
