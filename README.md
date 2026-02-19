# GoLang


## I have also impliment GOF Design patterns.


### type is use to create custom types (Domain Modeling)

In Java:

Everything revolves around class

In Go:

Everything revolves around type

Struct → type
Interface → type
Enum-like → type
Function → type
Custom primitive → type

It’s the foundational building block of Go.


- Without type, your code becomes this:

    func Transfer(from string, to string, amount int64)


- With type, your code becomes this:

    func Transfer(from AccountID, to AccountID, amount Money)


--> Http server with Routing -> running(listening) at localhost : 8080
    routing :
        1. /
        2. /users
        3. /hi

--> http client -> request to localhost:8080

go context package, channel cancelation, select statement
go concurrency patterns
grpc comminication in java
samll crud api with go
microservices develpement with go
main design patterns implemented with java and go