package main

import (
	"log"
	"net/http"
	"os"

	"dnd-5e-character-sheet/pkg/handlers"
)

func main() {
	log.Println("Starting func main()...")

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

	log.Print("Listening on port 8080.\n\n")
	http.ListenAndServe(":8080", nil)
}
