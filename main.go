package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Adresse struct {
	Rue   string `json:"rue"`
	Ville string `json:"ville"`
	Pays  string `json:"pays,omitempty"`
}

type User struct {
	Name string `json:"name"`
	// on ignore le password en json avec : -
	Password string  `json:"-"`
	Email    string  `json:"email"`
	Adresse  Adresse `json:"adresse"`
}

var users = []User{
	{
		Name:     "Samuel",
		Password: "secret",
		Email:    "samuel.michaux@gmail.com",
		Adresse: Adresse{
			Rue:   "17 chemin du moulin",
			Ville: "Comines",
			Pays:  "France",
		},
	},
	{
		Name:     "Sabrina",
		Password: "supersecret",
		Email:    "sabrina@sab.com",
		Adresse: Adresse{
			Rue:   "17 chemin du moulin",
			Ville: "Comines",
		},
	},
}

func user(w http.ResponseWriter, r *http.Request) {

	// encodage du tableau User en JSON
	b, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

type PasswordJsonBody struct {
	UserIndex         int    `json:"user_index"`
	OldPassword       string `json:"old_password"`
	NewPassword       string `json:"new_password"`
	NewPasswordRepeat string `json:"new_password_repeat"`
}

func updatePassword(w http.ResponseWriter, r *http.Request) {
	var p PasswordJsonBody

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("Update password parsed: %v\n", p)

	if p.UserIndex < 0 || p.UserIndex > len(users)-1 {
		msg := fmt.Sprintf("Invalid index. got user_index=%v, valid range=[0, %v]",
			p.UserIndex, len(users)-1)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	u := users[p.UserIndex]
	if u.Password != p.OldPassword {
		http.Error(w, "Old password do not match", http.StatusBadRequest)
		return
	}

	if p.NewPassword != p.NewPasswordRepeat {
		http.Error(w, "New passwords do not match", http.StatusBadRequest)
		return
	}

	u.Password = p.NewPassword

	fmt.Fprintf(w, "Password updated")
}
func main() {
	http.HandleFunc("/users", user)
	http.HandleFunc("/update_password", updatePassword)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
