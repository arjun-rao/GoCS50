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
Cash implements a program that calculates the minimum number of coins required to give a user change.

Usage:

	$ ./cash
	Change owed: 0.41
	4

*/
package main

import (
	"math"

	"github.com/arjun-rao/cs50/cs50"
)

// getChange prompts the user to enter a float, and reprompts on invalid input.
// It returns a valid 64 bit floating point number.
func getChange() float64 {
	for {
		change := cs50.GetFloat("Change: ")
		if change >= 0 {
			return change
		}
	}
}

// greedyCoins returns the minimum number of coins of denominations 25c, 10c,5c, 1c needed
// to sum to the given number of dollars and cents.
func greedyCoins(dollars int, cents int) int {
	coins := dollars * 4
	for cents > 0 {
		switch {
		case cents >= 25:
			cents -= 25
			coins++
		case cents >= 10:
			cents -= 10
			coins++
		case cents >= 5:
			cents -= 5
			coins++
		default:
			coins += cents
			cents = 0
		}
	}
	return coins
}

func main() {
	// Get required change from the User.
	change := getChange()

	// Multiply change by 100 and round to get dollars value in number of cents.
	temp := math.Floor(change*100 + 0.5)
	dollars := int(temp)
	cents := dollars % 100
	dollars /= 100
	println(greedyCoins(dollars, cents))
}
