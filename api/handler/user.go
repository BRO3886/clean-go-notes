package handler

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/BRO3886/clean-go-notes/pkg/user"
	"github.com/BRO3886/clean-go-notes/utils"
	"github.com/badoux/checkmail"
	"github.com/dgrijalva/jwt-go"
)

func regsiter(svc user.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			utils.ResponseWrapper(w, http.StatusMethodNotAllowed, "invalid request type")
			return
		}
		user := &user.User{}
		var err error
		if err = json.NewDecoder(r.Body).Decode(user); err != nil {
			utils.ResponseWrapper(w, http.StatusBadRequest, err.Error())
			return
		}

		if err = checkmail.ValidateFormat(user.Email); err != nil {
			utils.ResponseWrapper(w, http.StatusBadRequest, "inavlid email")
			return
		}
		if len(user.Name) == 0 {
			utils.ResponseWrapper(w, http.StatusBadRequest, "name should not be empty")
			return
		}
		if len(user.Password) == 0 {
			utils.ResponseWrapper(w, http.StatusBadRequest, "password should not be empty")
			return
		}

		user, err = svc.Register(user)

		if err != nil {
			utils.ResponseWrapper(w, http.StatusConflict, err.Error())
			return
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id":   user.Email,
			"role": "user",
		})

		tokenString, err := token.SignedString([]byte(os.Getenv("jwtsecret")))
		if err != nil {
			utils.ResponseWrapper(w, http.StatusConflict, err.Error())
			return
		}

		w.WriteHeader(http.StatusCreated)
		utils.WrapData(w, map[string]interface{}{
			"message": "account created",
			"token":   tokenString,
			"user":    *user,
		})
	})

}

func login(svc user.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		return
	})

}
