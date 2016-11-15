package main

import "fmt"

//iota generates increamenting values to all Constants.
const (
	first = iota + 1
	second
	third
)

func main() {
	myArray()
	myMap()

}

//trying Array and slice  in GO. slice more flexible than Array.
func myArray() {
	arr := [0]int{}                       //intailizing Array
	sli := arr[:]                         //intailizing slice
	sli = append(sli, 100, 200, 900, 827) //append(arrayname or slice name,eliment1,element2,......)
	fmt.Println(sli)

}

//trying map Data Structure.
func myMap() {
	m := map[string]string{} //initalizing map
	m["a"] = "apple"
	m["b"] = "Banana"
	for k, v := range m {
		fmt.Println("Key:", k, "value:", v)
	}
	i, tpresent := m["c"] //i will get value from map with "c" ,In this case it's null. Second variable will get boolean if is persent or not.
	fmt.Println(i, tpresent)
}
func arthmatic() {

}
