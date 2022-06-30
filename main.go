package main

import (
	"fmt"
	"math/rand"
	"time"
	"unicode"
)

func main() {
	passwordGenerator() // Calling a function
}

func passwordGenerator() {
	lowLet := "qwertyuiopasdfghjklzxcvbnm" // Creating a string of lower-case letters. Made an appropriate name for it.
	upLet := "QWERTYUIOPASDFGHJKLZXCVBNM"  // Creating a string of upper-case letters.
	nums := "0123456789"                   // Creating a string of numbers.
	symbols := "!?@#$^&*()_-[]"            // Creating a string of special symbols.
	password := ""                         // Creating an empty string to fill in later.

	var passwordLength int

	rand.Seed(time.Now().UnixNano()) // To initialise the pseudo random number generator seed.

	fmt.Print("What is length of your password?\n")
	fmt.Scan(&passwordLength)

	if passwordLength <= 0 { // Checking for correct input by user.
		fmt.Print("Incorrect input. Try again!\n")
		passwordGenerator()
	}

	for i := 0; i < passwordLength; i++ { // Creating cycle for fill user password.
		randLowLet := lowLet[rand.Intn(len(lowLet))]                                                // Initialisating a variable and fill it with random lower-case letters.
		randUpLet := upLet[rand.Intn(len(upLet))]                                                   // Initialisating a variable and fill it with random upper-case letters.
		randNums := nums[rand.Intn(len(nums))]                                                      // Initialisating a variable and fill it with random numbers.
		randSymbols := symbols[rand.Intn(len(symbols))]                                             // Initialisating a variable and fill it with random special symbols.
		password += string(randLowLet) + string(randUpLet) + string(randNums) + string(randSymbols) // Initialisating a variable and fill it with
	} // the result of the previously initialised variables.
	correctPassword := password[:passwordLength]                        // Cut off the length user want and initialise the variable with this length.
	fmt.Printf("\t\tGenerated password: %s\n", string(correctPassword)) // Output
	fmt.Printf("\t\tSelected length: %d\n", len(correctPassword))       // Output
	fmt.Print(isValid(correctPassword))                                 // Output the result of function isValid().
	repeat()
}

func isValid(check string) string {
	var choice int

	fmt.Print("Do you want to check is your password safe?\n[0][1] <- ?\t")
	fmt.Scan(&choice)

	if choice == 0 { // If user don't want to check is his password safe,
		return "continuing...\n" // enter "0", function will stop and returns message "continuing..."
	} else if choice == 1 { // If user want to check is his password safe, enter "1"
		enterRune := []rune(check)                           // Creating slice of runes.
		if unicode.Is(unicode.Latin, enterRune[0]) == true { // With help of package "unicode" checking first letter for latin symbols.
			counter := 0                      // Initialisating a counter for user password.
			for i := 0; i < len(check); i++ { // Creating cycle with counter inside for check how many symbols in password
				counter++
			}
			if counter >= 5 { // If our counter, which we are created, have more than 4 symbols
				return "It's OK. Password is safe!\n" // program returns to user that password is safe.
			} else { // If less, returns to user that password is not safe.
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

	fmt.Print("Do you want to create one more password?\n[0][1] <- ?\t") // If user want to create one more password,
	fmt.Scan(&choice)                                                    // programm calls function passwordGenerator() again

	if choice == 0 { // If he doesn't want to create one more password,
		fmt.Print("Thanks for using my password generator!\n") // programm returns message "Thanks for using..."
	} else if choice == 1 {
		passwordGenerator()
	} else {
		panic("incorrect input...")
	}
}
