# Go Clean Architecture
Example that shows core principles of the Clean Architecture in Golang projects.

## <a href="https://www.zhashkevych.com/clean-architecture">Blog Post</a>

## Rules of Clean Architecture by Uncle Bob:
- Independent of Frameworks. The architecture does not depend on the existence of some library of feature laden software. This allows you to use such frameworks as tools, rather than having to cram your system into their limited constraints.
- Testable. The business rules can be tested without the UI, Database, Web Server, or any other external element.
- Independent of UI. The UI can change easily, without changing the rest of the system. A Web UI could be replaced with a console UI, for example, without changing the business rules.
- Independent of Database. You can swap out Oracle or SQL Server, for Mongo, BigTable, CouchDB, or something else. Your business rules are not bound to the database.
- Idependent of any external agency. In fact your business rules simply don’t know anything at all about the outside world. 

More on topic can be found <a href="https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html">here</a>.

### Project Description&Structure:
REST API with custom JWT-based authentication system. Core functionality is about creating and managing bookmarks (Simple clone of <a href="https://app.getpocket.com/">Pocket</a>).

#### Structure:
4 Domain layers:

- Models layer
- Repository layer
- UseCase layer
- Delivery layer

## API:

### POST /auth/sign-up

Creates new user 

##### Example Input: 
```
{
	"username": "UncleBob",
	"password": "cleanArch"
} 
```


### POST /auth/sign-in

Request to get JWT Token based on user credentials

##### Example Input: 
```
{
	"username": "UncleBob",
	"password": "cleanArch"
} 
```

##### Example Response: 
```
{
	"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzEwMzgyMjQuNzQ0MzI0MiwidXNlciI6eyJJRCI6IjAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMCIsIlVzZXJuYW1lIjoiemhhc2hrZXZ5Y2giLCJQYXNzd29yZCI6IjQyODYwMTc5ZmFiMTQ2YzZiZDAyNjlkMDViZTM0ZWNmYmY5Zjk3YjUifX0.3dsyKJQ-HZJxdvBMui0Mzgw6yb6If9aB8imGhxMOjsk"
} 
```

### POST /api/bookmarks

Creates new bookmark

##### Example Input: 
```
{
	"url": "https://github.com/zhashkevych/go-clean-architecture",
	"title": "Go Clean Architecture example"
} 
```

### GET /api/bookmarks

Returns all user bookmarks

##### Example Response: 
```
{
	"bookmarks": [
            {
                "id": "5da2d8aae9b63715ddfae856",
                "url": "https://github.com/zhashkevych/go-clean-architecture",
                "title": "Go Clean Architecture example"
            }
    ]
} 
```

### DELETE /api/bookmarks

Deletes bookmark by ID:

##### Example Input: 
```
{
	"id": "5da2d8aae9b63715ddfae856"
} 
```


## Requirements
- go 1.13
- docker & docker-compose

## Run Project

Use ```make run``` to build and run docker containers with application itself and mongodb instance

