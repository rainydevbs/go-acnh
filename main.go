package main

import (
	"fmt"           // for formatting our text
	"html/template" // a library that allows us to interact with our html file
	"net/http"      // to access the core go http functionality
	"time"          // a library for working with date and time
)

// holds information to be displayed in our HTML file
type Welcome struct {
	Name string
	Time string
}

// entrypoint
func main() {
	welcome := Welcome{"Anonymous", time.Now().Format(time.Stamp)}

	// relative path
	// template.Must() handles any errors and halts if there are fatal errors
	templates := template.Must(template.ParseFiles("templates/welcome-template.html"))

	// create a handle that looks in the static directory
	// uses the "/static/" as a url that html can refer to find files (.css, .png)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	// looks in relative "static" directory first using http.FileServer()
	// matches it to a url of our choice as shown in http.Handle("/static/")
	// use this when referencing files
	// <link rel="stylesheet"  href="/static/stylesheet/...">

	// takes in the pattern "/" (will always execute)
	// passed in handler function handles when pattern matches
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {

		// URL query e.g ?name=Martin
		// declare and initialize name
		// the condition is (name != "")
		if name := request.FormValue("name"); name != "" {
			welcome.Name = name
		}

		// pass the welcome struct to the welcome-template.html file
		// also shows an error if it fails
		if err := templates.ExecuteTemplate(writer, "welcome-template.html", welcome); err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
	})

	// wrap serving the website in fmt for error message displays
	fmt.Println("Listening")
	fmt.Println(http.ListenAndServe(":19751", nil))
}
