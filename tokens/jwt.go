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

func GenerateJWT(id int) (string, error) {
	expireAt := time.Now().Add(time.Duration(time.Hour) * DAYS_TO_EXPIRE_IN).Unix()
	jwtSecret := viper.Get("jwt.secret").(string)

	claims := jwt.MapClaims{
		"nbf":     expireAt,
		"user_id": id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString([]byte(jwtSecret))

	if err != nil {
		log.Println("An error occurred when generating a JWT.")
		return "", err
	}

	return accessToken, nil
}

func ParseToken(header string) (any, error) {
	accessToken := strings.SplitAfter(header, "Bearer")[1]
	jwtSecret := viper.Get("jwt.secret").(string)

	// log.Printf("access token: %s\n", accessToken)

	result, err := jwt.Parse(strings.Trim(accessToken, " "), func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})

	// log.Printf("result: %v\n", result.Valid)

	if err != nil {
		log.Println("An error occurred when parsing the JWT.")
		return nil, err
	}

	return result, nil
}
