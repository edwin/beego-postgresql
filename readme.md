# beego-postgresql

A minimal **Beego + PostgreSQL** example application written in Go.

This repository demonstrates how to build a simple web API using the Beego framework with PostgreSQL as the database backend. It is intentionally small and easy to understand, making it suitable for learning or as a starting point for new projects.

---

## Features

- Beego web framework
- PostgreSQL database integration
- MVC-style project structure
- Simple REST endpoint example
- Ready for extension (OpenTelemetry, Docker, etc.)

---

## Project Structure

```
.
├── controllers
│   └── default.go
├── models
│   └── customer.go
├── routers
│   └── router.go
├── main.go
├── go.mod
├── Dockerfile
├── .gitignore
└── README.md
```

---

## Getting Started

### Prerequisites

- Go 1.18 or newer
- PostgreSQL
- Basic knowledge of Go modules

---

### Install Dependencies

From the project root:

```bash
$ go mod tidy
```

---

## Database Setup

Create a PostgreSQL database (example):

```bash
$ createdb beego
```

Make sure your database connection string matches your local setup.

---

## Running the Application

Build the binary:

```bash
$ GOOS=linux GOARCH=amd64 go build -o beego-postgresql .
```

Run the application:

```bash
$ ./beego-postgresql
```

Access the API:

```bash
$ curl -kv http://localhost:8080/
```

---

## Example

```
$ curl -kv http://localhost:8080/
> GET / HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.61.1
> Accept: */*
> 
< HTTP/1.1 200 OK
< Content-Length: 161
< Content-Type: application/json; charset=utf-8
< Date: Tue, 23 Dec 2025 15:24:07 GMT
< 
* Connection #0 to host customer-svc-beego left intact
[{"Id":1,"Name":"Ryoko Hirosue"},{"Id":2,"Name":"Kasumi Arimura"},{"Id":3,"Name":"Hikari Mitsushima"},{"Id":4,"Name":"Nanao Arai"},{"Id":5,"Name":"Mikako Tabe"}]
```

---

## Podman (Optional)

A Dockerfile is included.

Build image:

```bash
$ podman build -t beego-postgresql .
```

Run container:

```bash
$ podman run -p 8080:8080 beego-postgresql
```

---

## References

- Beego Framework: https://github.com/beego/beego
- PostgreSQL: https://www.postgresql.org/
- Go Modules: https://go.dev/doc/modules

---

## License

No license specified. Add one if you plan to distribute or reuse this project.
