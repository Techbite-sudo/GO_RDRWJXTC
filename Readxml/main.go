package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
)
type Notes struct{
	To string `xml:"to"`
	From string `xml:"from"`
	Heading string `xml:"heading"`
	Body string `xml:"body"`
}
func main(){
	//fetch the xml data
	data,err:= ioutil.ReadFile("notes.xml")
	if err != nil{
		log.Println("No such file found")
	}
	//create an instance of the struct
	note := new(Notes)
	//unmarshal xml data into struct
	xml.Unmarshal([]byte(data),note)
	fmt.Println(*note)
	fmt.Println(note.To)
	fmt.Println(note.From)
	fmt.Println(note.Heading)
	fmt.Println(note.Body)	
}