package cs50

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
	"strings"
	"testing"
)

// Source: https://stackoverflow.com/a/39571615
const (
	MinUint uint = 0 // binary: all zeroes

	// Perform a bitwise NOT to change every bit from 0 to 1
	MaxUint = ^MinUint // binary: all ones

	// Shift the binary number to the right (i.e. divide by two)
	// to change the high bit to 0
	MaxInt = int(MaxUint >> 1) // binary: all ones except high bit

	// Perform another bitwise NOT to change the high bit to 1 and
	// all other bits to 0
	MinInt = ^MaxInt // binary: all zeroes except high bit
)

func TestGetString(t *testing.T) {
	tests := []struct {
		name       string
		prompt     string
		rawInput   string
		wantInput  string
		wantOutput string
		wantError  string
	}{
		{
			"Input without Prompt",
			"",
			"My input string\n",
			"My input string",
			"",
			"",
		},
		{
			"Input with Prompt",
			"Enter a string:",
			"My input string\n",
			"My input string",
			"Enter a string:",
			"",
		},
		{
			"Blank Input",
			"",
			"\r\n",
			"",
			"",
			"",
		},
	}
	var outbuf, errbuf bytes.Buffer

	for _, tt := range tests {
		outbuf.Reset()
		errbuf.Reset()
		stdIn = strings.NewReader(tt.rawInput)
		stdOut = &outbuf
		stdErr = &errbuf
		t.Run(tt.name, func(t *testing.T) {
			if gotInput := GetString(tt.prompt); gotInput != tt.wantInput {
				t.Errorf("GetString() = %v, want %v", gotInput, tt.wantInput)
			}
			if gotOutput := outbuf.String(); gotOutput != tt.wantOutput {
				t.Errorf("GetString() produced output: %v, expected: %v", gotOutput, tt.wantOutput)
			}
			if gotError := errbuf.String(); gotError != tt.wantError {
				t.Errorf("GetString() produced output: %v, expected: %v", gotError, tt.wantError)
			}
		})
	}
}

func TestGetInt(t *testing.T) {
	tests := []struct {
		name       string
		prompt     string
		rawInput   string
		wantInput  int
		wantOutput string
		wantError  string
	}{
		{
			"Maximum Integer Value",
			"",
			fmt.Sprintf("%s\n", strconv.Itoa(MaxInt)),
			MaxInt,
			"",
			"",
		},
		{
			"Input with Prompt",
			"Enter a number:",
			fmt.Sprintf("%s\n", strconv.Itoa(MaxInt)),
			MaxInt,
			"Enter a number:",
			"",
		},
		{
			"Lowest Negative Number",
			"Enter a number:",
			fmt.Sprintf("%s\n", strconv.Itoa(MinInt)),
			MinInt,
			"Enter a number:",
			"",
		},
	}
	var outbuf, errbuf bytes.Buffer

	for _, tt := range tests {
		outbuf.Reset()
		errbuf.Reset()
		stdIn = strings.NewReader(tt.rawInput)
		stdOut = &outbuf
		stdErr = &errbuf
		t.Run(tt.name, func(t *testing.T) {
			if gotInput := GetInt(tt.prompt); gotInput != tt.wantInput {
				t.Errorf("GetInt() = %v, want %v", gotInput, tt.wantInput)
			}
			if gotOutput := outbuf.String(); gotOutput != tt.wantOutput {
				t.Errorf("GetInt() produced output: %v, expected: %v", gotOutput, tt.wantOutput)
			}
			if gotError := errbuf.String(); gotError != tt.wantError {
				t.Errorf("GetInt() produced output: %v, expected: %v", gotError, tt.wantError)
			}
		})
	}
}

