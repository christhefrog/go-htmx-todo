package main

import (
	"log"
	"net/http"
	"strconv"
	"text/template"
)

type TodoItem struct {
	Id   int
	Text string
	Date string
}

func remove(slice []TodoItem, s int) []TodoItem {
	return append(slice[:s], slice[s+1:]...)
}

func main() {
	items := map[int]TodoItem{}
	lastItemID := -1

	index := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("www/index.html"))
		tmpl.Execute(w, nil)
	}

	listItems := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("www/listItems.html"))

		var itemsSlice []TodoItem
		for _, value := range items {
			itemsSlice = append(itemsSlice, value)
		}
		context := map[string][]TodoItem{
			"Items": itemsSlice,
		}

		tmpl.Execute(w, context)
	}

	addItem := func(w http.ResponseWriter, r *http.Request) {
		text := r.PostFormValue("text")
		date := "27.04.2024"

		lastItemID += 1
		items[lastItemID] = TodoItem{Id: lastItemID, Text: text, Date: date}

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

	log.Fatal(http.ListenAndServe(":8000", nil))
}
