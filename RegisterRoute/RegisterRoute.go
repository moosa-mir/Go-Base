package registerRoute

import (
	"fmt"
	"log"
	account "myproject/Handler/AccountInfo"
	login "myproject/Handler/Login"
	register "myproject/Handler/Register"
	"net/http"
)

// RegisterRoutes registers the login route with the HTTP server
func RegisterRoutes() {
	http.HandleFunc("/login", login.LoginHandler)
	http.HandleFunc("/accountInfo/", account.AccountInfoHandler)
	http.HandleFunc("/register", register.RegisterHandler)
	// Start the HTTP server on port 8080
	fmt.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
