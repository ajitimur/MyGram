package helpers

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"errors"
	"fmt"
	"strings"

	"github.com/beego/beego/v2/server/web/context"
	jwt "github.com/dgrijalva/jwt-go"
)

var SecretKey *ecdsa.PrivateKey

func init() {
	PrivateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}

	SecretKey = PrivateKey
}

func GenerateToken(id int, email string) string {
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
	}

	// secretByte := []byte(SecretKey)

	parseToken := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	signedToken, err := parseToken.SignedString(SecretKey) //error aneh banget, minta type *ecdsa.PrivateKey
	if err != nil {
		panic(err)
	}

	return signedToken
}

func VerifyToken(c *context.Context) (interface{}, error) {
	errRes := errors.New("sign in to proceed")
	headerToken := c.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")

	if !bearer {
		return nil, errRes
	}

	stringToken := strings.Split(headerToken, " ")[1]

	token, _ := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errRes
		}
		return SecretKey, nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, errRes
	}

	return token.Claims.(jwt.MapClaims), nil
}

func ValidateToken(encodedToken string) (*jwt.Token, error) {
	fmt.Println(encodedToken, "<<<<<")
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			fmt.Println("not ok")
			return nil, errors.New("Invalid token")
		}

		return SecretKey, nil
	})

	fmt.Println(token)
	if err != nil {
		fmt.Println("err")
		return token, err
	}

	return token, nil
}
