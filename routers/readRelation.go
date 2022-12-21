package routers

import (
	"encoding/json"
	"net/http"

	"github.com/lucassingh/go-backend/db"
	"github.com/lucassingh/go-backend/models"
)

func ReadRelation(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	var t models.Relation

	t.UserID = IDUser

	t.UserRelation = ID

	var resp models.ResponseRelation

	status, err := db.ReadRelation(t)

	if err != nil || status == false {
		resp.Status = false
	} else {
		resp.Status = true
	}

	w.Header().Set("Content-type", "aplication/json")

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(resp)
}
