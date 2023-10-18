package handlers

import (
	"encoding/json"
	"net/http"
)

type User struct {
	UserId  int
	Hobbies []string
}

func Home(w http.ResponseWriter, r *http.Request) {
	// w.WriteHeader(http.StatusBadRequest)
	// w.Write([]byte("hello this is our home page updated1"))
	w.Header().Set("Content-Type", "application/json")
	user1 := User{UserId: 1, Hobbies: []string{"chess", "carrom"}}
	b, err := json.Marshal(user1)
	if err != nil {
		w.Write([]byte("something went wrong"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
