package main
import ("fmt")
func main(){
	type NoKtp string

	var ktpEkal string = "1234"
	
	var contoh string = "3456"
	var contohktp NoKtp = NoKtp(contoh)

	fmt.Println(ktpEkal)
	fmt.Println(contohktp)
}