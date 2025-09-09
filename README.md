# 🚀 Go Projects Collection

This repository contains a collection of my projects, and learning exercises in **Go (Golang)**.  
The goal of this repo is to document my journey with Go while building practical applications, utilities, and exploring language features like concurrency, modules, and APIs.

##Installing and Running Guide for Blog Aggregator (Gator) ⚡
###📦 Installation
---

You will need PostgreSQL and Go installed on your system to run Gator.For Arch-based systems, run the following (for other distributions, refer to your package manager):  
```$ sudo pacman -S go postgresql```


Set up PostgreSQL and the Go environment:a. Configure PostgreSQL  
```$ sudo passwd postgres  # Set a password for the PostgreSQL user $ sudo systemctl start postgresql.service  # Start the PostgreSQL service```

b. Create the Gator database  
```$ sudo -u postgres psql  # Use psql to interact with the database # Inside the psql prompt: $ CREATE DATABASE gator; $ALTER USER postgres PASSWORD 'postgres';```

c. Install Goose for migrations and sqlc for converting SQL queries to Go code  
```$ go install github.com/pressly/goose/v3/cmd/goose@latest $ go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest```


