package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}
type contactInfo struct {
	email   string
	zipCode int32
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Owen",
		contactInfo: contactInfo{
			email:   "Jim@aa.com",
			zipCode: 90501,
		},
	}
	jim.print()
	pJim := &jim

	pJim.update("Alex")
	jim.print()

}

func (p *person) update(firstName string) {
	(*p).firstName = firstName
}
func (p *person) print() {
	fmt.Printf("+%v\n", *p)
}
