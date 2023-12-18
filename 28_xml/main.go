package main

import (
	"encoding/xml"
	"fmt"
)

type Plant struct {
	XMLName xml.Name `xml:"plant"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v",
		p.Id, p.Name, p.Origin)
}
func Example1() {
	//Xml Oluşturma
	coffee := &Plant{Id: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}

	out, _ := xml.MarshalIndent(coffee, " ", "  ")
	/*
	 <plant id="27">
	   <name>Coffee</name>
	   <origin>Ethiopia</origin>
	   <origin>Brazil</origin>
	 </plant>
	*/
	fmt.Println(string(out))

	/*
		<?xml version="1.0" encoding="UTF-8"?> //xml.Header
		 <plant id="27">
		   <name>Coffee</name>
		   <origin>Ethiopia</origin>
		   <origin>Brazil</origin>
		 </plant>
	*/
	fmt.Println(xml.Header + string(out))

	//Xml i objeye cevirme
	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	//Plant id=27, name=Coffee, origin=[Ethiopia Brazil]
	fmt.Println(p)

	tomato := &Plant{Id: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}

	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants  []*Plant `xml:"parent>child>plant"`
	}

	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, tomato}

	out, _ = xml.MarshalIndent(nesting, " ", "  ")
	fmt.Println(string(out))

}

func main() {
	out := []byte{}
	type User struct {
		XMLName   xml.Name `xml:"User"`
		FirstName string   `xml:"FirstName"`
		LastName  string   `xml:"LastName"`
	}
	type UserList struct {
		XMLName xml.Name `xml:"UserList"`
		// Users   []User   `xml:"parent>child>user"` //Burdaki parent chil user sadece nesting yapı oluşturmak için kullanılıyor
		Users []User `xml:"user"`
	}

	userList := UserList{
		Users: []User{
			{FirstName: "User1", LastName: "Lastname1"},
			{FirstName: "User2", LastName: "Lastname2"},
			{FirstName: "User3", LastName: "Lastname3"},
			{FirstName: "User4", LastName: "Lastname4"},
			{FirstName: "User5", LastName: "Lastname5"},
		}}

	out, _ = xml.MarshalIndent(userList, " ", "  ")
	fmt.Println(string(out))
}
