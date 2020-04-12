//This file is under Model package
package Model

//To connect to SQL Database
import "database/sql"
//Importing MYSQL driver to connect to SQL Database
import _ "github.com/go-sql-driver/mysql"
//To display log messages
import "log"
//fmt package - to print out statements
import "fmt"

//Variable to store instance of connection to DB
var con *sql.DB
//Function that returns the instance of database 
func Connect() *sql.DB{

	//Opening the DB
	db, err := sql.Open("mysql", "root:fanatic@tcp(localhost:3306)/sk7_db")
	
	//If there exists an error
	if err != nil{
		log.Fatal(err)
	}

	//Assigning the instance of DB for using this in DB Operations under Model package
	con = db
	//Printed in terminal after successfuly connected with specified DB
	fmt.Println("Connected to DB")
	fmt.Println("Serving...")
	//Returns the DB instance to Main file - Go_MVC.go
	return db
}
