package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)
type Salary struct{
	Basic float64
	HRA float64
	TA float64
}
type Employee struct{
	FirstName string
	LastName string
	Email string 
	Age int
	MonthlySalary []Salary
}
func main(){
	Emp := Employee{
		FirstName: "Boniface",
		LastName: "Mwema",
		Email: "bonfii@gmail.com",
		Age: 21,
		MonthlySalary: []Salary{
			Salary{
				Basic: 15000.00,
                HRA:   5000.00,
                TA:    2000.00,
			},
			Salary{
				Basic: 16000.00,
                HRA:   5000.00,
                TA:    2100.00,
			},
			Salary{
				Basic: 17000.00,
                HRA:   5000.00,
                TA:    2200.00,
			},
		},

	}
	file,err := json.MarshalIndent(&Emp,"","")
	if err != nil{
		log.Println("Did not Marshal Emp data into File")
	}
	ioutil.WriteFile("test.json",file,0644)


}