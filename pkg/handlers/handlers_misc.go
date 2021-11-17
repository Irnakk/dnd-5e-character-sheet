package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"dnd-5e-character-sheet/pkg/csdata"
	"dnd-5e-character-sheet/pkg/dice"
)

// Function outputs the request info into the terminal output:
//  * Request Method
//  * Request URI
func printRequestInfo(req *http.Request) {
	fmt.Println("----------------")
	fmt.Printf("req.Method:\t%v\n", req.Method)
	fmt.Printf("req.RequestURI:\t%v\n", req.RequestURI)
	fmt.Println("----------------")
}

// Function performs <Number> rolls of <Value>-sided dice.
//
// Input:
//
// Function receives {Number int, Value int} structure in the
// request body.
//
// Output:
//
// Function returns a marshaled {Number, Value, Result} object in
// the request reply.
//
// URI: /roll-result
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

// Function creates a random ("4d6 per stat") SixStats object.
//
// Input:
//
// Function does not receive any particular data in the request body.
//
// Output:
//
// Function replies with a marshaled SixStats object.
//
// URI: /roll-stats
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

// Function attempts to read the sheet data from a DB if the
// sheet id is numerical, otherwise the data is read from a JSON file.
//
// Input:
//
// Function receives {Identifier string} in the request body.
//
// Output:
//
// Function replies with a marshaled CharacterSheet object.
//
// URI: /read-sheet
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

	var (
		sheet     csdata.CharacterSheet
		byteValue []byte // the marshaled response will be stored here
	)

	// If the sheet identifier is not a number, we read the data
	// from the corresponding JSON file.
	// Otherwise (id is a number), we read the data from the
	// correspodning DB row.
	if id, err := strconv.Atoi(requestReply.Identifier); err != nil {
		fmt.Printf("Error in parsing int in identifier <<%s>>:\t%v\n", requestReply.Identifier, err)
		fmt.Printf("Trying to read data/%s.json\n", requestReply.Identifier)

		jsonFile, err := os.Open("data/" + requestReply.Identifier + ".json")
		if err != nil {
			fmt.Printf("Error in opening file:\t%v\n", err)
			return
		}
		defer jsonFile.Close()

		byteValue, err = ioutil.ReadAll(jsonFile)
		if err != nil {
			fmt.Printf("Error in reading file:\t%v\n", err)
			return
		}
	} else {
		fmt.Printf("Reading from DB, id=%d\n", id)

		err = sheet.ReadFromDB(id)
		if err != nil {
			fmt.Printf("Error in reading from DB, id=%d:\t%v\n", id, err)
			return
		}

		byteValue, err = json.MarshalIndent(sheet, "", "	")
		if err != nil {
			fmt.Printf("Error in marshalling sheet:\t%v\n", err)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json") // Not really sure why this is required
	w.Write(byteValue)
}

// Function updates a file with a given ID, saving stats base and bonuses.
//
// Input:
//
// Function receives {Identifier string, StatsBase SixStats,
// StatsBonuses SixStats} object in the request body.
//
// Output:
//
// Function does not return any particular data as a request reply.
//
// URI: /write-stats
func WriteStatsHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("In WriteStatsHandler()")
	defer fmt.Print("Out of WriteStatsHandler()\n\n****************\n\n")
	printRequestInfo(req)

	var requestReply struct {
		Identifier   string
		StatsBase    csdata.SixStats
		StatsBonuses csdata.SixStats
	}

	err := json.NewDecoder(req.Body).Decode(&requestReply)
	if err != nil {
		fmt.Printf("Error in decoding stats info:\t%v\n", err)
		return
	}

	fmt.Printf("~~~~~~~~~~~~~~~~\nrequestReply:\n%v\n~~~~~~~~~~~~~~~~\n", requestReply)

	fmt.Printf("Parsed id:\t%v\n", requestReply.Identifier)

	var sheet csdata.CharacterSheet

	if err = sheet.ReadFromFile(requestReply.Identifier); err == nil {
		fmt.Printf("Read from file data/%s.json successfully!\n", requestReply.Identifier)
	}

	// Does not matter whether we could read the sheet from the file or not,
	// we still assign the new values to the fields

	sheet.StatsBase = requestReply.StatsBase
	sheet.StatsBonuses = requestReply.StatsBonuses
	sheet.Update()
	err = sheet.WriteFile(requestReply.Identifier)
	if err != nil {
		fmt.Printf("Error in writing sheet to file:\t%v\n", err)
		return
	}
}

// Function updates a file with a given ID, saving saving throws
// proficiency and modifiers.
//
// Input:
//
// Function receives {Identifier string, STProficiency SixStatsCheck}
// object in the request body.
//
// Output:
//
// Function does not return any particular data as a request reply.
//
// URI: /write-st
func WriteSTHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("In WriteSTHandler()")
	defer fmt.Print("Out of WriteSTHandler()\n\n****************\n\n")
	printRequestInfo(req)

	var requestReply struct {
		Identifier    string
		STProficiency csdata.SixStatsCheck
	}

	err := json.NewDecoder(req.Body).Decode(&requestReply)
	if err != nil {
		fmt.Printf("Error in decoding saving throws info:\t%v\n", err)
		return
	}

	fmt.Printf("~~~~~~~~~~~~~~~~\nrequestReply:\n%v\n~~~~~~~~~~~~~~~~\n", requestReply)

	fmt.Printf("Parsed id:\t%v\n", requestReply.Identifier)

	var sheet csdata.CharacterSheet

	if err = sheet.ReadFromFile(requestReply.Identifier); err == nil {
		fmt.Printf("Read from file data/%s.json successfully!\n", requestReply.Identifier)
	}

	// Does not matter whether we could read the sheet from the file or not,
	// we still assign the new values to the fields

	sheet.STProficiency = requestReply.STProficiency
	sheet.Update()
	err = sheet.WriteFile(requestReply.Identifier)
	if err != nil {
		fmt.Printf("Error in writing sheet to file:\t%v\n", err)
		return
	}
}

// Function updates a file with a given ID, saving skills modifiers.
//
// Input:
//
// Function receives {Identifier string, SkillsProficiency SkillsCheck}
// object in the request body.
//
// Output:
//
// Function does not return any particular data as a request reply.
//
// URI: /write-skills
func WriteSkillsHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("In WriteSkillsHandler()")
	defer fmt.Print("Out of WriteSkillsHandler()\n\n****************\n\n")
	printRequestInfo(req)

	var requestReply struct {
		Identifier        string
		SkillsProficiency csdata.SkillsCheck
	}

	err := json.NewDecoder(req.Body).Decode(&requestReply)
	if err != nil {
		fmt.Printf("Error in decoding skills info:\t%v\n", err)
		return
	}

	fmt.Printf("~~~~~~~~~~~~~~~~\nrequestReply:\n%v\n~~~~~~~~~~~~~~~~\n", requestReply)

	fmt.Printf("Parsed id:\t%v\n", requestReply.Identifier)

	var sheet csdata.CharacterSheet

	if err = sheet.ReadFromFile(requestReply.Identifier); err == nil {
		fmt.Printf("Read from file data/%s.json successfully!\n", requestReply.Identifier)
	}

	fmt.Printf("CharacterSheet:\n%v\n", sheet)

	// Does not matter whether we could read the sheet from the file or not,
	// we still assign the new values to the fields

	sheet.SkillsProficiency = requestReply.SkillsProficiency
	sheet.Update()
	err = sheet.WriteFile(requestReply.Identifier)
	if err != nil {
		fmt.Printf("Error in writing sheet to file:\t%v\n", err)
		return
	}
}
