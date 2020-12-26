package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Nerzal/gocloak/v3"

	"github.com/sergiogoh/go-vue-keycloak/backend/internal/user"
)

// Authenticate tries to authenticate a user by its given username and password
func Authenticate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var u user.LoginUser
	json.NewDecoder(r.Body).Decode(&u)
	fmt.Println("decoded to user", u)

	client := gocloak.NewClient("http://localhost:8000")
	// token, err := client.Login("vue-go-keycloak-auth", "", "master", u.Username, u.Password)
	token, err := client.Login("admin-cli", "", "master", u.Username, u.Password)

	if err != nil {
		http.Error(w, "Forbidden", http.StatusForbidden)
	} else {
		fmt.Println("token: ", token)
		json.NewEncoder(w).Encode(token)
	}

}
