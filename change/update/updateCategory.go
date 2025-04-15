package update

import (
	"basicbackend/data"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func UpdateCategory(id string, category data.Category) error {
	jsonCategory, err := json.Marshal(category)
	if err != nil {
		return err
	}
	url := fmt.Sprintf("http://localhost:3000/categories/%s", id)
	response, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonCategory))
	if err != nil {
		fmt.Println("HTTP Request hatası:", err)
		return err
	}
	response.Header.Set("Content-Type", "application/json;charset=utf-8")
	client := &http.Client{}
	req, err := client.Do(response)
	if err != nil {
		fmt.Println("HTTP Response hatası:", err)
		return err
	}
	defer req.Body.Close()
	if req.StatusCode == http.StatusOK {
		fmt.Println("Kategori başarılı bir şekilde güncellendi.")
	} else {
		fmt.Println("Güncelleme başarısız : ", req.StatusCode)
	}
	return nil
}
