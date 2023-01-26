package tokens

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

const (
	DAYS_TO_EXPIRE_IN = 24
)

type UserClaims struct {
	ID    int
	Email string
}

func GenerateJWT(userClaims UserClaims) (string, error) {
	expireAt := time.Now().Add(time.Duration(time.Hour) * DAYS_TO_EXPIRE_IN).Unix()
	jwtSecret := viper.Get("jwt.secret").(string)

	claims := jwt.MapClaims{
		"exp":        expireAt,
		"user_id":    userClaims.ID,
		"user_email": userClaims.Email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString([]byte(jwtSecret))

	if err != nil {
		log.Println("An error occurred when generating a JWT.")
		return "", err
	}

	return accessToken, nil
}

func ParseToken(header string) (*UserClaims, error) {
	accessToken := strings.SplitAfter(header, "Bearer")[1]
	jwtSecret := viper.Get("jwt.secret").(string)

	result, err := jwt.Parse(strings.Trim(accessToken, " "), func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})

	claims := result.Claims.(jwt.MapClaims)
	data := UserClaims{
		ID:    int(claims["user_id"].(float64)),
		Email: claims["user_email"].(string),
	}

	if err != nil {
		log.Println("An error occurred when parsing the JWT.")
		return nil, err
	}

	return &data, nil
}
