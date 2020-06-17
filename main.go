package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/BRO3886/clean-go-notes/pkg/note"

	"github.com/BRO3886/clean-go-notes/api/handler"
	"github.com/BRO3886/clean-go-notes/pkg/user"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func connectDB() (*gorm.DB, error) {

	//Heroku
	if os.Getenv("DATABASE_URL") != "" {
		return gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	}
	return gorm.Open("sqlite3", "test.db")
}

func main() {
	log.Println("start test server")

	db, err := connectDB()
	if err != nil {
		log.Fatal("error conneting to DB")
	}
	log.Println("connected to db")

	defer db.Close()

	if os.Getenv("DB_LOG_MODE") == "true" {
		db.LogMode(true)
	}

	db.AutoMigrate(&user.User{}, &note.Note{})

	userRepo := user.NewSqliteRepo(db)
	userSvc := user.NewService(userRepo)

	noteRepo := note.NewSqliteRepo(db)
	noteSvc := note.NewService(noteRepo)

	myRouter := mux.NewRouter().StrictSlash(true)

	handler.MakeUserHandlers(myRouter, userSvc)
	handler.MakeNotesHandler(myRouter, noteSvc)

	//health route
	myRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ping"))
		return
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "1729"
	}

	log.Println("All ok. Serving on port=", port)

	// log.Fatal(http.ListenAndServe("localhost:3000", myRouter))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), myRouter))
}
