package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct{
    FirstName string `json:"firstname"` 
    LastName string `json:"lastname"`
    City string `json:"city"`
}
func main() {
	json_string := `{
        "firstname":"Boniface",
        "lastname":"Mwema",
        "city":"Nairobi"
    }`
    //json into struct    
    emp1 := new(Employee)
    json.Unmarshal([]byte(json_string),emp1)
    fmt.Println(*emp1)
    //normal    
    emp2 := new(Employee)
    emp2.FirstName = "Ladesha"
    emp2.LastName = "Mwanza"
    emp2.City = "Eldoret"
    fmt.Println(*emp2)
    //struct into json
    emp3 := new(Employee)
    emp3.FirstName = "Brian"
    emp3.LastName = "Yegon"
    emp3.City = "Kitale"
    jsonStr,_:= json.Marshal(*emp3)
    fmt.Printf("%s\n",jsonStr)
}