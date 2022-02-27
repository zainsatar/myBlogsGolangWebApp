package main

import (
	"log"
	"main/requesthandlers"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	fs := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	r.HandleFunc("/signup", requesthandlers.ShowSignupPage).Methods("GET")
	r.HandleFunc("/signup", requesthandlers.SignupHandler).Methods("POST")
	r.HandleFunc("/login", requesthandlers.ShowLoginPage).Methods("GET")
	r.HandleFunc("/login", requesthandlers.LoginHandler).Methods("POST")
	r.HandleFunc("/home", requesthandlers.ShowHomePage).Methods("GET")
	r.HandleFunc("/addNewPost", requesthandlers.ShowAddNewPostPage).Methods("GET")
	r.HandleFunc("/addNewPost", requesthandlers.AddNewPostHandler).Methods("POST")
	r.HandleFunc("/myPosts", requesthandlers.ShowMyPostsPage).Methods("GET")
	r.HandleFunc("/logout", requesthandlers.Logout).Methods("GET")
	r.HandleFunc("/settings", requesthandlers.ShowSettingsPage).Methods("GET")
	r.HandleFunc("/settings", requesthandlers.SettingsHandler).Methods("POST")
	Port := os.Getenv("PORT")
	// log.Fatal(http.ListenAndServe(":3000", r))
	log.Fatal(http.ListenAndServe(":"+Port, r))
}
