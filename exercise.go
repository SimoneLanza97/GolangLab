package main 
import ( 
	"fmt"
)
func som(a, b int) int {return a + b}
func diff(a, b int) int {return a - b}
func molt(a, b int) int {return a * b}
func div(a, b int) int {return a / b}
func somf(a, b float64) float64 {return a + b}
func difff(a, b float64) float64 {return a - b}
func moltf(a, b float64) float64 {return a * b}
func divf(a, b float64) float64 {return float64(a) / float64(b)}


func main() {
    // Dichiarazione di variabili int
    var intVar1 int = 10
    var intVar2 int = 20
    var intVar3 int = 30
    var intVar4 int = 0
    var intVar5 int = 100
    var intVar6 int = 50
    var intVar7 int = 75
    var intVar8 int = 42

    // Dichiarazione di variabili float64
    var floatVar1 float64 = 323.14
    var floatVar2 float64 = 5.67
    var floatVar3 float64 = 100.0
    var floatVar4 float64 = 12.0
    var floatVar5 float64 = 123.456
    var floatVar6 float64 = 789.123
    var floatVar7 float64 = 25.75
    var floatVar8 float64 = 8.99

	// var array1 [5]int = [5]int{1,2,3,4,5,} 

	ex1 := som(intVar1, intVar2)
	ex2 := diff(intVar3, intVar4)
	ex3 := molt(intVar5, intVar6)
	ex4 := div(intVar7, intVar8)
    fmt.Println("results of integer operation:")
	fmt.Printf("%v\n%v\n%v\n%v\n",ex1,ex2,ex3,ex4)

    ex5 := somf(floatVar1,floatVar2)
    ex6 := difff(floatVar3,floatVar4)
    ex7 := moltf(floatVar5,floatVar6)
    ex8 := divf(floatVar7,floatVar8)
    fmt.Println("results of float operation:")
    fmt.Printf("%.2f\n%.2f\n%.2f\n%.2f\n",ex5,ex6,ex7,ex8)
}
