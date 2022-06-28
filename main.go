package main

import (
	"fmt"
	"math/rand"
	"time"
	"unicode"
)

func main() {
	passwordGenerator() //
}

func passwordGenerator() {
	lowLet := "qwertyuiopasdfghjklzxcvbnm"
	upLet := "QWERTYUIOPASDFGHJKLZXCVBNM"
	nums := "0123456789"
	symbols := "!?@#$^&*()_-[]"
	password := ""

	var passwordLength int

	rand.Seed(time.Now().UnixNano())

	fmt.Print("What is length of your password?\n")
	fmt.Scan(&passwordLength)

	if passwordLength <= 0 {
		fmt.Print("Incorrect input. Try again!\n")
		passwordGenerator()
	}

	for i := 0; i < passwordLength; i++ {
		randLowLet := lowLet[rand.Intn(len(lowLet))]
		randUpLet := upLet[rand.Intn(len(upLet))]
		randNums := nums[rand.Intn(len(nums))]
		randSymbols := symbols[rand.Intn(len(symbols))]
		password += string(randLowLet) + string(randUpLet) + string(randNums) + string(randSymbols)
	}
	correctPassword := password[:passwordLength]
	fmt.Printf("\t\tGenerated password: %s\n", string(correctPassword))
	fmt.Printf("\t\tSelected length: %d\n", len(correctPassword))
	fmt.Print(isValid(correctPassword))
	repeat()
}

func isValid(check string) string {
	var choice int

	fmt.Print("Do you want to check is your password safe?\n[0][1] <- ?\t")
	fmt.Scan(&choice)

	if choice == 0 {
		return "continuing...\n"
	} else if choice == 1 {
		enterRune := []rune(check)
		if unicode.Is(unicode.Latin, enterRune[0]) == true {
			counter := 0
			for i := 0; i < len(check); i++ {
				counter++
			}
			if counter >= 5 {
				return "It's OK. Password is safe!\n"
			} else {
				return "Password is not safe.\n"
			}
		} else {
			return "Password is not safe.\n"
		}
	} else {
		return "continuing...\n"
	}
}

func repeat() {
	var choice int

	fmt.Print("Do you want to create one more password?\n[0][1] <- ?\t")
	fmt.Scan(&choice)

	if choice == 0 {
		fmt.Print("Thanks for using my password generator!\n")
	} else if choice == 1 {
		passwordGenerator()
	} else {
		panic("incorrect input...")
	}
}
