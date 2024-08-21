package main

import (
	"fmt"

	"example.com/bank-calculator/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {

	accountBalance, err := fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("------------")
		panic("Sorry, cannot continue")
	}

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Reach out to us 24/7. Phone:", randomdata.PhoneNumber(), "\n")

	for {

		presentOptions()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println("Current Balance:", accountBalance, "\n")
		} else if choice == 2 {
			fmt.Print("Deposit Amount: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid Amount! Must be greater than 0\n")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Balance Updated! Current Balance:", accountBalance, "\n")
		} else if choice == 3 {
			fmt.Print("Withdraw Amount: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid Amount! Must be greater than 0")
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Println("Invalid Amount! You can't withdraw more than you have\n")
				continue
			}

			accountBalance -= withdrawAmount
			fmt.Println("Balance Updated! Current Balance:", accountBalance, "\n")
		} else {
			fmt.Println("Goodbye!")
			break
		}

	}

	fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
	fmt.Println("Thanks for choosing our bank")

}
