## LOAN ENGINE WITH GO
This project is base structure to create a RestFul using Go. In this project I tried to solve loan mechanism process.

---

## Features

- STAFF
- BORROWER
- LOAN
- INVESTMENT

--- 

## Accessibility
All request required `userIp` in header.

## Requirements
Ensure your system meets the following requirements:
- Go >= 1.24.1
- PostgreSQL >= 14,17

## Installation

### Step 1: Clone the Repository
```bash
git clone https://github.com/AradeaTechno/loanEngine.git
cd loanEngine
```

## Step 2: Install Dependencies

Run the following command to install dependencies:
```bash
go mod tidy
```
Incase your application mod not changing, please open `go.mod` and change the modules name into your root project's name.

## Step 3: Environment Configuration
Change env requirements with yours:
```bash
nano .env
```

## Step 4: Run Application
Run the following command to run the application:
```bash
go run main.go
```

## Step 5: API Tester
I've included postman collection to be used to test the API. 
Import this postman collection below to your postman.
```bash
AMARTHA.postman_collection.json
```

## Step 6: Database upload
I've included all publics to be used on your postgre.
```bash
dump-amarthaloan-202503261629.sql
```
