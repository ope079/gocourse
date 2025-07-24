package basics

import "fmt"

type Employee struct {
	Firstname string
	LastName  string
	Age       int
}

func main() {
	const MAXRETRIES = 5

	var employeeID = 1001

	fmt.Println("EmloyeeID: ", employeeID)
}
