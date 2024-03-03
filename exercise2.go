package main

import (
	"fmt"
)
func es_1(x [4]string) []string {var slice1 []string = x[:];return slice1}
func es_2(x,y []float64) []float64 {var slice1 = append(x,y...);return slice1}
func es_3(x []int, y []int, z int) []int {var slice1 = append(x,y[:z]...);return slice1}

func main() {
	// for es_1
	var input1 [4]string = [4]string{"George","Steven","Albert","Simon"}
	// for es_2
	var input2 []float64 = []float64{0.3,5.2,7.9,3.21}
	var input3 []float64 = []float64{5.0,9.2,88.6,46.5,17.111}
	// for es_3
	var input4 []int = []int{555,78,56,5341,48}
	var input5 []int = []int{964,54684,15,326,45}
	var input6 int = 3
	
	result1 := es_1(input1)
	fmt.Println(result1)
	
	result2 := es_2(input2,input3)
	fmt.Printf("%.3f\n",result2)

	result3 := es_3(input4,input5,input6)
	fmt.Printf("%v\n",result3)
}
