//This file is under Controller package
package Controller

//Views package has the "Response" structure
import "Views"
//Model package has the "Insertion()"
import "Model"
//fmt package - to print out statements
import "fmt"
//encoding/json - For JSON encoding of custom data types
import "encoding/json"
//net/http - for http send/receive requests
import "net/http"

//This function returns the default response
func Home_Function() http.HandlerFunc{

	return func(w http.ResponseWriter, r *http.Request) {
		//This case works only if the request is of type GET
		if r.Method == http.MethodGet{
			//Printed in server terminal after query processed
			fmt.Println("Request received")
			//Initializing values for items in structure "Response" which is in "Go_Structs.go" under Views package
			data := Views.Response{
				//Value corresponding to StatusOK(200) will be assigned
				Code : http.StatusOK,
				Body : "Response",
			}
			//Outputting statusOK(200) - for GET request successfully processed
			w.WriteHeader(http.StatusOK)
			//JSON Encoder is wrapped around ResponseWriter and data is encoded into JSON
			json.NewEncoder(w).Encode(data)
		}
		if r.Method != http.MethodGet{
			//Printed in Client side stating error
			w.Write([]byte("Some error (Try with GET method)"))
		}
	}
}

//This function returns repsonse for Insertion Query
func Insertion_Function() http.HandlerFunc{

	return func(w http.ResponseWriter, r *http.Request) {
		//This case works only if the request is of type POST
		if r.Method == http.MethodPost{
			//Creating an empty instance of PostRequest structure
			data := Views.PostRequest{}
			//JSON Decoder is wrapped around body of request and data is decoded from JSON
			json.NewDecoder(r.Body).Decode(&data)
			//Printing the request data
			fmt.Println(data)
			//Getting return from Insertion()
			err := Model.Insertion(data.Name, data.Todo)
			if err != nil{
				//Printed in Client side stating error
				w.Write([]byte("Some error"))
				return
			}
			//Outputting statusCretaed(201) - for POST request successfully processed
			w.WriteHeader(http.StatusCreated)
			//Printed in client terminal after query processed
			//Status is sent as a structure which is encoded as json
			//JSON Encoder is wrapped around ResponseWriter and data is encoded into JSON
			json.NewEncoder(w).Encode(struct{
				Status string `json : "status"`
				}{Status : "Insertion Succesful"})
			//Printed in server terminal after query processed
			fmt.Println("Insertion successful")
		}
		if r.Method != http.MethodPost{
			//Printed in Client side stating error
			w.Write([]byte("Some error (Try with POST method)"))
		}
	}
}

func ReadAll_Function() http.HandlerFunc{

	return func(w http.ResponseWriter, r *http.Request) {
		//This case works only if the request is of type GET
		if r.Method == http.MethodGet{
			//Calling the ReadAll query
			data, err := Model.ReadAll()
			if err != nil{
				//Printed in Client side stating error
				w.Write([]byte("Some error"))
				return
			}
			//Outputting statusOK(200) - for GET request successfully processed
			w.WriteHeader(http.StatusOK)
			//JSON Encoder is wrapped around ResponseWriter and data is encoded into JSON
			json.NewEncoder(w).Encode(data)
			//Printed in server terminal after query processed
			fmt.Println("Reading successful")
		}
		if r.Method != http.MethodGet{
			//Printed in Client side stating error
			w.Write([]byte("Some error (Try with GET method)"))
		}
	}
}

func ReadById_Function() http.HandlerFunc{

	return func(w http.ResponseWriter, r *http.Request) {
		//This case works only if the request is of type GET
		if r.Method == http.MethodGet{
			//Getting the value of name from URL Query
			name := r.URL.Query().Get("name")
			//Calling the ReadById Query
			data, err := Model.ReadById(name)
			if err != nil{
				//Printed in Client side stating error
				w.Write([]byte("Some error"))
				return
			}
			//Outputting statusOK(200) - for GET request successfully processed
			w.WriteHeader(http.StatusOK)
			//JSON Encoder is wrapped around ResponseWriter and data is encoded into JSON
			json.NewEncoder(w).Encode(data)
			//Printed in server terminal after query processed
			fmt.Println("Reading by Name successful")
		}
		if r.Method != http.MethodGet{
			//Printed in Client side stating error
			w.Write([]byte("Some error (Try with GET method)"))
		}
	}
}

func Deletion_Function() http.HandlerFunc{

	return func(w http.ResponseWriter, r *http.Request) {
		//This case works only if the request is of type GET
		if r.Method == http.MethodDelete{
			//Getting the value of name from URL path
			//This extracts name value from url after /SK7server/delete/...
			name := r.URL.Path[18:]

			err := Model.Deletion(name)
			if err != nil{
				//Printed in Client side stating error
				w.Write([]byte("Some error"))
				return				
			}
			//Outputting statusOK(200) - for GET request successfully processed
			w.WriteHeader(http.StatusOK)
			//Printed in client terminal after query processed
			//Status is sent as a structure which is encoded as json
			//JSON Encoder is wrapped around ResponseWriter and data is encoded into JSON
			json.NewEncoder(w).Encode(struct{
				Status string `json : "status"`
				}{Status : "Deletion Succesful"})
			//Printed in server terminal after query processed
			fmt.Println("Deletion successful")
		}
		if r.Method != http.MethodDelete{
			//Printed in Client side stating error
			w.Write([]byte("Some error (Try with DELETE method)"))
		}
	}
}



