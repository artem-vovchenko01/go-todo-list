Simple Golang API for creating todo notes

To run:

```console
foo@bar:~$ cd backend
foo@bar:~$ go run .
```

Application will be listening on localhost:8080

API routes are listed in main.go 

Example of signin in:

```console
foo@bar:~$ curl -X POST localhost:8080/user/signup --data '{ "email" : "test@mail.com", "password" : "12345678" } '
```

To retrieve JWT token run:
```console
foo@bar:~$ JWT=$(curl -X POST localhost:8080/user/signin --data '{ "email" : "test@mail.com", "password" : "12345678" }' )
```

