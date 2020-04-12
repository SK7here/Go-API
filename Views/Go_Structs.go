//This file is under Views package
package Views

//Creating a structure(Custom type) which is to be encoded as JSON
type Response struct{

	//Tags are appended to each field to serialize into JSON format from structure
	Code int `json:"code"`
	Body interface{} `json:"body"`
}

type PostRequest struct{

	//Tags are appended to each field to deserialize from JSON format into structure
	Name string `json:"name"`
	Todo string `json:"todo"`
}
