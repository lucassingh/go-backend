package routers

import (
	"net/http"

	"github.com/lucassingh/go-backend/db"
)

/*DeleteTwt*/
func DeleteTwt(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "You must send id parameter", http.StatusBadRequest)
		return
	}

	err := db.DeleteTweet(ID, IDUser)

	if err != nil {
		http.Error(w, "An error ocurred to try delete this tweet"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "aplication/json")

	w.WriteHeader(http.StatusCreated)
}
