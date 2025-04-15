package category

import (
	"basicbackend/data"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func GetAllCategories() ([]data.Category, error) {
	response, err := http.Get("http://localhost:3000/categories")
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var categories []data.Category
	err = json.Unmarshal(bodyBytes, &categories)
	if err != nil {
		return nil, err
	}
	return categories, nil
}
