package main

var input string

func main() {	
	database := connDB()
	routes()
	
	/* fmt.Print("ingrese: ")	
	fmt.Scanf("%s", &input)	
	*/

	// close the connection to the DB
	defer database.Close() 	
}
