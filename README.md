# Go Error Handling

* Author: Cameron Wright

## Overview

Repository for understanding Golang error handling, syntax, libraries, and testing.

The program asks for an address and will give the sun rise and sun set at that 
location on the current date.

## Manifest

|Filename       | Description                                                       |
|---------------|-------------------------------------------------------------------|
|main.go         | sunset & sunrise calculating program                             |
|README.md      | This file                                                         |

## Building the project

Go as a language has a built in build automation tool that makes life easier.  Assuming
that you have go installed, to build the main file, run the command 

$ go build

A binary file titled "GoErrorHandling" will be created that can be executed with

$ ./GoErrorHandling

Follow instructins of program and output will be pushed to console. To clean 
repository run following command

$ go clean

## Features and usage

The program finds the sunset and sunrise of a given address.  When first started
the program will ask for user input of an address to be typed in to the console.
Afterwards the sunrise and sunset will be shown as times printed to the console.

## Testing

This section should detail how you tested your code. Simply stating "I ran it a few times and
it seems to work" is not sufficient. Your testing needs to be detailed here or even better,
this section should document the appropriate smoke/unit/system tests that you wrote. For a
homework, I don't expect much but for a project this section should have significant content.

### Known Bugs

List known bugs that you weren't able to fix (or ran out of time to fix).

## Discussion

This project was an introduction into Go because while I have known about it before
I have never dived into it as much as this project made me.  The first library
I was going to focus on was the color library and I was going to take a text file,
scan it, and color it someway based on its contects.  The problem I faced was that
color newlines after coloring something so I could not return the text file back to
its original newlines.  The color library mentions newlining alot, but unlike other
flags that can be changed apparently newlining is something that to my review can
not be turned off.  I may be wrong but this was with me using the first examples
on the example page.

The second "verson" of this project was something I wanted to do with date & time I
saw in the Awesome-Go library.  I first wanted to make a program that would find your
computers bit system aka 32bit, 64bit, etc and calculate how much more time you have
before the next Y2K scenario.  I past that idea after looking at the kind of data
Golang can get about an OS and thinking about how would you take a date and add the
total number of bytes to it.  This is a doable project but this is me just getting 
into Golang.  The final idea was to take a library that finds when the sun sets and when
it rises and find when it does so with a given address.  This seemed like a challenge
because sunset/sunrise library takes coordinates not address so I would need to find
a way to convert them over.  Thankfully Google Maps has an API that you can send a
GET Request to with the address and get the coordinates back.  Though with this
route I would need to use the REST golang library and I am still new at REST calls
and using JSON.
