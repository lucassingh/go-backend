package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/lucassingh/go-backend/db"
	"github.com/lucassingh/go-backend/models"
)

/*UploadAvatar*/
func UploadAvatar(w http.ResponseWriter, r *http.Request) {

	file, handler, err := r.FormFile("avatar")

	var extension = strings.Split(handler.Filename, ".")[1]

	var uploadsFIle = "uploads/avatars/" + IDUser + "." + extension

	f, err := os.OpenFile(uploadsFIle, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		http.Error(w, "An error ocurred when try upload avatar. Try again later..."+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)

	if err != nil {
		http.Error(w, "An error ocurred when try copy avatar. Try again later..."+err.Error(), http.StatusBadRequest)
		return
	}

	var user models.User

	var status bool

	user.Avatar = IDUser + "." + extension

	status, err = db.UpdateRegister(user, IDUser)

	if err != nil || status == false {
		http.Error(w, "An error ocurred when try create avatar in Data base. Try again later..."+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "aplication/json")

	w.WriteHeader(http.StatusCreated)
}
