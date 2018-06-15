# Addressbook
This project is a RESTful implementation of an addressbook.

It supports:

    * Listing all entries
    * Modifying specific entries
    * Deleting specific entries
    * Adding an entry
    * Import from CSV
    * Export to CSV


## Instructions to run locally

Get the source code:

- `go get github.com/JessTheBell/addressbook` 


Navigate to the project folder:

- `cd $GOPATH/src/github.com/JessTheBell/addressbook `

Start the project:

- `go run main.go`

Open `localhost:8080` in the browser of your choice.

Have Fun!


There is a disabled populate() function. To start the server with some default entries, uncomment that line in main.

-------- 

## Problem 
The following project is intended to give you an opportunity to demonstrate your 
understanding of web service and micro service concepts.

The requirements are laid out in a simple user story that intentionally leaves room for interpretation. 

Please use your best judgment as you decide how to best fulfill the requirements outlined. 


### Requirements: 

As a user I need an online address book exposed as a REST API.

I need the data set to include the following data fields: 

First Name, Last Name, Email Address, and Phone Number

I need the api to follow standard rest semantics to support listing entries, 
showing a specific single entry, and adding, modifying, and deleting entries.

The code for the address book should include regular go test files that demonstrate how to exercise all operations of the service.

Finally I need the service to provide endpoints that can export and import the address book data in a CSV format.


## How to turn in the project:

Please post the code to a publicly accessible github, or other CSV repository and provide a link to the completed code.
If your solution includes a dependency on any external system like a database or other system,
please provide instructions for configuring the external system in the readme in your repository.


----------------- 
