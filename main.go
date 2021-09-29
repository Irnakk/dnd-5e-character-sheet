package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"dnd-5e-character-sheet/src/csdata"
	"dnd-5e-character-sheet/src/dice"
	"dnd-5e-character-sheet/src/handlers"
)

func testWrite() error {
	sheet := csdata.CharacterSheet{}
	sheet.ProficiencyBonus = 2
	sheet.StatsBase, _ = dice.RollStats()

	return sheet.WriteFile("test.txt")
}

func main() {
	// err := testWrite()
	// if err != nil {
	// 	fmt.Printf("Error in writing file:\t%v\n", err)
	// 	return
	// }

	fmt.Println("Starting func main()...")
	fmt.Printf("at:\t%v\n", time.Now())
	wd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error in getting current working directory:\t%v\n", err)
	} else {
		fmt.Printf("from:\t%v\n", wd)
	}

	defer fmt.Printf("Finishing func main() at:\t%v\n", time.Now())

	fmt.Println("Binding handler functions...")
	http.HandleFunc("/", handlers.RootHandler)
	http.HandleFunc("/roll", handlers.RollHandler)
	http.HandleFunc("/stats", handlers.StatsHandler)
	http.HandleFunc("/6stats", handlers.SixStatsHandler)
	http.HandleFunc("/favicon.ico", handlers.IconHandler)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("web")))) // Not entirely sure why it is written like this
	http.Handle("/scripts/", http.StripPrefix("/scripts/", http.FileServer(http.Dir("web"))))
	http.HandleFunc("/roll-result", handlers.RollResultHandler)
	http.HandleFunc("/saving-throws", handlers.SavingThrowsHandler)

	http.HandleFunc("/roll-stats", handlers.RollStatsHandler)
	http.HandleFunc("/read-sheet", handlers.ReadSheetHandler)
	http.HandleFunc("/write-sheet", handlers.WriteSheetHandler)

	fmt.Print("Listening on port 8080.\n\n")
	http.ListenAndServe(":8080", nil)
}
