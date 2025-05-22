package main

import "fmt"

const appName = "Budgify v1.0"

/*const currency = "$"

var totalIncome float64
var totalExpenses float64
var balance float64 */

func main() {
	var separator = "========================"

	fmt.Println(separator)
	fmt.Println(appName)
	fmt.Println(separator)

	fmt.Println("1. Add income")
	fmt.Println("2. Add expense")
	fmt.Println("3. Show summary")
	fmt.Println("4. Exit")

	var choice int
	fmt.Scan(&choice)

	fmt.Println("Your choice: ", choice)

}
