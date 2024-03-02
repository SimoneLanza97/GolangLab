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
	fmt.Println(array1)
}