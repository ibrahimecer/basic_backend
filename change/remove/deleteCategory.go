package remove

import (
	"fmt"
	"net/http"
)

func DeleteCategory(id string) error {
	url := fmt.Sprintf("http://localhost:3000/categories/%s", id)
	response, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		fmt.Println("HTTP Request hatası:", err)
		return err
	}
	client := &http.Client{}
	req, err := client.Do(response)
	if err != nil {
		fmt.Println("HTTP Response hatası:", err)
		return err
	}
	defer req.Body.Close()
	if req.StatusCode == http.StatusOK {
		fmt.Println("Ürün başarılı bir şekilde silincii.")
	} else {
		fmt.Println("Silinme başarısız : ", req.StatusCode)
	}
	return nil
}
