package main

import "fmt"

const appName = "Budgify v1.0"

const currency = "$"

var totalIncome float64
var totalExpenses float64
var balance float64

func main() {
	var separator = "========================"
	var choice int

	for {
		fmt.Println(separator)
		fmt.Println(appName)
		fmt.Println(separator)

		fmt.Println("1. Add income")
		fmt.Println("2. Add expense")
		fmt.Println("3. Show summary")
		fmt.Println("4. Exit")
		fmt.Scan(&choice)

		if choice == 1 {
			addIncome()
		} else if choice == 2 {
			addExpense()
		} else if choice == 3 {
			showSummary()
		} else if choice == 4 {
			fmt.Println("Goodbye")
			break
		} else {
			fmt.Println("Invalid Choice. Try again.")
		}

	}

}

func getUserInput(prompt string) float64 {
	var amount float64
	fmt.Print(prompt)
	fmt.Scan(&amount)
	return amount
}

func addIncome() {
	amount := getUserInput("Enter income amount: ")
	totalIncome += amount
	balance = totalIncome - totalExpenses
	fmt.Println("Income added.")
}

func addExpense() {
	amount := getUserInput("Enter expense amount: ")
	totalExpenses += amount
	balance = totalIncome - totalExpenses
	fmt.Println("Expense added.")
}

func showSummary() {
	fmt.Println("========================")
	fmt.Println("SUMMARY")
	fmt.Println("========================")
	fmt.Printf("Total Income: %s%.2f\n", currency, totalIncome)
	fmt.Printf("Total Expenses: %s%.2f\n", currency, totalExpenses)
	fmt.Printf("Balance: %s%.2f\n", currency, balance)
}
