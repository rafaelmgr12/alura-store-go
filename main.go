package main

import (
	"html/template"
	"net/http"
)

type Product struct {
	Name        string
	Description string
	Price       float64
	Quantity    int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	products := []Product{
		{Name: "Gopher", Description: "A gopher is a small, furry, and unique animal.", Price: 1.99, Quantity: 1},
		{Name: "T-Shirt", Description: "A t-shirt is a style of fabric, usually textured, made for human use.", Price: 9.99, Quantity: 1},
		{Name: "Socks", Description: "Socks are a type of item of clothing worn on the feet.", Price: 5.99, Quantity: 1},
		{Name: "Hat", Description: "A hat is a piece of head covering worn by humans.", Price: 12.99, Quantity: 1},
	}
	temp.ExecuteTemplate(w, "Index", products)
}
