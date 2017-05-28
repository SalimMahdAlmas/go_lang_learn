package main

import "fmt"

type Person struct {
	name string
	age int
	student bool
}
func main() {

	evil := &Person{"Salim Mahammad Almas",15,true}
	dump_person(evil)

	// just to learn . not true tho

	evil.student = false
	evil.name = "chuvil"
	evil.age = 69
	dump_person(evil)

	switchStudent(evil)
	dump_person(evil)
}
func dump_person(person *Person) {
	fmt.Println()
	fmt.Println("Name    := ",person.name)
	fmt.Println("Age     := ",person.age)
	fmt.Println("Student := ",person.student)
}
// pointer stuff
func switchStudent(person *Person) {
	person.student = !person.student
}

func newPerson(name string, age int , student bool) Person {
	return Person{
		name,age,student,
	}
}