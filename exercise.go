package main 
import ( 
	"fmt"
)
func main() {
	array := [2]int{1,2}
	fmt.Println("this is your array", array)
	array[0] = 44
	fmt.Println("This is your array after changes: ", array)
	array1 := [5]int{1:23,2:44}
	length := len(array1)
	capacity := cap(array1)
	fmt.Println("length=",length,"capacity=", capacity)
	fmt.Println(array1)
	slice := []int{1,2,3}
	fmt.Println("your slice is :", slice)
	slice = append(slice,23,24)
	fmt.Println("your slice is :", slice)
	slice1 := []int{55,33}
	slice = append(slice, slice1...)
	fmt.Println("your slice is :", slice)
	slice2 := slice[:2]
	fmt.Println("your slice2 is :", slice2)


}