package middleware

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/BRO3886/clean-go-notes/pkg"
	"github.com/dgrijalva/jwt-go"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
)

//Validate validates jwt
func Validate(handler http.Handler) http.Handler {
	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("jwtsecret")), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})
	return jwtMiddleware.Handler(handler)
}

//ValidateAndGetClaims token creation
func ValidateAndGetClaims(ctx context.Context, role string) (map[string]interface{}, error) {
	token, ok := ctx.Value("user").(*jwt.Token)
	if !ok {
		log.Println(token)
		return nil, pkg.ErrInvalidToken
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		log.Println(claims)
		return nil, pkg.ErrInvalidToken
	}

	if claims["role"].(string) != role {
		log.Println(claims["role"])
		return nil, pkg.ErrUnauthorised
	}
	return claims, nil
}
