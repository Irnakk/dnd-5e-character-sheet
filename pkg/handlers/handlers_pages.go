/*
Package handlers keeps the handler functions for each URI:
    /
	/roll
	/stats
	/favicon.ico
	/css/XXX
	/scripts/XXX
	/roll-result
	/saving-throws
	/skills
	/register
	/login
	/sheet
	/roll-stats
	/read-sheet
	/write-stats
	/write-st
*/
package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

// Function serves an icon file.
//
// URI: /favicon.ico
func IconHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("In iconHandler()")
	defer fmt.Print("Out of iconHandler()\n\n****************\n\n")
	http.ServeFile(w, req, "web/content/icon.png")
}

// Function serves "index.html" page.
//
// URI: /
func RootHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("In rootHandler()")
	defer fmt.Print("Out of rootHandler()\n\n****************\n\n")
	printRequestInfo(req)

	t, err := template.ParseFiles("web/index.html")
	if err != nil {
		fmt.Printf("Error in parsing index.html:\t%v\n", err)
		return
	}
	t.Execute(w, nil)
}

// Function serves "roll.html" page.
//
// URI: /roll
func RollHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("In rollHandler()")
	defer fmt.Print("Out of rollHandler()\n\n****************\n\n")
	printRequestInfo(req)

	t, err := template.ParseFiles("web/roll.html")
	if err != nil {
		fmt.Printf("Error in parsing roll.html:\t%v\n", err)
		return
	}
	t.Execute(w, nil)
}

// Function serves "stats.html" page.
//
// URI: /stats
func StatsHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("In StatsHandler()")
	defer fmt.Print("Out of StatsHandler()\n\n****************\n\n")
	printRequestInfo(req)

	t, err := template.ParseFiles("web/stats.html")
	if err != nil {
		fmt.Printf("Error in parsing stats.html:\t%v\n", err)
		return
	}
	t.Execute(w, nil)
}

// Function serves "saving_throws.html" page.
//
// URI: /saving-throws
func SavingThrowsHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("In SavingThrowsHandler()")
	defer fmt.Print("Out of SavingThrowsHandler()\n\n****************\n\n")
	printRequestInfo(req)

	t, err := template.ParseFiles("web/saving_throws.html")
	if err != nil {
		fmt.Printf("Error in parsing saving_throws.html:\t%v\n", err)
		return
	}
	t.Execute(w, nil)
}

// Function serves "register.html" page.
//
// URI: /register
func RegisterHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("In RegisterHandler()")
	defer fmt.Print("Out of RegisterHandler()\n\n****************\n\n")
	printRequestInfo(req)

	t, err := template.ParseFiles("web/register.html")
	if err != nil {
		fmt.Printf("Error in parsing register.html:\t%v\n", err)
		return
	}
	t.Execute(w, nil)
}

// Function serves "login.html" page.
//
// URI: /login
func LoginHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("In LoginHandler()")
	defer fmt.Print("Out of LoginHandler()\n\n****************\n\n")
	printRequestInfo(req)

	t, err := template.ParseFiles("web/login.html")
	if err != nil {
		fmt.Printf("Error in parsing login.html:\t%v\n", err)
		return
	}
	t.Execute(w, nil)
}

// Function serves "skills.html" page.
//
// URI: /skills
func SkillsHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("In SkillsHandler()")
	defer fmt.Print("Out of SkillsHandler()\n\n****************\n\n")
	printRequestInfo(req)

	t, err := template.ParseFiles("web/skills.html")
	if err != nil {
		fmt.Printf("Error in parsing skills.html:\t%v\n", err)
		return
	}
	t.Execute(w, nil)
}

// Function serves "character_sheet.html" page.
//
// URI: /sheet
func CharacterSheetHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("In CharacterSheetHandler()")
	defer fmt.Print("Out of CharacterSheetHandler()\n\n****************\n\n")
	printRequestInfo(req)

	t, err := template.ParseFiles("web/character_sheet.html")
	if err != nil {
		fmt.Printf("Error in parsing character_sheet.html:\t%v\n", err)
		return
	}
	t.Execute(w, nil)
}
