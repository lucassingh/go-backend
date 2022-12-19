package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/lucassingh/go-backend/db"
	"github.com/lucassingh/go-backend/models"
)

/*CreateTweet*/
func CreateTweet(w http.ResponseWriter, r *http.Request) {

	var message models.Tweet

	err := json.NewDecoder(r.Body).Decode(&message)

	register := models.InsertTweet{
		UserID:  IDUser,
		Message: message.Message,
		Date:    time.Now(),
	}

	_, status, err := db.InsertTweet(register)

	if err != nil {
		http.Error(w, "An error ocurred when try create a tweet. Try again later..."+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "Cannot create a tweet", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
