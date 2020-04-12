//This file is under Model package
package Model

//Views package has the "Response" structure
import "Views"

//Function to read all data from table
func ReadAll() ([]Views.PostRequest, error){

	//Read Query
	rows, err := con.Query("SELECT * FROM TODO")
	if err != nil{
		return nil, err
	}

	//Creating an array of type struct(PostRequest) to store its items
	todos := []Views.PostRequest{}
	for rows.Next(){
		//Creating an empty instance of PostRequest structure
		data := Views.PostRequest{}
		//Storing each row's data into corresponding items in PostRequest structure's empty instance
		rows.Scan(&data.Name, &data.Todo)
		//Appending to the array
		todos = append(todos, data)
	}

	//Returns table data as array of structure type(PostRequest)
	return todos, nil
}

//Function to read data from table by condition
func ReadById(name string) ([]Views.PostRequest, error){

	//Read Query with respect to "name" id
	rows, err := con.Query("SELECT * FROM TODO WHERE name = ?", name)
	if err != nil{
		return nil, err
	}

	//Creating an array of type struct(PostRequest) to store its items
	todos := []Views.PostRequest{}
	for rows.Next(){
		//Creating an empty instance of PostRequest structure
		data := Views.PostRequest{}
		//Storing each row's data into corresponding items in PostRequest structure's empty instance
		rows.Scan(&data.Name, &data.Todo)
		//Appending to the array
		todos = append(todos, data)
	}

	//Returns table data as array of structure type(PostRequest)
	return todos, nil
}