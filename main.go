package main

import (
	"basicbackend/change/remove"
	"basicbackend/change/update"
	"basicbackend/data"
	"basicbackend/view/category"
	"basicbackend/view/product"
	"fmt"
)

func main() {
	update.UpdateCategory("1", data.Category{"1", "Teknoloji"})
	prd, err1 := product.GetAllProducts()
	if err1 != nil {
		fmt.Println("GetAllProducts Error: ", err1)
	}
	fmt.Println(prd)
	// add.AddCategory(data.Category{"3", "Home"})
	ctg, err2 := category.GetAllCategories()
	if err2 != nil {
		fmt.Println("GetAllCategories Error: ", err2)
	}
	fmt.Println(ctg)
	err4 := remove.DeleteCategory("3")
	if err4 != nil {
		fmt.Println("DeleteCategory: ", err4)
	}
}