func TestGetFloat(t *testing.T) {
	tests := []struct {
		name       string
		prompt     string
		rawInput   string
		wantInput  float64
		wantOutput string
		wantError  string
	}{
		{
			"Maximum Float Value",
			"",
			fmt.Sprintf("%s\n", strconv.FormatFloat(math.MaxFloat64, 'e', -1, 64)),
			math.MaxFloat64,
			"",
			"",
		},
		{
			"Input with Prompt",
			"Enter a number:",
			fmt.Sprintf("%s\n", strconv.FormatFloat(math.MaxFloat64, 'e', -1, 64)),
			math.MaxFloat64,
			"Enter a number:",
			"",
		},
		{
			"Lowest Negative Float",
			"Enter a number:",
			fmt.Sprintf("%s\n", strconv.FormatFloat(math.SmallestNonzeroFloat64, 'g', -1, 64)),
			math.SmallestNonzeroFloat64,
			"Enter a number:",
			"",
		},
	}
	var outbuf, errbuf bytes.Buffer

	for _, tt := range tests {
		outbuf.Reset()
		errbuf.Reset()
		stdIn = strings.NewReader(tt.rawInput)
		stdOut = &outbuf
		stdErr = &errbuf
		t.Run(tt.name, func(t *testing.T) {
			if gotInput := GetFloat(tt.prompt); gotInput != tt.wantInput {
				t.Errorf("GetFloat() = %v, want %v", gotInput, tt.wantInput)
			}
			if gotOutput := outbuf.String(); gotOutput != tt.wantOutput {
				t.Errorf("GetFloat() produced output: %v, expected: %v", gotOutput, tt.wantOutput)
			}
			if gotError := errbuf.String(); gotError != tt.wantError {
				t.Errorf("GetFloat() produced output: %v, expected: %v", gotError, tt.wantError)
			}
		})
	}
}

func TestGetLong(t *testing.T) {
	tests := []struct {
		name       string
		prompt     string
		rawInput   string
		wantInput  int64
		wantOutput string
		wantError  string
	}{
		{
			"Maximum Integer Value",
			"",
			fmt.Sprintf("%s\n", strconv.FormatInt(math.MaxInt64, 10)),
			math.MaxInt64,
			"",
			"",
		},
		{
			"Input with Prompt",
			"Enter a number:",
			fmt.Sprintf("%s\n", strconv.FormatInt(math.MaxInt64, 10)),
			math.MaxInt64,
			"Enter a number:",
			"",
		},
		{
			"Lowest Negative Number",
			"Enter a number:",
			fmt.Sprintf("%s\n", strconv.FormatInt(math.MinInt64, 10)),
			math.MinInt64,
			"Enter a number:",
			"",
		},
	}
	var outbuf, errbuf bytes.Buffer

	for _, tt := range tests {
		outbuf.Reset()
		errbuf.Reset()
		stdIn = strings.NewReader(tt.rawInput)
		stdOut = &outbuf
		stdErr = &errbuf
		t.Run(tt.name, func(t *testing.T) {
			if gotInput := GetLong(tt.prompt); gotInput != tt.wantInput {
				t.Errorf("GetLong() = %v, want %v", gotInput, tt.wantInput)
			}
			if gotOutput := outbuf.String(); gotOutput != tt.wantOutput {
				t.Errorf("GetLong() produced output: %v, expected: %v", gotOutput, tt.wantOutput)
			}
			if gotError := errbuf.String(); gotError != tt.wantError {
				t.Errorf("GetLong() produced output: %v, expected: %v", gotError, tt.wantError)
			}
		})
	}
}

func TestGetChar(t *testing.T) {
	tests := []struct {
		name       string
		prompt     string
		rawInput   string
		wantInput  rune
		wantOutput string
		wantError  string
	}{
		{
			"Input without Prompt",
			"",
			"A\n",
			rune('A'),
			"",
			"",
		},
		{
			"Input with Prompt",
			"Enter Y/N:",
			"Y\n",
			rune('Y'),
			"Enter Y/N:",
			"",
		},
	}
	var outbuf, errbuf bytes.Buffer

	for _, tt := range tests {
		outbuf.Reset()
		errbuf.Reset()
		stdIn = strings.NewReader(tt.rawInput)
		stdOut = &outbuf
		stdErr = &errbuf
		t.Run(tt.name, func(t *testing.T) {
			if gotInput := GetChar(tt.prompt); gotInput != tt.wantInput {
				t.Errorf("GetChar() = %v, want %v", gotInput, tt.wantInput)
			}
			if gotOutput := outbuf.String(); gotOutput != tt.wantOutput {
				t.Errorf("GetChar() produced output: %v, expected: %v", gotOutput, tt.wantOutput)
			}
			if gotError := errbuf.String(); gotError != tt.wantError {
				t.Errorf("GetChar() produced output: %v, expected: %v", gotError, tt.wantError)
			}
		})
	}
}
