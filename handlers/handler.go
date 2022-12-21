package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/lucassingh/go-backend/middlew"
	"github.com/lucassingh/go-backend/routers"
	"github.com/rs/cors"
)

/*HandlersCon organize conn*/
func HandlersCon() {

	router := mux.NewRouter()

	router.HandleFunc("/register", middlew.CheckDB(routers.Register)).Methods("POST")

	router.HandleFunc("/login", middlew.CheckDB(routers.Login)).Methods("POST")

	router.HandleFunc("/view-profile", middlew.CheckDB(middlew.ValidateJWT(routers.ViewProfile))).Methods("GET")

	router.HandleFunc("/update-profile", middlew.CheckDB(middlew.ValidateJWT(routers.UpdateProfile))).Methods("PUT")

	router.HandleFunc("/create-tweet", middlew.CheckDB(middlew.ValidateJWT(routers.CreateTweet))).Methods("POST")

	router.HandleFunc("/read-tweets", middlew.CheckDB(middlew.ValidateJWT(routers.ReadTwts))).Methods("GET")

	router.HandleFunc("/delete-tweet", middlew.CheckDB(middlew.ValidateJWT(routers.DeleteTwt))).Methods("DELETE")

	router.HandleFunc("/upload-avatar", middlew.CheckDB(middlew.ValidateJWT(routers.UploadAvatar))).Methods("POST")

	router.HandleFunc("/get-avatar", middlew.CheckDB(middlew.ValidateJWT(routers.GetAvatar))).Methods("GET")

	router.HandleFunc("/upload-banner", middlew.CheckDB(middlew.ValidateJWT(routers.UploadBanner))).Methods("POST")

	router.HandleFunc("/get-banner", middlew.CheckDB(middlew.ValidateJWT(routers.GetBanner))).Methods("GET")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
