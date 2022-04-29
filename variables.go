package main

import "fmt"

func declare_variables() {
	var c, python, java bool
	fmt.Println(c, python, java)
}

func declare_variables_with_init_values() {
	//If an init value is present then type can be omitted
	var a, b, c = true, false, "No"
	fmt.Println(a, b, c)
}

func short_variable_declaration() {
	var i, j int = 1, 2

	//this short assignment can only be done inside functions
	c, python, java := true, false, "Yes"
	fmt.Println(i, j, c, python, java)
}


// var i, j int = 1, 2
// var c, python, java = true, false, "Yes"

//var statement declares a list of variables
func main() {
	declare_variables()
	declare_variables_with_init_values()
	short_variable_declaration()
}