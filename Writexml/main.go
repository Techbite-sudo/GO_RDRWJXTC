package main

import (
	"encoding/xml"
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
	//creating an instance of Notes and assigning values 
	note := new(Notes)
	note.To = "Boniface"
	note.From = "Mwema"
	note.Heading = "Software engineering"
	note.Body = "This is the most interesting course anyone can ever take "
	//Marshalling the struct into the xml file 
	xmlFile,err := xml.MarshalIndent(note, "", "\n")
	if err != nil{
		log.Println("Did not Marshal")
	}
	//Writing the marshalled data into xml file called notes.xml
	error:= ioutil.WriteFile("notes.xml",xmlFile,0644)
	if error != nil{
		log.Println("Failed to write into notes.xml")
	}

}