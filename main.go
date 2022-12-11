package main

import (
	"fmt"
	"strconv"
)

func main() {

	var operator int
	var num1 int
	var num2 int
	var result int
	history := make([][]string , 0)

	for {
		fmt.Println("SIMPLE CALCUCALTOR")
		fmt.Println("======================")
		fmt.Println("Menu options :")
		fmt.Println("1. Addition")
		fmt.Println("2. Subtraction")
		fmt.Println("3. Division")
		fmt.Println("4. Multiplication")
		fmt.Println("5. History")
		fmt.Println("-")
		fmt.Println("9. Exit")
		fmt.Println("======================")

		fmt.Print("Please choose an option [1, 2, 3, 4, 5, 9] : ")
		fmt.Scanln(&operator)

		if operator == 9 {
			fmt.Println("Exiting ....")
			fmt.Println("THANKS, HAVE A NICE DAY!")
			fmt.Println("======================")
			break
		}

		var operatorSymbol string

		switch operator {
		case 1:
			operatorSymbol = "+"
		case 2:
			operatorSymbol = "-"
		case 3:
			operatorSymbol = "÷"
		case 4:
			operatorSymbol = "x"
		}

		switch operator {
		case 1:
			fmt.Print("Enter the first number : ")
			fmt.Scanln(&num1)
			fmt.Print("Enter the second number : ")
			fmt.Scanln(&num2)
			result = num1 + num2
			fmt.Printf("RESULT : %v %v %v = %v\n", num1, operatorSymbol, num2, result)
		case 2:
			fmt.Print("Enter the first number : ")
			fmt.Scanln(&num1)
			fmt.Print("Enter the second number : ")
			fmt.Scanln(&num2)
			result = num1 - num2
			fmt.Printf("RESULT : %v %v %v = %v\n", num1, operatorSymbol, num2, result)
		case 3:
			fmt.Print("Enter the first number : ")
			fmt.Scanln(&num1)
			fmt.Print("Enter the second number : ")
			fmt.Scanln(&num2)
			if num2 == 0 {
				result = 0
			} else {
				result = num1 / num2
			}
			fmt.Printf("RESULT : %v %v %v = %v\n", num1, operatorSymbol, num2, result)
		case 4:
			fmt.Print("Enter the first number : ")
			fmt.Scanln(&num1)
			fmt.Print("Enter the second number : ")
			fmt.Scanln(&num2)
			result = num1 * num2
			fmt.Printf("RESULT : %v %v %v = %v\n", num1, operatorSymbol, num2, result)
		case 5:
			fmt.Println("HISTORY :")
			for _, v := range history {
			fmt.Println(v[0], v[1], v[2], v[3], v[4])
			}
		}

		convNum1 := strconv.Itoa(num1)
		convNum2 := strconv.Itoa(num2)
		convResult := strconv.Itoa(result)

		listHistory := [5]string{convNum1, operatorSymbol, convNum2, "=", convResult}
		history = append(history, listHistory[:])

		fmt.Println("======================")
		fmt.Println()
		fmt.Println("Starting again ....")
		fmt.Println()
		fmt.Println("======================")
	}
}
