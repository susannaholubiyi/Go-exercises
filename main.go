package main

import (
	"fmt"
)

func main() {
	name := "Susannah"
	fmt.Printf("Hello my name is %s\n", name)
	//addition()
	//gasMilleage()
	//creditLimitCalculator()
	//salesCommissionCalculator()
	largestNumber()

}
func addition() {
	result := 321325 * 424521
	fmt.Println(result)
}
func gasMilleage() {
	tripCounter := 0
	var totalMilesDriven, totalGallonsUsed int
	fmt.Println("Enter the miles driven and gallons used or -1 to end: ")
	for {
		var milesDriven, gallonUsed int
		fmt.Print("Miles driven (or -1 to finish): ")
		_, err := fmt.Scanf("%d", &milesDriven)

		if err != nil {
			fmt.Println("Error:", err)
			break
		}
		if milesDriven == -1 {
			break
		}
		fmt.Print("Gallons used: ")
		_, err = fmt.Scanf("%d", &gallonUsed)

		if err != nil {
			fmt.Println("Error:", err)
			break
		}
		milesPerGallon := float64(milesDriven) / float64(gallonUsed)
		fmt.Printf("The miles per gallon for this trip is: %.2f\n", milesPerGallon)

		totalMilesDriven += milesDriven
		totalGallonsUsed += gallonUsed
		tripCounter++
	}

	if tripCounter > 0 {
		totalMilesPerGallon := float64(totalMilesDriven) / float64(totalGallonsUsed)
		fmt.Printf("\n The total miles per gallon used for all the trips is: %.2f\n", totalMilesPerGallon)
	} else {
		fmt.Println("No Trips were recorded.")
	}

}
func creditLimitCalculator() {
	var accountNumber, beginningBalance, totalItemCharges, totalCreditApplied, creditLimit int
	fmt.Println("Enter your account Information to check if your credit limit is valid  ")
	fmt.Println("Enter your account number: ")
	_, err := fmt.Scanf("%d", &accountNumber)

	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Enter your balance at the beginning of the month: ")
	_, err2 := fmt.Scanf("%d", &beginningBalance)

	if err2 != nil {
		fmt.Println("Error:", err2)
	}
	fmt.Println("Enter your total item charges in a month: ")
	_, err3 := fmt.Scanf("%d", &totalItemCharges)

	if err3 != nil {
		fmt.Println("Error:", err3)
	}
	fmt.Println("Enter your total credits applied in a month: ")
	_, err4 := fmt.Scanf("%d", &totalCreditApplied)

	if err4 != nil {
		fmt.Println("Error:", err4)
	}
	fmt.Println("Enter your allowed credit limit: ")
	_, err5 := fmt.Scanf("%d", &creditLimit)

	if err5 != nil {
		fmt.Println("Error:", err5)
	}

	newBalance := (beginningBalance + totalItemCharges) - totalCreditApplied
	fmt.Printf("\nYour new balance is  %d", newBalance)
	if newBalance > creditLimit {
		fmt.Println("Credit limit exceeded")
	} else {
		fmt.Println("Credit limit has not been exceeded")
	}
}

func salesCommissionCalculator() {
	const fixedPrice = 200.00
	totalPriceOfItemSold := 0.00
	var priceOfItem float64
	fmt.Println("Enter the price of items you sold or -1 to end To calculate Your sales commission")
	for {

		fmt.Println("Price of item in $: ")
		_, err := fmt.Scanf("%f", &priceOfItem)
		if err != nil {
			fmt.Println("Error", err)
			break
		}

		if priceOfItem == -1 {
			break
		}
		totalPriceOfItemSold += priceOfItem
	}

	ninePercent := totalPriceOfItemSold * (0.09)

	finalCommission := fixedPrice + ninePercent

	fmt.Printf("\nYour commission is $%.2f", finalCommission)
}

func largestNumber() {
	largest := 0
	fmt.Println("Enter 10 numbers to determine the largest")
	for index := 0; index < 10; index++ {
		var number int
		fmt.Println("Enter number: ")
		_, err := fmt.Scanf("%d", &number)
		if err != nil {
			fmt.Println("Error: ", err)
			break
		}
		if index == 0 || number > largest {
			largest = number
		}
	}
	fmt.Printf("\nThe largest number you entered is %d", largest)
}
