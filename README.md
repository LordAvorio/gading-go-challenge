
# GO Mid Test Assignments

A repository of my mid test assignment in iSwift golang-cohort-3 class.

## How To Run this project locally

Clone this repository into your local PC

```bash
  git clone https://github.com/LordAvorio/gading-go-challenge.git
```

Enter the project folder and change the branch into mid-test-assignment

```bash
  git checkout mid-test-assignment
```

Edit the environment file on config.env

```bash
# CONFIG MYSQL DATABASE
HOST="127.0.0.1" # INSERT YOUR SERVER HOST
PORT="3307" # INSERT YOUR SERVER PORT
DB_USER="gadingdev" # INSERT YOUR DATABASE USER
DB_PASS="admin1!" # INSERT YOUR DATABASE PASSWORD
DB_NAME="gading-golang-mid-test" # INSERT YOUR DATABASE NAME

# CONFIG APP
APP_PORT=":8080" # INSERT YOUR PORT APPLICATION

# FLAG
AUTO_MIGRATE=false # CHANGE TRUE IF YOU WANT RUN MIGRATION DB
```

Run the main.go file

```bash
  go run main.go
```

