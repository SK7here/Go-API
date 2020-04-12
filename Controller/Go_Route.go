//This file is under Controller package
package Controller


//net/http - for http send/receive requests
import "net/http"

//This function creates a multiplexer, prepares the response and returns the multiplexer
func Mux_Register() *http.ServeMux{

	//Creating a multiplexer/router to route data
	mux := http.NewServeMux()

	//Define the function to be preformed by multiplexer along with the invoker - /SK7server
	//Add '/' at end of trigger to enable base route 
	//Multiplexer will be created corresponding to type of url/query

	//"Home_Function()" is in file "Go_HandlerFunc.go" in "Controller" package
	mux.HandleFunc("/SK7server/", Home_Function())
	//"Insertion_Function()" is in file "Go_HandlerFunc.go" in "Controller" package
	mux.HandleFunc("/SK7server/insert/", Insertion_Function())
	//"ReadAll_Function()" is in file "Go_HandlerFunc.go" in "Controller" package
	mux.HandleFunc("/SK7server/readAll/", ReadAll_Function())
	//"ReadId_Function()" is in file "Go_HandlerFunc.go" in "Controller" package
	mux.HandleFunc("/SK7server/readbyid/", ReadById_Function())
	//"Deletion_Function()" is in file "Go_HandlerFunc.go" in "Controller" package
	mux.HandleFunc("/SK7server/delete/", Deletion_Function())
	return mux
}
