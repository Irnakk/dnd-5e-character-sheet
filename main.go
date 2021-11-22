package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"database/sql"
	"io/ioutil"

	_ "github.com/lib/pq"

	"dnd-5e-character-sheet/pkg/csdata"
	"dnd-5e-character-sheet/pkg/handlers"
)

var DB *sql.DB

func main() {
	log.Println("Starting func main()...")

	signal_channel := make(chan os.Signal)
	signal.Notify(signal_channel, os.Interrupt, syscall.SIGTERM)

	go func() {
		for sig := range signal_channel {
			log.Printf("Captured %v, stopping profiler and exiting..", sig)
			log.Println("Closing DB...")
			if err := DB.Close(); err != nil {
				log.Printf("Error in closing DB:\t%v\n", err)
			}
			log.Println("Finishing func main()")
			os.Exit(1)
		}
	}()

	wd, err := os.Getwd()
	if err != nil {
		log.Printf("Error in getting current working directory:\t%v\n", err)
	} else {
		log.Printf("from:\t%v\n", wd)
	}

	defer log.Println("Finishing func main()")

	log.Println("Binding handler functions...")
	http.HandleFunc("/", handlers.RootHandler)
	http.HandleFunc("/roll", handlers.RollHandler)
	http.HandleFunc("/stats", handlers.StatsHandler)
	http.HandleFunc("/favicon.ico", handlers.IconHandler)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("web/css"))))
	http.Handle("/scripts/", http.StripPrefix("/scripts/", http.FileServer(http.Dir("web/scripts"))))
	http.HandleFunc("/roll-result", handlers.RollResultHandler)
	http.HandleFunc("/saving-throws", handlers.SavingThrowsHandler)
	http.HandleFunc("/skills", handlers.SkillsHandler)
	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/sheet", handlers.CharacterSheetHandler)

	http.HandleFunc("/roll-stats", handlers.RollStatsHandler)
	http.HandleFunc("/read-sheet", handlers.ReadSheetHandler)
	http.HandleFunc("/write-stats", handlers.WriteStatsHandler)
	http.HandleFunc("/write-st", handlers.WriteSTHandler)
	http.HandleFunc("/write-skills", handlers.WriteSkillsHandler)

	// TODO:
	// http.HandleFunc("/reg-form", handlers.WriteSheetHandler)
	// http.HandleFunc("/login-form", handlers.WriteSheetHandler)

	log.Println("Connecting to the database...")

	file_content, err := ioutil.ReadFile("C:/d5cs/db_login")
	if err != nil {
		log.Printf("Error in opening login file:\t%v\n", err)
	}

	psqlconn := string(file_content)

	DB, err = sql.Open("postgres", psqlconn)
	if err != nil {
		log.Printf("Error in opening DB:\t%v\n", err)
	}
	defer DB.Close()
	csdata.DB = DB

	log.Print("Listening on port 8080.\n\n")
	http.ListenAndServe(":8080", nil)
}
