package routers

import (
	"net/http"

	"github.com/lucassingh/go-backend/db"
	"github.com/lucassingh/go-backend/models"
)

/*DeleteRelation*/
func DeleteRelation(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	var t models.Relation

	t.UserID = IDUser

	t.UserRelation = ID

	status, err := db.DeleteRelation(t)

	if err != nil {
		http.Error(w, "An error ocurred when try delete relation"+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "Can not delete relation"+err.Error(), http.StatusBadRequest)
		return
	}
}
