package routers

import (
	"errors"
	"strings"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/lucassingh/go-backend/db"
	"github.com/lucassingh/go-backend/models"
)

/*Email*/
var Email string

/*IDUser*/
var IDUser string

/*ProcessToken*/
func ProcessToken(tk string) (*models.Claim, bool, string, error) {

	key := []byte("TH!s!SDev3L0pmenTK3eyFoRc0Urs3GO")

	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")

	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("token format invalid")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err == nil {
		_, find, _ := db.CheckUserExist(claims.Email)

		if find == true {
			Email = claims.Email
			IDUser = claims.ID.Hex()
		}

		return claims, find, IDUser, nil
	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("Invalid token")
	}

	return claims, false, string(""), err
}
