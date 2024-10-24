# 📚 Library Management System

A simple **Library Management System** built with **Go (Golang)** and Cassandra. This system allows **administrators** to manage books and users, and **students** to borrow and return books through a set of well-defined APIs. Below, you’ll find how the system works in a clear and organized way.

## 🛠 How the System Works

The system has two main roles: **Administrators** and **Students**. Each role has specific tasks they can perform using HTTP endpoints. Below is a breakdown of the key operations and how they are handled by the Go backend.

## 🔌 How It All Starts
1. `main.go`
- The entry point of the application.
- **Initializes the Cassandra** connection.
- Defines routes for admin and student operations.
- Starts the server at `http://localhost:8080.`
    ```bash
    handlers.InitCassandraSession("127.0.0.1")
    log.Fatal(http.ListenAndServe(":8080", nil))
## 🚀 Admin Operations

- **API Configuration in documentation**

## 🧩 Summary

The **Library Management System** offers:

- 📚 **Admin Features:** Manage books and users.
- 🎓 **Student Features:** Borrow and return books.
- ⚡ **Cassandra Integration:** Stores all data in a NoSQL database.
This project provides a simple way to manage a library, ensuring data integrity through backend validation and fast access through **Cassandra**.

## 🎯 How to Start

1. **Run Cassandra** on your machine

2. **Initialize the Go project**
    ```bash
    go mod init library-system
    go mod tidy
3. **Start the server:**
    ```bash
    go run main.go
4. **Test the endpoints** using Postman or Curl

#### Enjoy building with Go and Cassandra!

## 🐳 Running Cassandra with Docker
Follow these steps to start Cassandra using Docker.

1. Pull the Cassandra Docker image:
```docker pull cassandra:latest```
2. Run Cassandra in a container:
```docker run --name cassandra -d -p 9042:9042 cassandra:latest```