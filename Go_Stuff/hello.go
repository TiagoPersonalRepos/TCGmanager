package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func http_request() {
	// url := "https://jsonplaceholder.typicode.com/posts/1"

	url := "https://db.ygoprodeck.com/api/v7/cardinfo.php?format=tcg"

	resp, err := http.Get(url)

	if err != nil {
		log.Fatalln(err)
	}

	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	   log.Fatalln(err)
	}

 	//Convert the body to type string
	sb := string(body)
	log.Printf(sb)
}

func main() {
	a := app.New()
	w := a.NewWindow("Hello World")

	w.SetContent(widget.NewLabel("Hello World!"))
	w.ShowAndRun()
}
