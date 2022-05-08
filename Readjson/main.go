package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)
type CatlogNodes struct{
	CatlogNodes []Catlog `json:"catlog_nodes"`
}
type Catlog struct{
	Product_id string `json:"product_id"`
	Product_name string `json:"product_name"`
	Quantity int `json:"quantity"`
}
func main(){
	//Read the json file
	jsonData,err:= ioutil.ReadFile("test.json")
	if err != nil{
		log.Println(err)
	}
	catlog1 := new(CatlogNodes)
	json.Unmarshal([]byte(jsonData),&catlog1)

	for _,v := range catlog1.CatlogNodes{
		fmt.Println("Product Id:",v.Product_id)
		fmt.Println("Product Name:",v.Product_name)
		fmt.Println("Product Id:",v.Quantity)
		println("----------------------------------")

	}
	// for i := 0; i < len(catlog1.CatlogNodes); i++ {
	// 	// fmt.Println("Product Id: ", catlog1.CatlogNodes[i].Product_id)
	// 	// fmt.Println("Product Name: ", catlog1.CatlogNodes[i].Product_name)
	// 	// fmt.Println("Quantity: ", catlog1.CatlogNodes[i].Quantity)
	// 	fmt.Println(catlog1.CatlogNodes[i])
	// }

	
}