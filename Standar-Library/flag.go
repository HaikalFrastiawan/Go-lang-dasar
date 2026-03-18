package main

import (
	"flag"
	"fmt"
)

func main(){
	var username *string = flag.String("username", "root", "database username")
	var password *string = flag.String("password", "root", "database password")
	var host *string = flag.String("host", "localhost", "database host")
	var port *int = flag.Int("port", 0 , "database port")

	flag.Parse()

	fmt.Println("username", *username)
	fmt.Println("password", *password)
	fmt.Println("host", *host)
	fmt.Println("port", *port)


	//masukin di terminalnya
	// go run flag.go -username=haikal -password=123 -host=123.456.789 -port=8080
}