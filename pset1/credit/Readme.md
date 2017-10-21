## Credit.go
---
### tl;dr

Implement a program that determines whether a provided credit card number is valid according to Luhn’s algorithm.

```
$ ./credit
Number: 378282246310005
AMEX
```

### Specification
* Write a program that prompts the user for a credit card number and then reports (via printf) whether it is a valid American Express, MasterCard, or Visa card number, per the definitions of each’s format herein.

* So that we can automate some tests of your code, we ask that your program’s last line of output be AMEX\n or MASTERCARD\n or VISA\n or INVALID\n, nothing more, nothing less, and that main always return 0.

* For simplicity, you may assume that the user’s input will be entirely numeric (i.e., devoid of hyphens, as might be printed on an actual card).

* See http://docs.cs50.net/problems/credit/credit.html

### To Run
```
go build credit.go
./credit
```