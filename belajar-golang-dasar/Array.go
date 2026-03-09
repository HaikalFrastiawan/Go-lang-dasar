package main
import ("fmt")

func main(){
	var names [3] string
	names[0] = "haikal"
	names[1] = "fras"
	names[2] = "tiawan"

	var values = [...]int{1,2,3}


	
	fmt.Println(values)
	fmt.Println(len(values))
	fmt.Println(names)
}