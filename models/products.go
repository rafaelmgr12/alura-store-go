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

func EditProduct(id string) Product {
	db := db.ConnectDatabase()

	dbProduct, err := db.Query("select * from products where id=$1", id)
	if err != nil {
		panic(err.Error())
	}
	updateProduct := Product{}

	for dbProduct.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = dbProduct.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}
		updateProduct.Id = id
		updateProduct.Name = name
		updateProduct.Description = description
		updateProduct.Price = price
		updateProduct.Quantity = quantity

	}
	defer db.Close()
	return updateProduct
}

func UpdateProduct(id int, name string, description string, price float64, quantity int) {
	db := db.ConnectDatabase()

	updateData, err := db.Prepare("update products set name=$1,description=$2,price=$3,quantity=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}
	updateData.Exec(name, description, price, quantity, id)
	defer db.Close()
}
