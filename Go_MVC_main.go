//Every stand-alone go program must have a main package
package main

//Controller package has the "Mux_Register()"
import "Controller"
//Model package has the "Connect()"
import "Model"
//net/http - for http send/receive requests
import "net/http"
//Importing MYSQL driver to connect to SQL Database
import _ "github.com/go-sql-driver/mysql"
//To display log messages
import "log"

//Entry point for the program
func main() {

	//Creating a multiplexer/router to route data
	//Mux_register() is in file "Go_Route.go" in "Controller" package
	mux := Controller.Mux_Register()
	//Connecting to database
	//Connect() is in file "Go_DBConnect.go" in "Model" package
	SK7db := Model.Connect()
	//defer will make this line as the last line of execution in program
	defer SK7db.Close()
	//Http server made to listen in specified address waiting for request from client
	//Multiplexer to be performed when request is made from client side
	//log.Fatal() - If error exists, log the error. If no error, do nothing
	log.Fatal(http.ListenAndServe("localhost:3000", mux))
}