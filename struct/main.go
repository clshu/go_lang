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
	names := []string{"Hello", "You're", "the", "best."}

	jim.print()

	jim.update("Alex")
	jim.print()

	updateStr(names)
	fmt.Println(names)
}

func updateStr(s []string) {
	s[0] = "Puppy"
}
func (p *person) update(firstName string) {
	(*p).firstName = firstName
}
func (p *person) print() {
	fmt.Printf("+%v\n", *p)
}
