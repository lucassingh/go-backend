package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/lucassingh/go-backend/db"
)

/*ReadTwts*/
func ReadTwts(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "You must send id parameter", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "You must send page parameter", http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))

	if err != nil {
		http.Error(w, "You must send page parameter. Value should be greater than 0", http.StatusBadRequest)
		return
	}

	pag := int64(page)

	response, success := db.ReadTweets(ID, pag)

	if !success {
		http.Error(w, "Error to try get all tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "aplication/json")

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(response)
}
