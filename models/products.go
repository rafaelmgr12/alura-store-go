package models

import "aluraStoreGo/db"

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func SearchAllProducts() []Product {
	db := db.ConnectDatabase()

	selectAll, err := db.Query("SELECT * FROM products")
	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for selectAll.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = selectAll.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}
		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity

		products = append(products, p)
	}
	defer db.Close()
	return products

}

func CreateNewProduct(name, description string, price float64, quantity int) {
	db := db.ConnectDatabase()

	insertData, err := db.Prepare("insert into products(name,description,price,quantity) values($1,$2,$3,$4)")
	if err != nil {
		panic(err.Error())
	}
	insertData.Exec(name, description, price, quantity)
	defer db.Close()

}

func DeleteProduct(id string) {
	db := db.ConnectDatabase()

	deleteData, err := db.Prepare("delete from products where id=$1")
	if err != nil {
		panic(err.Error())
	}
	deleteData.Exec(id)
	defer db.Close()
}
