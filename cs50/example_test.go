// Copyright 2017 Arjun Rao
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cs50

import "fmt"
import "math"

// GetString() prompts a user with the given message for input, and returns their input as a string.
// It reads one line of input from standard input.
func ExampleGetString() {
	fmt.Println(GetString("Enter a string:"))
}

// GetFloat() can be used to conveniently get float64 value from the user via standard input.
// It reprompts the user if anything other than a float was entered.
func ExampleGetFloat() {
	var radius float64
	for {
		radius := GetInt("Enter radius of a circle:")
		if radius > 0 {
			break
		}
	}
	fmt.Printf("Area: %f", math.Pi*radius*radius)
}

// GetLong() can be used to conveniently get an int64 value from the user via standard input.
// It reprompts the user if anything other than an int64 was entered.
func ExampleGetLong() {
	var card int64

	for {
		card := GetLong("Get Credit Card Number:")
		if card > 0 {
			break
		}
	}
	fmt.Printf("Card Number: %d", card)

}

// GetInt() can be used to conveniently get an integer value input from the user via standard input.
// It reprompts the user if anything other than a number was entered.
func ExampleGetInt() {
	var age int

	for {
		age := GetInt("Enter your age:")
		if age > 0 {
			break
		}
	}

	switch {
	case age >= 18:
		fmt.Println("You are an adult!")
	default:
		fmt.Printf("Wait for %d year(s)\n", 18-age)
	}

}

// GetChar() can be used to conveniently get a single char value from the user via standard input.
// It reprompts the user if anything other than a single char was entered.
func ExampleGetChar() {
	var choice rune

	for {
		choice = GetChar("Enter Y/N:")
		switch choice {
		case 'Y':
			fmt.Printf("Thank you for being positive!")
			return
		case 'N':
			fmt.Printf("We all have our opinions!")
			return
		}
	}

}
