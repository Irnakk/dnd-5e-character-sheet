package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"dnd-5e-character-sheet/pkg/csdata"
	"dnd-5e-character-sheet/pkg/dice"
)

func printRequestInfo(req *http.Request) {
	fmt.Println("----------------")
	fmt.Printf("req.Method:\t%v\n", req.Method)
	fmt.Printf("req.RequestURI:\t%v\n", req.RequestURI)
	fmt.Println("----------------")
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

func ReadSheetHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("In ReadSheetHandler()")
	defer fmt.Print("Out of ReadSheetHandler()\n\n****************\n\n")
	printRequestInfo(req)

	var requestReply struct {
		Identifier string
	}

	err := json.NewDecoder(req.Body).Decode(&requestReply)
	if err != nil {
		fmt.Printf("Error in decoding sheet id:\t%v\n", err)
		return
	}

	fmt.Printf("Parsed id:\t%v\n", requestReply.Identifier)

	jsonFile, err := os.Open("data/" + requestReply.Identifier + ".json")
	if err != nil {
		fmt.Printf("Error in opening file:\t%v\n", err)
		return
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Printf("Error in reading file:\t%v\n", err)
		return
	}

	w.Header().Set("Content-Type", "application/json") // Not really sure why this is required
	w.Write(byteValue)
}

func WriteSheetHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("In WriteSheetHandler()")
	defer fmt.Print("Out of WriteSheetHandler()\n\n****************\n\n")
	printRequestInfo(req)

	var requestReply struct {
		Identifier   string
		StatsBase    csdata.SixStats
		StatsBonuses csdata.SixStats
	}

	err := json.NewDecoder(req.Body).Decode(&requestReply)
	if err != nil {
		fmt.Printf("Error in decoding sheet info:\t%v\n", err)
		return
	}

	fmt.Printf("~~~~~~~~~~~~~~~~\nrequestReply:\n%v\n~~~~~~~~~~~~~~~~\n", requestReply)

	fmt.Printf("Parsed id:\t%v\n", requestReply.Identifier)

	var sheet csdata.CharacterSheet

	sheet.StatsBase = requestReply.StatsBase
	sheet.StatsBonuses = requestReply.StatsBonuses
	sheet.Update()
	err = sheet.WriteFile(requestReply.Identifier)
	if err != nil {
		fmt.Printf("Error in writing sheet to file:\t%v\n", err)
		return
	}
}
