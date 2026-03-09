package main
import("fmt")

func main(){
	//versi panjang
	// var person map[string]string = map[string]string{}
	// person["name"] = "Haikal"
	// person["Adress"] = "subang"
	
	person := map[string]string {
		"name" : "Haikal",
		"address": "jakarta",
	}

	fmt.Println(person["name"])
	fmt.Println(person["address"])

	book := make(map[string]string)
	book["title"] ="Buku Golang"
	book["author"] = "Haikal"
	book["ups"] = "Salah"

	fmt.Println(book)
	delete(book, "ups")
	fmt.Println(book)


	//latihan map
	inventory := map[string]int{
		"Buku" : 10,
		"Pena" : 25,
		"Penghapus" : 5,
	}

	fmt.Println(inventory)

	//menambahkan buku = 5pcs
	inventory["Buku"] = inventory["Buku"] + 5
	fmt.Println(inventory)

	//latihan map 2
	users := map[string]string{
		"admin" : "12345",
	}

	

	//cara comma ok idiom
	value, ok := users["guest"]

	if ok{
		fmt.Println("Data ditemukan", value)
	} else {
		fmt.Println("data tidak di temukan")
	}

	password, found := users["admin"]
	if found {
		fmt.Println("admin password",password)
	}	
}