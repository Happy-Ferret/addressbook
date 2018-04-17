# Addressbook

This is a very early in the life cycle implementation of a RESTful addressbook implementation.

I had hoped to get it done this weekend, but unfortantly did not get very much downtime to work on it.

Currently, the application is hard coded with 3 entries. 

It supports:
* a list all function 
  - `localhost:8080/entry`
* a search function that lists entries based on id, first or last name.
  - `localhost:8080/entry/1` -> returns json object entry with id=1
  - `localhost:8080/entry/Doe` -> returns json object with entries with the lastName=Doe

TODO:
  [x] Create entry -> html form (basic implementation done)
  [] Modify entry -> html form (?)
  [] Delete entry
  [] add some sort of data persistance (database)
  [] Import from CSV
  [] Export to CSV

My original implementation, which you can see in `old/` I had tried to use 
protobufs because of their efficiency but I ended up spiraling down a gRPC rabbit hole that took me away from the original problem.
So I started over and just went with simple json objects for storage. 


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

### Domain model
Domain models hold data
* Address book 
    * Holds people
* Person
    - First Name
    - Last Name
    - Email
    - Phone 

### Services
  * List All Entries
  * Show specific single entry
  * Add entry
  * Modify entry
  * delete entry




## (potential) Future additions

* users -> each user has their own addressbook.

