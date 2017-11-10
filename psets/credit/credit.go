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

/*
credit implements a program that determines whether a provided credit card number is valid
according to Luhn’s algorithm.

Usage:

	$ ./credit
	Number: 378282246310005
	AMEX

*/
package main

import (
	"strconv"

	"github.com/arjun-rao/cs50/cs50"
)

// getCreditCard prompts the user to enter a credit card number, and reprompts on invalid input.
// It returns a valid 64 bit integer number.
func getCreditCard() (int64, string) {
	for {
		input := cs50.GetString("Number: ")
		number, err := strconv.ParseInt(input, 10, 64)
		if err == nil && number >= 0 {
			return number, input
		}
	}
}

// validate returns the credit card type if card is valid, else returns "INVALID".
func validate(card string) string {
	oddSum, evenSum := 0, 0
	cardLength := len(card)
	// Handle invalid length cases first
	if cardLength < 13 || cardLength > 16 || cardLength == 14 {
		return "INVALID"
	}
	isEven := cardLength%2 == 0
	// Multiply every other digit by 2 and add non multiplied digits starting with
	// second last digit. Add the totals and check last digit to see if 0 (valid syntactically)
	for _, digit := range card {
		val, _ := strconv.Atoi(string(digit))
		if !isEven {
			oddSum += val
			isEven = true
		} else {
			for temp := val * 2; temp != 0; temp /= 10 {
				evenSum += temp % 10
			}
			isEven = false
		}
	}
	totalSum := oddSum + evenSum
	if totalSum%10 == 0 {
		// Handle validity of card types
		switch {
		case card[0] == '4' && (cardLength == 13 || cardLength == 16):
			return "VISA"
		case (card[:2] == "34" || card[:2] == "37") && cardLength == 15:
			return "AMEX"
		case card[0] == '5' && card[1] > '0' && card[1] < '6' && cardLength == 16:
			return "MASTERCARD"
		default:
			return "INVALID"
		}
	}
	return "INVALID"
}

func main() {
	// Get credit card number from the user.
	_, card := getCreditCard()

	// Multiply change by 100 and round to get dollars value in number of cents.
	println(validate(card))
}
