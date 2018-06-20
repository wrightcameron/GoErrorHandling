# Go Error Handling

* Author: Cameron Wright

## Overview

Repository for understanding Golang error handling, syntax, libraries, and testing.

The program asks for an address and will give the sun rise and sun set at that 
location on the current date.

## Manifest

|Filename       | Description                                                       |
|---------------|-------------------------------------------------------------------|
|main.go        | sunset & sunrise calculating program                              |
|README.md      | This file                                                         |
|benchmark_test.go|Contains functions for benchmarking main.go                      |
|results.txt    |Information about current unit tests, test coverage, and benchmarks|
|unit_test.go   |Contains functions for unit testing main.go                        |

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

Unit Testing was started on this project, but not as far as I wanted.  Two functions
have unit tests for them out of three.  One function getUserAddress() requires user
input so that makes it harder to automate, and I did not count it.  Unit Testing has
a test coverate at 20%

Benchmarking was also done on the same functions and from the results shown in
results.txt getCoordinatesFromResponse is the bottleneck but that is where more work is
happening compared to calculateSunRiseSet.  

### Known Bugs

The exact times for the sunrise and sunset seem alittle off, and this could
be due to a logical bug.

## Discussion

This project was an introduction into Go because while I have known about it before
I have never dived into it as much as this project made me.  The first library
I was going to focus on was the color library and I was going to take a text file,
scan it, and color it someway based on its contects.  The problem I faced was that
color newlines after coloring something so I could not return the text file back to
its original newlines.  The color library mentions newlining alot, but unlike other
flags that can be changed apparently newlining is something that to my review can
not be turned off.

The second "verson" of this project was something I wanted to do with date & time I
saw in the Awesome-Go library.  I first wanted to make a program that would find your
computers bit system aka 32bit, 64bit, etc and calculate how much more time you have
before the next Y2K scenario.  I gave up that idea after looking at the kind of data
Golang can get about an OS and thinking about how would you take a date and add the
total number of bytes to it.  This is a do able project but this is me just getting 
into Golang.  The final idea was to take a library that finds when the sun sets and when
it rises and find when it does so with a given address.  This seemed like a challenge
because sunset/sunrise library takes coordinates not address so I would need to find
a way to convert them over.  Thankfully Google Maps has an API that you can send a
GET Request to with the address and get the coordinates back.  Though with this
route I would need to use the REST golang library and I am still alittle new at 
REST calls and using JSON.

Testing wise Golang has awsome built in testing tools in its language.  All I had
to do was create the test file with the correct name and Go would run the test for
me with the correct command.  I also found out that Go also has built in benchmarking
which makes getting these build not hard at all.  This is a dream compared to java
where if you want to unit test you need the library JUnit and some other benchmarking
library.
