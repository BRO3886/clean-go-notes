package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/BRO3886/clean-go-notes/pkg/note"

	"github.com/BRO3886/clean-go-notes/api/handler"
	"github.com/BRO3886/clean-go-notes/pkg/user"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	log.Println("start test server")
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal("error conneting to DB")
	}
	defer db.Close()
	log.Println("connected to db")
	db.LogMode(false)
	db.AutoMigrate(&user.User{}, &note.Note{})

	userRepo := user.NewSqliteRepo(db)
	userSvc := user.NewService(userRepo)

	noteRepo := note.NewSqliteRepo(db)
	noteSvc := note.NewService(noteRepo)

	myRouter := mux.NewRouter().StrictSlash(true)

	handler.MakeUserHandlers(myRouter, userSvc)
	handler.MakeNotesHandler(myRouter, noteSvc)

	myRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Helth"))
		return
	})
	fmt.Println("All ok. Serving...")
	log.Fatal(http.ListenAndServe("localhost:3000", myRouter))
}
