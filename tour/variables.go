package tour

import "fmt"

func declareVariables() {
	var c, python, java bool
	fmt.Println(c, python, java)
}

func declareVariablesWithInitValues() {
	//If an init value is present then type can be omitted
	var a, b, c = true, false, "No"
	fmt.Println(a, b, c)
}

func shortVariableDeclaration() {
	var i, j int = 1, 2

	//this short assignment can only be done inside functions
	c, python, java := true, false, "Yes"
	fmt.Println(i, j, c, python, java)
}

// var i, j int = 1, 2
// var c, python, java = true, false, "Yes"

// RunVariables var statement declares a list of variables
func RunVariables() {
	declareVariables()
	declareVariablesWithInitValues()
	shortVariableDeclaration()
}
