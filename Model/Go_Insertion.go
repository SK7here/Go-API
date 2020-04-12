//This file is under Model package
package Model

//Importing MYSQL driver to connect to SQL Database

//Function to insert a data into table in DB
func Insertion(name string, todo string) error{

	//Insertion Query
	insertQ, err := con.Query("INSERT INTO TODO VALUES(?, ?)", name, todo)
	//Insertion Query closed at last
	defer insertQ.Close()

	//Returns error if anything else returns nil
	if err != nil{
		return err
	}	
	return nil
}