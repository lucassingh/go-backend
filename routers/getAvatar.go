package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/lucassingh/go-backend/db"
)

/*GetAvatar*/

func GetAvatar(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "You must send ID parameter", http.StatusBadRequest)
		return
	}

	profile, err := db.FindProfile(ID)

	if err != nil {
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}

	OpenFile, err := os.Open("uploads/avatars/" + profile.Avatar)

	if err != nil {
		http.Error(w, "Avatar not found", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, OpenFile)

	if err != nil {
		http.Error(w, "An error ocurred when try copy avatar. Try again later...", http.StatusBadRequest)
	}
}
