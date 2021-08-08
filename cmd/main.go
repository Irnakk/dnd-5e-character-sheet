package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

func iconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../web/content/icon.png")
}

func printRequestInfo(req *http.Request) {
	fmt.Println("----------------")
	fmt.Printf("req.Method:\t%v\n", req.Method)
	fmt.Printf("req.RequestURI:\t%v\n", req.RequestURI)
	fmt.Println("----------------")
}

func rootHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("In rootHandler()")
	defer fmt.Print("Out of rootHandler\n\n****************\n\n")
	printRequestInfo(req)

	t, _ := template.ParseFiles("../web/index.html")
	t.Execute(w, nil)
}

func rollHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("In rollHandler()")
	defer fmt.Print("Out of rollHandler\n\n****************\n\n")
	printRequestInfo(req)

	t, _ := template.ParseFiles("../web/roll.html")
	t.Execute(w, nil)

	switch req.Method {
	case "GET":
		//http.ServeFile(w, req, "../web/roll")
	case "POST":
		if err := req.ParseForm(); err != nil {
			fmt.Printf("Error in parsing form in rollHandler:\t%v", err)
			return
		}

		fmt.Printf("Successful POST; req.PostForm:\t%v\n", req.PostForm)
		info := req.FormValue("info")
		fmt.Printf("Received <info>:\t%v\n", info)
	}
}

func main() {
	fmt.Printf("Starting func main() at:\t%v\n", time.Now())

	defer fmt.Printf("Finishing func main() at:\t%v\n", time.Now())

	fmt.Println("Binding handler functions...")
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/roll", rollHandler)
	http.HandleFunc("/favicon.ico", iconHandler)

	fmt.Print("Listening on port 8080.\n\n")
	http.ListenAndServe(":8080", nil)
}
