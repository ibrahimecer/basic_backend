package add

import (
	"basicbackend/data"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func AddCategory(category data.Category) {
	jsonCategory, err := json.Marshal(category)

	response, err := http.Post("http://localhost:3000/categories", "application/json;charset=utf-8", bytes.NewBuffer(jsonCategory))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()
	fmt.Println("Ürün başarılı bir şekilde kaydedildi")
}
