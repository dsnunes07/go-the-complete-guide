package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/pallinder/go-randomdata"
)

const accountBalanceFileName = "balance.txt"
const defaultInitialBalance = 0

func main() {
	accountBalance, err := fileops.GetFloatFromFile(accountBalanceFileName, defaultInitialBalance)
	if err != nil {
		fmt.Println("------------")
		fmt.Println("Error:", err)
		fmt.Println("------------")
	}

	clientName := randomdata.FirstName(randomdata.RandomGender)

	fmt.Printf("Hi, %s, welcome to Go Bank!\n", clientName)

	// go has only for loop
	for {
		choice := requestUserInput()
		switch choice {
		case 1:
			fmt.Println("Your balance is: ", accountBalance)
		case 2:
			fmt.Print("Your deposit: ")
			var amountToDeposit float64
			fmt.Scan(&amountToDeposit)
			if amountToDeposit > 0 {
				accountBalance += amountToDeposit
				fmt.Println("Balance updated! New amount: ", accountBalance)
				fileops.WriteFloatToFile(accountBalance, accountBalanceFileName)
			} else {
				fmt.Printf("Invalid amount: %.2f\n", amountToDeposit)
			}
		case 3:
			var withdrawAmount float64
			fmt.Println("Enter withdraw value:")
			fmt.Scan(&withdrawAmount)
			if withdrawAmount > 0 && withdrawAmount <= accountBalance {
				accountBalance -= withdrawAmount
				fileops.WriteFloatToFile(accountBalance, accountBalanceFileName)
			} else {
				fmt.Println("Cannot withdraw value ", withdrawAmount)
			}
		case 4:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}
