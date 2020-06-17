<div align="center"><h1>Clean Architecture Notes API</h1></div>

This is my implementation of [clean architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) for a golang REST API.

### What is clean architecture

Clean architecture is a highly modular architecture pioneered by Robert C. Martin (Uncle Bob). There are several layers of abstraction for the separation of core business rules from everything else.

### Why clean architecture?

It creates a system which is 

1. Independent of Frameworks

    The architecture does not depend on the existence of some library of feature laden software. This allows you to use such frameworks as tools, rather than having to cram your system into their limited constraints.

2. Testable

    The business rules can be tested without the UI, Database, Web Server, or any other external element.

3. Independent of UI

    The UI can change easily, without changing the rest of the system. A Web UI could be replaced with a console UI, for example, without changing the business rules.
4. Independent of Database 
    You can swap out Oracle or SQL Server, for Mongo, BigTable, CouchDB, or something else. Your business rules are not bound to the database.
5. Independent of any external agency

    In fact your business rules simply don’t know anything at all about the outside world.


### Clean architecture components

![](https://blog.cleancoder.com/uncle-bob/images/2012-08-13-the-clean-architecture/CleanArchitecture.jpg)


1. Entities

    Entities are similar to models in MVC. They are an OOPS representation of a real world entity.

2. Use Cases (Repository)

    Repository is an abstraction over the entity. It houses the lower level business rules related to database connections. This is independent of the database service used

3. Services

    Services are yet another abstraction layer over the application level business rules. This is an interface for higher level business logic oriented functions.
    
4. Interface Adapters

    Each interface has its separate directory. Since we made an API we have an api directory

5. Frameworks and Drivers

    In our case, it is Web since this is the layer where all the detials go. In our case, we have implemented handlers for request-response model


### References

See this article on Medium - [Clean Architecture, the right way (by Angad Sharma)](https://medium.com/gdg-vit/clean-architecture-the-right-way-d83b81ecac6)

DSC VIT's template repo of clean architecute in Go - [Go template repository](https://github.com/GDGVIT/go-template)

An implementation for a full fleged rest api and clean code - [Bandersnatch by supercmmetry](https://github.com/supercmmetry/bandersnatch)

Another implementation I've referred to while learning this - [Clean Notes API by Rithik Jain](https://github.com/rithikjain/CleanNotesApi)