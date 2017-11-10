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
//
// CS50 Library for Go
// Based on CS50 Library for Python
// https://github.com/cs50/python-cs50

// Package cs50 implements helper functions for reading from standard input. It allows for
// reading Chars, Strings, Ints, Floats and Longs. Each of the input functions allows
// users to supply a prompt string which will be presented to the user via standard output before
// reading their input. If input is invalid or erroneous the user will be prompted to try again.
package cs50

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var stdOut, stdErr io.Writer
var stdIn io.Reader

func init() {
	// Define the standard input, output and error streams for use by the helper functions
	stdIn = os.Stdin
	stdOut = os.Stdout
	stdErr = os.Stderr
}

// GetInt reads a line of text from standard input and returns the equivalent int.
// Takes a prompt string as parameter which is used to ask for input.
// If the text does not represent an int, user is prompted to retry.
func GetInt(prompt string) (result int) {
	for {
		input := GetString(prompt)
		result, err := strconv.Atoi(input)
		if err == nil {
			return result
		}
		if len(prompt) == 0 {
			fmt.Fprint(stdOut, "Retry: ")
		}
	}
}

// GetChar reads a line of text from standard input and returns the equivalent rune.
// Takes a prompt string as parameter which is used to ask for input.
// If the text does not represent a char, user is prompted to retry.
func GetChar(prompt string) (result rune) {
	for {
		input := GetString(prompt)
		if len(input) == 1 {
			return rune(input[0])
		}
		if len(prompt) == 0 {
			fmt.Fprint(stdOut, "Retry: ")
		}
	}
}

// GetLong reads a line of text from standard input and returns the equivalent int64.
// Takes a prompt string as parameter which is used to ask for input.
// If the text does not represent an int, user is prompted to retry.
func GetLong(prompt string) (result int64) {
	for {
		input := GetString(prompt)
		result, err := strconv.ParseInt(input, 10, 64)
		if err == nil {
			return result
		}
		if len(prompt) == 0 {
			fmt.Fprint(stdOut, "Retry: ")
		}
	}
}

// GetFloat reads a line of text from standard input and returns the equivalent float64.
// Takes a prompt string as parameter which is used to ask for input.
// If the text does not represent an int, user is prompted to retry.
func GetFloat(prompt string) (result float64) {
	for {
		input := GetString(prompt)
		result, err := strconv.ParseFloat(input, 64)
		if err == nil {
			return result
		}
		if len(prompt) == 0 {
			fmt.Fprint(stdOut, "Retry: ")
		}
	}
}

// GetString reads a line of text from standard input and returns it as a string,
// sans trailing line ending, leading and trailing spaces.
// If user inputs only a line ending, returns "".
// Takes a prompt string as parameter which is used to ask for input.
func GetString(prompt string) (input string) {
	inputScanner := bufio.NewScanner(stdIn)
	if len(prompt) != 0 {
		fmt.Fprint(stdOut, prompt)
	}
	inputScanner.Scan()
	input = inputScanner.Text()
	if err := inputScanner.Err(); err != nil {
		fmt.Fprintf(stdErr, "%v", err)
		input = ""
		return
	}
	input = strings.TrimSpace(input)
	return
}
