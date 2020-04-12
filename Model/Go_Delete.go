//This file is under Model package
package Model

//Importing MYSQL driver to connect to SQL Database
//import _ "github.com/go-sql-driver/mysql"

//Function to delete a data from table in DB based on condition
func Deletion(name string) error{

	//Deletion Query
	deleteQ, err := con.Query("DELETE FROM TODO WHERE name=?", name)
	//Deletion Query closed at last
	defer deleteQ.Close()

	//Returns error if anything else returns nil
	if err != nil{
		return err
	}	
	return nil
}