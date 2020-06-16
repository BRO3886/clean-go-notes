package handler

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/BRO3886/clean-go-notes/api/middleware"
	"github.com/BRO3886/clean-go-notes/pkg/user"
	"github.com/BRO3886/clean-go-notes/utils"
	"github.com/badoux/checkmail"
	"github.com/dgrijalva/jwt-go"
)

func regsiter(svc user.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		token := middleware.Token{Email: user.Email, ID: uint64(user.ID)}

		tkString := jwt.NewWithClaims(jwt.SigningMethodHS512, token)

		tokenString, err := tkString.SignedString([]byte(os.Getenv("jwtsecret")))
		if err != nil {
			utils.ResponseWrapper(w, http.StatusConflict, err.Error())
			return
		}

		user.Password = ""

		utils.JsonifyHeader(w)
		w.WriteHeader(http.StatusCreated)
		utils.WrapData(w, map[string]interface{}{
			"message": "account created",
			"token":   tokenString,
			"user":    *user,
		})
	}

}

func login(svc user.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		user, err = svc.Login(user.Email, user.Password)
		if err != nil {
			utils.ResponseWrapper(w, http.StatusBadRequest, err.Error())
		}

		token := middleware.Token{Email: user.Email, ID: uint64(user.ID)}

		tkString := jwt.NewWithClaims(jwt.SigningMethodHS512, token)

		tokenString, err := tkString.SignedString([]byte(os.Getenv("jwtsecret")))
		if err != nil {
			utils.ResponseWrapper(w, http.StatusConflict, err.Error())
			return
		}

		user.Password = ""

		utils.JsonifyHeader(w)
		w.WriteHeader(http.StatusOK)
		utils.WrapData(w, map[string]interface{}{
			"message": "login successful",
			"token":   tokenString,
			"user":    *user,
		})
	}
}

func userDetails(svc user.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		tk := ctx.Value(middleware.JwtContextKey("token")).(*middleware.Token)

		user, err := svc.GetUserByID(tk.ID)
		if err != nil {
			utils.ResponseWrapper(w, http.StatusConflict, err.Error())
			return
		}

		utils.JsonifyHeader(w)
		w.WriteHeader(http.StatusFound)
		utils.WrapData(w, map[string]interface{}{
			"message": "user found",
			"user":    user,
		})

	}
}

//MakeUserHandlers handlers for user related route
func MakeUserHandlers(r *mux.Router, svc user.Service) {
	r.HandleFunc("/api/user/register", regsiter(svc)).Methods(http.MethodPost)
	r.HandleFunc("/api/user/details", middleware.JwtAuth(userDetails(svc))).Methods(http.MethodGet)
	r.HandleFunc("/api/user/login", login(svc)).Methods(http.MethodPost)
}
