# 🚀 Go Projects Collection

This repository contains a collection of my projects **Go**.
The goal of this repo is to document my journey with Go while building practical applications, utilities, and exploring language features like concurrency, modules, and APIs.

## 🏗️ Projects


### WebServer (X clone)


**Features:**
- A simple web server built in Go.
- API endpoints for creating and retrieving "tweets".
- User creation and management.
- **JWT-based** authentication with refresh tokens.
 - Secure login with password hashing using bcrypt.
- Uses sqlc to generate type-safe Go code from raw SQL queries.

### BlogAggregator(Gator)


**Features:**
- User registration and login
- Add, list, follow, and unfollow RSS feeds
- Aggregate feeds at specified intervals
- Reset the database

<details>
<summary>Installation</summary>

### 📦 Installation

You will need PostgreSQL and Go installed on your system to run Gator.For Arch-based systems, run the following (for other distributions, refer to your package manager):
`$ sudo pacman -S go postgresql`

Set up PostgreSQL and the Go environment:<br>
a. Configure PostgreSQL
```bash
$ sudo passwd postgres  # Set a password for the PostgreSQL user
$ sudo systemctl start postgresql.service  # Start the PostgreSQL service
```

b. Create the Gator database
```bash
$ sudo -u postgres psql  # Use psql to interact with the database # Inside the psql prompt:
$ CREATE DATABASE gator; $ALTER USER postgres PASSWORD 'postgres';
```

c. Install Goose for migrations and sqlc for converting SQL queries to Go code
```bash
$ go install github.com/pressly/goose/v3/cmd/goose@latest
$ go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

d. Clone this repo
```bash
    git clone https://github.com/shri771/Go.git
    cd Go/BlogAggregator
```
### 🚀 Usage
First you need to create user and then you can run other cmds.<br>
***Here are cmds you can try.***
```bash
   $ go run . login usrename         #Login a  user
   $ go run . register username             #register a user
   $ go run . reset                 # Reset database use with caution
   $ go run . users                 # List registed user
   $ go run . agg   1s                # Aggregate the feed every 1 sec
   $ go run . addfeed name_feed user             # Add New feed
   $ go run . feeds                # List feed for current user
   $ go run . follow   url             # Follow a feed
   $ go run . following         # Following feed
   $ go run . unfollow  url            # unfollow Feed

```

### 🏗️ Project Documentation
#### Introdution
t

</details>

### Pokedexcli


**Features:**
- A CLI-based Pokedex application.
- Explore different locations to find Pokemon.
- Catch Pokemon and add them to your Pokedex.
- View details of the Pokemon you have caught.

