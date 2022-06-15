package main

import "fmt"

type contact struct {
	email    string
	postCode int
}

type person struct {
	firstName string
	lastName  string
	//contact contact
	contact
}

func main() {
	jim := person{
		firstName: "jim",
		lastName:  "moriaty",
		contact: contact{
			email:    "jim@gmail.com",
			postCode: 123456,
		},
	}

	//Points to address
	//pointerToJim := &jim

	jim.updateFirstName("Jimmy")
	jim.print()
}

func (pointerToPerson *person) updateFirstName(firstName string) {
	(*pointerToPerson).firstName = firstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

/* func main() {
	alex := person{"Alex", "Anderson"}
	fmt.Println(alex)
	irshad := person{lastName: "Ahmed", firstName: "Irshad"}
	fmt.Println(irshad)

	var john person
	fmt.Println(john)

	john.firstName = "john"
	john.lastName = "smith"

	fmt.Printf("%+v", john)
} */
