package middlew

import (
	"net/http"

	"github.com/lucassingh/go-backend/db"
)

/*CheckDB validate connection to database when call route */
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.CheckConnection() == 0 {
			http.Error(w, "Connection lost with DB", 500)
		}
		next.ServeHTTP(w, r)
	}
}
