package main
import("fmt")

func main()  {
	const firstName string = "haikal"
	const lastName = "akbar"

	fmt.Println("Hallo saya ", firstName, lastName)

	//multiple constant
	const (
		namaDepan = "haikal"
		namaBelakang = "akbar"
	)
	fmt.Println("Hallo saya ", namaDepan, namaBelakang)
}

