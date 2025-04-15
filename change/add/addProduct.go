package add

import (
	"basicbackend/data"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func AddProduct(product data.Product) {
	jsonProduct, err := json.Marshal(product)

	response, err := http.Post("http://localhost:3000/products", "application/json;charset=utf-8", bytes.NewBuffer(jsonProduct))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()
	fmt.Println("Ürün başarılı bir şekilde kaydedildi")
}
