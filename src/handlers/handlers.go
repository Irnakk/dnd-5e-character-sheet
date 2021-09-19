package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"dnd-5e-character-sheet/src/dice"
)

func printRequestInfo(req *http.Request) {
	fmt.Println("----------------")
	fmt.Printf("req.Method:\t%v\n", req.Method)
	fmt.Printf("req.RequestURI:\t%v\n", req.RequestURI)
	fmt.Println("----------------")
}

func IconHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("In iconHandler()")
	defer fmt.Print("Out of iconHandler()\n\n****************\n\n")
	http.ServeFile(w, req, "web/content/icon.png")
}

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

func RollResultHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("In rollResultHandler()")
	defer fmt.Print("Out of rollResultHandler()\n\n****************\n\n")
	printRequestInfo(req)

	type rollInfo struct {
		Number int
		Value  int
	}

	var roll rollInfo

	err := json.NewDecoder(req.Body).Decode(&roll)
	if err != nil {
		fmt.Printf("Error in decoding request body:\t%v\n", err)
		return
	}

	response_data, err := dice.MarshalRollSingle(roll.Number, roll.Value)
	if err != nil {
		fmt.Printf("Could not marshal a roll:\t%v\n", err)
		return
	}

	w.Header().Set("Content-Type", "application/json") // Not really sure why this is required
	w.Write(response_data)
}

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

func RollStatsHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("In RollStatsHandler()")
	defer fmt.Print("Out of RollStatsHandler()\n\n****************\n\n")
	printRequestInfo(req)

	statsObj, err := dice.RollStats()
	if err != nil {
		fmt.Printf("Error in rolling stats:\t%v\n", err)
		return
	}

	response_data, err := json.MarshalIndent(statsObj, "", "	")
	if err != nil {
		fmt.Printf("Error in marshalling statsObj:\t%v\n", err)
		return
	}

	w.Header().Set("Content-Type", "application/json") // Not really sure why this is required
	w.Write(response_data)
}
func SixStatsHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("In SixStatsHandler()")
	defer fmt.Print("Out of SixStatsHandler()\n\n****************\n\n")
	printRequestInfo(req)

	t, err := template.ParseFiles("web/6stats.html")
	if err != nil {
		fmt.Printf("Error in parsing 6stats.html:\t%v\n", err)
		return
	}
	t.Execute(w, nil)
}
