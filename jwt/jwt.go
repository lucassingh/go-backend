package jwt

import (
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/lucassingh/go-backend/models"
)

/*GenerateJWT*/
func GenerateJWT(t models.User) (string, error) {

	privateKey := []byte("TH!s!SDev3L0pmenTK3eyFoRc0Urs3GO")

	payload := jwt.MapClaims{
		"_id":      t.ID,
		"name":     t.Name,
		"lastName": t.LastName,
		"email":    t.Email,
		"dateBorn": t.DateBorn,
		"bio":      t.Bio,
		"location": t.Location,
		"webSite":  t.WebSite,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenStr, err := token.SignedString(privateKey)

	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
