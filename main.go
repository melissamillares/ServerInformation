package main

func main() {	
	database := connDB()
	routes()

	// close the connection to the DB
	defer database.Close() 	
}
