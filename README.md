# 🚀 Go Projects Collection

This repository contains a collection of my projects, and learning exercises in **Go (Golang)**.  
The goal of this repo is to document my journey with Go while building practical applications, utilities, and exploring language features like concurrency, modules, and APIs.

## Installing and ruuning Guide Blog Aggregator(Gator) ⚡
### 📦 Installation
---
1.You will need postgress and Go installed on you system to run gator.
**For Arch System Do this** (for other distros you can find it on other pkg manager)
```  $ Sudo pacman -S go postgresql  ```
2.Now we will Setup **postgres** and **go enviroment**
  1.Setup postgres
'''$ sudo passwd postgres \\ to set pasword for postgres
  $ sudo systemctl start postgres.service \\ to start posgress service```
  
  2.Create gator dabase
  ```$ sudo -u postgres psql \\ we will be using psql to interact with database and after this you will be in new psql prompt
  $ CREATE DATABASE gator;
  $ ALTER USER postgres PASSWORD 'postgres';```

  3. Also we will need goose for  the migration and sqlc to convert SQL Queries to Go code
  ''' $ go install github.com/pressly/goose/v3/cmd/goose@latest
      $ go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest '''
