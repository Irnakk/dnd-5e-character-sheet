package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"time"

	"dnd-5e-character-sheet/src/dice"
)

func iconHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("In iconHandler()")
	defer fmt.Print("Out of iconHandler()\n\n****************\n\n")
	http.ServeFile(w, req, "web/content/icon.png")
}

func printRequestInfo(req *http.Request) {
	fmt.Println("----------------")
	fmt.Printf("req.Method:\t%v\n", req.Method)
	fmt.Printf("req.RequestURI:\t%v\n", req.RequestURI)
	fmt.Println("----------------")
}

func rootHandler(w http.ResponseWriter, req *http.Request) {
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

func rollHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("In rollHandler()")
	defer fmt.Print("Out of rollHandler()\n\n****************\n\n")
	printRequestInfo(req)

	t, err := template.ParseFiles("web/roll.html")
	if err != nil {
		fmt.Printf("Error in parsing roll.html:\t%v\n", err)
		return
	}
	t.Execute(w, nil)

	switch req.Method {
	case "GET":
		if err := req.ParseForm(); err != nil {
			fmt.Printf("Error in parsing form in rollHandler():\t%v\n", err)
			return
		}

		fmt.Printf("Successful GET; req.PostForm:\t%v\n", req.PostForm)
		for key, value := range req.PostForm {
			fmt.Printf("<%v>:\t%v\n", key, value)
		}
	case "POST":
		if err := req.ParseForm(); err != nil {
			fmt.Printf("Error in parsing form in rollHandler():\t%v\n", err)
			return
		}

		fmt.Printf("Successful POST; req.PostForm:\t%v\n", req.PostForm)
		for key, value := range req.PostForm {
			fmt.Printf("<%v>:\t%v\n", key, value)
		}
	}
}

func rollAjaxHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("In rollAjaxHandler()")
	defer fmt.Print("Out of rollAjaxHandler()\n\n****************\n\n")
	printRequestInfo(req)

	t, err := template.ParseFiles("web/roll_ajax.html")
	if err != nil {
		fmt.Printf("Error in parsing roll_ajax.html:\t%v\n", err)
		return
	}
	t.Execute(w, nil)

	switch req.Method {
	case "GET":
		if err := req.ParseForm(); err != nil {
			fmt.Printf("Error in parsing form in rollAjaxHandler():\t%v\n", err)
			return
		}

		fmt.Printf("Successful GET; req.PostForm:\t%v\n", req.PostForm)
		for key, value := range req.PostForm {
			fmt.Printf("<%v>:\t%v\n", key, value)
		}
	case "POST":
		if err := req.ParseForm(); err != nil {
			fmt.Printf("Error in parsing form in rollAjaxHandler():\t%v\n", err)
			return
		}

		fmt.Printf("Successful POST; req.PostForm:\t%v\n", req.PostForm)
		for key, value := range req.PostForm {
			fmt.Printf("<%v>:\t%v\n", key, value)
		}
	}
}

func testTextHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("In testTextHandler()")
	defer fmt.Print("Out of testTextHandler()\n\n****************\n\n")
	printRequestInfo(req)

	type rollResponse struct {
		Number int
		Value  int
		Result int
	}

	roll_result, err := dice.DiceRoll(1, 20)
	if err != nil {
		fmt.Printf("Error in rolling dice:\t%v\n", err)
		return
	}

	response_obj := rollResponse{
		Number: 1,
		Value:  20,
		Result: roll_result,
	}

	response_data, err := json.Marshal(response_obj)
	if err != nil {
		fmt.Printf("Error in marshalling response_object:\t%v\n", err)
		return
	}

	w.Header().Set("Content-Type", "application/json") // Not really sure why this is required
	w.Write(response_data)

	//http.ServeFile(w, req, "web/text.txt")

}

func main() {
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
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/roll", rollHandler)
	http.HandleFunc("/favicon.ico", iconHandler)

	http.HandleFunc("/roll_ajax", rollAjaxHandler)
	http.HandleFunc("/text.txt", testTextHandler)

	fmt.Print("Listening on port 8080.\n\n")
	http.ListenAndServe(":8080", nil)
}
