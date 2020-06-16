package handler

import (
	"encoding/json"
	"net/http"

	"github.com/BRO3886/clean-go-notes/utils"

	"github.com/BRO3886/clean-go-notes/api/middleware"
	"github.com/BRO3886/clean-go-notes/pkg/note"
	"github.com/gorilla/mux"
)

func createNote(svc note.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		note := &note.Note{}
		if err := json.NewDecoder(r.Body).Decode(note); err != nil {
			utils.ResponseWrapper(w, http.StatusBadRequest, err.Error())
			return
		}

		ctx := r.Context()
		tk := ctx.Value(middleware.JwtContextKey("token")).(*middleware.Token)

		note.UserID = tk.ID

		n, err := svc.CreateNote(note)
		if err != nil {
			utils.ResponseWrapper(w, http.StatusConflict, err.Error())
			return
		}
		utils.JsonifyHeader(w)
		w.WriteHeader(http.StatusCreated)
		utils.WrapData(w, map[string]interface{}{
			"message": "note created",
			"note":    n,
		})
	}
}

func fetchAllNotes(svc note.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()
		tk := ctx.Value(middleware.JwtContextKey("token")).(*middleware.Token)

		notes, err := svc.FetchAllNotes(tk.ID)
		if err != nil {
			utils.ResponseWrapper(w, http.StatusBadRequest, err.Error())
		}
		utils.JsonifyHeader(w)
		w.WriteHeader(http.StatusOK)
		utils.WrapData(w, map[string]interface{}{
			"message": "notes found",
			"notes":   notes,
		})

	}
}

//MakeNotesHandler handles routes related to notes
func MakeNotesHandler(r *mux.Router, svc note.Service) {
	r.HandleFunc("/api/notes/create", middleware.JwtAuth(createNote(svc))).Methods(http.MethodPost)
	r.HandleFunc("/api/notes", middleware.JwtAuth(fetchAllNotes(svc))).Methods(http.MethodGet)
}
