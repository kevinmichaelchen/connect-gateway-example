# connect-gateway-example

[Connect](https://buf.build/blog/connect-a-better-grpc) is a better gRPC.
It's debuggable and usable from the browser and with curl.

[gRPC-Gateway](https://grpc-ecosystem.github.io/grpc-gateway/) was a project
designed to seamlessly convert gRPC APIs to REST APIs.

## Getting started

### Making proto changes
```shell
make buf-mod-update buf-lint buf-gen
```

### Run the server
Run the Go server with:
```bash
go mod download
go run main.go
```

### Hit the Connect API
Using [HTTPie](https://httpie.io/):
```bash
# Add a user
http -v localhost:8080/users.v1.UserService/AddUser email="foo@bar.com"

# List users
http -v POST localhost:8080/users.v1.UserService/ListUsers "Content-Type":application/json Accept:"application/json, */*;q=0.5"
```

### Hit the REST API

gRPC-Gateway gives us a REST API for free.

#### Using curl
```bash
# Add user
curl -X POST localhost:8080/api/v1/users -d '{"email": "foo@bar.com"}'

# List users
curl localhost:8080/api/v1/users
```

#### Using HTTPie
Using [HTTPie](https://httpie.io/):
```bash
# Add user
http -b POST localhost:8080/api/v1/users email="foo@bar.com"

# List users
http -b localhost:8080/api/v1/users
```

### Swagger Spec
Open [http://localhost:8080/](http://localhost:8080/) in your browser.

## Making changes

Whenever you make a change to a Protocol Buffer (protobuf), generation will automatically update the Swagger file. 
