package main
import ("fmt")

func main(){
	names := [...] string{"hai","kal","fras","tiaewan","alvin", "akbar"}
	fmt.Println(names)

	slice1 := names[4:6]
	fmt.Println(slice1)

	var slice3 []string = names[3:]
	fmt.Println(slice3)


	days := [...]string{"senin","selasa","rabu","kamis","jumat","sabtu","minggu"}
	fmt.Println(days)
	
	daySlice1 := days[5:]//sabtu, minggu
	fmt.Println(daySlice1)

	daySlice1[0] = "sabtu Baru"
	daySlice1[1] = "minggu Baru"
	fmt.Println(daySlice1)
	fmt.Println(days)

	daySlice2:= append(daySlice1, "Libur Baru")
	daySlice2[0] = "sabtu lama"
	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	fmt.Println(days)

	var newSlice []string = make([]string, 2, 5)
	newSlice[0] = "Haikal"
	newSlice[1] = "Haikal"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))
	
	//untuk menambahkan array harus pakai append 
	newSlice2 := append(newSlice, "frass")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	//copy slice
	fromSlice := days[:]
	toSlice := make([]string,len(fromSlice),cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)


	iniArray := [...]int{1,2,3}
	iniSlice := []int{1,2,3,4}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)


	//latihan slice

	Angka := [...]int{10,20,30,40,50,60}
	fmt.Println(Angka)

	sliceA := Angka[3:]
	sliceA[0] = 99

	fmt.Println(sliceA)

}