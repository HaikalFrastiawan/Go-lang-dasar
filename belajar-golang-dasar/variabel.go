package main
import("fmt")

func main()  {
	var name string

	name = "haikal"
	fmt.Println("Hallo saya ", name)

	var nama = "haikal"
	fmt.Println("Hallo saya ", nama)
	
	//deklarasi pertama
	iniNama :=	 "haikal"
	fmt.Println("Hallo saya ", iniNama)

	//multiple deklarasi
	var (
		firstName = "haikal"
		lastName  = "akbar"
	)
	fmt.Println("Hallo saya ", firstName, lastName)

}

