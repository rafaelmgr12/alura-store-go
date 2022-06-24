package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"aluraStoreGo/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.SearchAllProducts()
	temp.ExecuteTemplate(w, "Index", products)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		priceToFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Error in price conversion")
		}
		quantityToInt, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Error in quantity conversion")
		}

		models.CreateNewProduct(name, description, priceToFloat, quantityToInt)

	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	models.DeleteProduct(id)
	http.Redirect(w, r, "/", 301)
}
