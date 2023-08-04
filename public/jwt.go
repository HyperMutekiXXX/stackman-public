package public

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

const super = "HyperMutekiXXX"
const SignedKey = "stackmanxxx"

type Payload struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
}

type Claims struct {
	Payload
	jwt.RegisteredClaims
}

func Signed(id int32, name string) (string, error) {
	c := Claims{
		Payload: Payload{
			Id:   id,
			Name: name,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: &jwt.NumericDate{time.Now()},
			ExpiresAt: &jwt.NumericDate{time.Now().Add(time.Minute * 5)},
			Issuer:    super,
		},
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	return claims.SignedString([]byte(SignedKey))
}

func ParseClaims(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SignedKey), nil
	})

	return token.Claims.(*Claims), err
}
