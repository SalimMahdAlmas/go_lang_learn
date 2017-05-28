package main

import "fmt"

func get_my_name() string {
	return "Salim Mahammad Almas"
}
func get_leet_number() int {
	return 1337
}

func main() {
	var leet int
	leet = 1337
	fmt.Printf("Leet : %d\n",leet)

	var my_name string
	my_name = "Salim Mahammad Almas"
	fmt.Printf("My name : %s\n",my_name)


	leet_ := 1337
	fmt.Printf("Leet : %d\n",leet_)

	my_name_ := "Salim Mahammad Almas"
	fmt.Printf("My name : %s\n",my_name_)


	fmt.Printf("leet : %d\n",get_leet_number())

	fmt.Printf("My name : %s\n",get_my_name())

}
