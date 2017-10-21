## Mario.go
---
### tl;dr
Implement a program that prints out a half-pyramid of a specified height, per the below.

```
$ ./mario
Height: 5
    ##
   ###
  ####
 #####
######

$ ./mario
Height: 3
  ##
 ###
####
```

### Specification

* Write, in a file called mario.c in ~/workspace/pset1/mario/less/, a program that recreates this half-pyramid using hashes (#) for blocks.

* To make things more interesting, first prompt the user for the half-pyramid’s height, a non-negative integer no greater than 23. (The height of the half-pyramid pictured above happens to be 8.)

* If the user fails to provide a non-negative integer no greater than 23, you should re-prompt for the same again.

* Then, generate (with the help of printf and one or more loops) the desired half-pyramid.

* Take care to align the bottom-left corner of your half-pyramid with the left-hand edge of your terminal window.

### To Run

```
go build mario.go
./mario
```