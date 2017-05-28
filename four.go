package main

import "fmt"

func my_name_and_roll_no() (string,int)  {
	return "Salim Mahammad Almas",3
}

func main() {
	name , rollno := my_name_and_roll_no()
	fmt.Println("My name ",name)
	fmt.Println("My roll no " , rollno )
}
