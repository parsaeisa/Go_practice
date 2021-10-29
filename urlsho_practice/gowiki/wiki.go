package main

import (
	"fmt"
	"log"
	"net/http"
	"handlers/handlers"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi , I love %s ! ", r.URL.Path[1:])
}

func main() {
	// p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page")}
	// p1.save()
	// p2, _ := loadPage("TestPage")
	// fmt.Println(string(p2.Body))
	//====================================================
	// http.HandleFunc("/", handler)
	// log.Fatal(http.ListenAndServe(":8080", nil))
	//====================================================
	http.HandleFunc("/view/", handlers.makeHandler(viewHandlers))
	http.HandleFunc("/edit/", handlers.makeHandler(editHandler))
	http.HandleFunc("/save/", handlers.makeHandler(saveHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
