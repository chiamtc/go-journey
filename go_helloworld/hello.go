package main //package declacraion

//package = a project or workspace
// a collection of golang file

/*
main.go
helper.go
utils.go
- all have to have "package main" at the very first line in that file

types of packages
1. executables = .exe if it has "package main" and uses "go build main" and in that project it has to have "func main()"
2. reusable  = like library if it has "package xxx" other than main and uses "go build xxx" then it will not have executable file


1. package main
2. package calculator or package utils

*/

import "fmt" //list of imports

/*
standard packages , source: golang.org/pkg
*/

//declare functions and write actual source codes
func main() {
	fmt.Println("hello, world")
	fmt.Printf("hello word from printF")
}

/*
go run = go run xx.go = executes the go file
go build = go build = compile all the go files in the dir
go fmt =  format all the code in each file
go install = compiles and install a dependency
go get = donwloads raw source code of someone else's package
go test = test

*/
