package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

type TodoItem struct {
	Id   int
	Text string
	Date string
}

func main() {
	items := map[int]TodoItem{}
	lastInsertedItemID := 0

	index := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("www/index.html"))
		tmpl.Execute(w, nil)
	}

	listItems := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("www/listItems.html"))

		var itemsSlice []TodoItem
		for i := 0; i < lastInsertedItemID; i++ {
			value, ok := items[i]
			if ok {
				itemsSlice = append(itemsSlice, value)
			}

		}
		context := map[string][]TodoItem{
			"Items": itemsSlice,
		}

		tmpl.Execute(w, context)
	}

	addItem := func(w http.ResponseWriter, r *http.Request) {
		text := r.PostFormValue("text")
		date := time.Now().Format("02.01.2006")

		if text != "" {
			items[lastInsertedItemID] = TodoItem{Id: lastInsertedItemID, Text: text, Date: date}
			lastInsertedItemID += 1
		}

		listItems(w, r)
	}

	removeItem := func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(r.PostFormValue("itemid"))

		delete(items, id)

		listItems(w, r)
	}

	http.HandleFunc("/", index)
	http.HandleFunc("/listItems", listItems)
	http.HandleFunc("/addItem", addItem)
	http.HandleFunc("/removeItem", removeItem)

	fmt.Println("Listening...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
