package model

import (
	"errors"
	// "log"
)

var lastProductId int = 3

type Product struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Category string `json:"category"`
	Status   string `json:"status"`
}

var dataBaseProducts = []*Product{
	{1, "bread", 3500, "food", "active"},
	{2, "mineral water", 3000, "food", "active"},
	{3, "chocolate", 3000, "Minuman", "active"},
}

func Index() ([]*Product, error) {
	return dataBaseProducts, nil
}

func Show(id int) (*Product, error) {
	products, err := Index()
	if err != nil {
		return nil, errors.New("something went wrong")
	}

	for index, product := range products {
		if product.Id == id {
			return product, nil
		} else if index+1 == len(products) {
			return nil, errors.New("product not found")
		}
	}

	return nil, errors.New("product not found")
}

func Create(product *Product) (*Product, error) {
	product.Id = currentProductId()
	product.Status = "active"

	dataBaseProducts = append(dataBaseProducts, product)

	return product, nil
}

func Update(data *Product) (*Product, error) {
	products, err := Index()
	if err != nil {
		return nil, errors.New("something went wrong")
	}

	for _, product := range products {
		if product.Id == data.Id {
			product.Name = data.Name
			product.Price = data.Price
			product.Category = data.Category
			product.Status = data.Status

			return product, nil
		}
	}

	return nil, errors.New("product not found")
}

func Delete(id int) (bool, error) {
	for index, _ := range dataBaseProducts {
		if index == id {
			dataBaseProducts = append(dataBaseProducts[:index], dataBaseProducts[index+1:]...)
			return true, nil
		}
	}

	return false, errors.New("product not found")
}

func currentProductId() int {
	lastProductId++
	return lastProductId
}

