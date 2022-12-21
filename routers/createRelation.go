package routers

import (
	"net/http"

	"github.com/lucassingh/go-backend/db"
	"github.com/lucassingh/go-backend/models"
)

/*CreateRelation*/
func CreateRelation(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "You must send ID parameter", http.StatusBadRequest)
		return
	}

	var t models.Relation

	t.UserID = IDUser

	t.UserRelation = ID

	status, err := db.InsertRelation(t)

	if err != nil {
		http.Error(w, "An error ocurred when try insert relation"+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "Can not insert relation"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "aplication/json")

	w.WriteHeader(http.StatusCreated)
}
