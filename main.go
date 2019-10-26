package main

import "fmt"

type person struct { //custom type
	fname string
	lname string
}

func main() {

	fmt.Println("Hello Go")

	slices := []int{4, 5, 6, 7, 8} //slices
	fmt.Println(slices)

	m := map[string]int{ //map
		"Raj": 23,
		"Adi": 24,
	}
	fmt.Println(m)

	p1 := person{

		"nafisur", "ahmed",
	}
	fmt.Println(p1)

}
