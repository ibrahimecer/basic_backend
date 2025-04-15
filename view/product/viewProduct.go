package product

import (
	"basicbackend/data"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func GetAllProducts() ([]data.Product, error) {
	response, err := http.Get("http://localhost:3000/products")
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var products []data.Product

	err = json.Unmarshal(bodyBytes, &products)
	if err != nil {
		return nil, err
	}
	responseCtg, err := http.Get("http://localhost:3000/categories")
	if err != nil {
		return nil, err
	}

	defer responseCtg.Body.Close()
	bodyBytesCtg, err := ioutil.ReadAll(responseCtg.Body)
	if err != nil {
		return nil, err
	}

	var categories []data.Category
	err = json.Unmarshal(bodyBytesCtg, &categories)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(products); i++ {
		for j := 0; j < len(categories); j++ {
			if products[i].CategoryId == categories[j].Id {
				products[i].CategoryId = categories[j].CategoryName
			}
		}
	}
	return products, nil
}
