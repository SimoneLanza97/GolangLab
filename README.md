# LEARN GO  
![](https://www.filepicker.io/api/file/O8dz87hXSheB05h3nO4M)
## **VARIABLES DECLARATION**
The first step in learning Go consists of understanding how Go handles the basics of programming, from syntax to the creation of conditions, loops, and functions.

In Go, a variable can be declared in two ways :
- the method where we declare the variable type

        var firstvar string = "this is a string"

- the method where the variable type will be inferred by Go itself:

        firstvar := "this is a string" 

You can use the method you prefer , but essentialy , it's the same thing , you are declaring a string type variable whit a string value.

As others programming language , in go there are 4 type of data :

- **data type** -> **how to declare it**

- strings -> int , int32 , int64 , uint , uint32 , uint64 
- integers -> string
- floats -> float32 , float64
- booleans -> bool 

## **ARRAYS**

In Go, an array is a collection of variables of the same type.
An array is declared in Go in the following way:

        var array [2]string = [2]string{"i'm", "an array"}

where:

- **var array** -> is the variables's type array declaration.
- **[2]** -> it's the index of the array and it determines how many values the array contains. in Go **you have to declare the index** of the array and you cannot change it.
- **string** -> is the type of data in the array.
- **{"i'm", "an array"}** -> is the value of the array , you must put the values **between curly braces** separated by a comma. 

You can also declare an array in this way :

        array := [2]string{"i'm" "an array"}

But keep in mind , if you declare an array of 5 values , declaring only some of these 5 values , the others will be 0.
Example:

        array := [5]int{1:23,2:44}
        // we are declaring an array with index 5 , but we are valorizing only the values at positions number 1 and 2.
        // if you print to the screen you array will be something like
        //[0 23 44 0 0]

### **how to canghe arrays values ?**

You can modify a value in the array simply declaring a new value for that index , like this: 

        array := [2]int{1,2}
        array[0] = 44
If you print the value of your array on the screen, it will be [44 2] and not [1 2] because you changed the index 0 from its previous value to 44.

### **how to print the length and capacity of an array**

You can use the build-in functions to see the length and capacity of an array:
- len(array) -> the length of an array (number of values)
- cap(array) -> the capacity of an array (max number of values)

this is not very useful with arrays because arrays have a predefined capacity and the length will always correspond to the capacity, but all this will be useful with slices, which resemble arrays but do not have a predefined length.

## **OPERATORS** 
Matematichs operators:
- \+ -> addition
- \- -> subtraction
- \* -> moltiplication
- / -> division
- % -> modulus , returns the remainder of a division operation 
- ++ -> increment 
- \-- -> decrement

## **EXERCISES**

### **FMT MODULE**
In Go, fmt is a package that provides functionality for formatting and printing output to the console. You can use fmt.Println() to print a new line to the console, fmt.Printf() to print formatted output.

                package main
                import "fmt"
                func main() {
                    fmt.Println("Hello, world!")
                    fmt.Printf("Hello, world!\n")
                }

Now you know how to print something on the screen and you might be able to do the next exercises:

- 1.Write a Go program that performs the following operations with integer numbers:
Addition of two integer numbers.
Subtraction of two integer numbers.
Multiplication of two integer numbers.
Division of two integer numbers.

- 2.Write a Go program that performs the following operations with floating-point numbers:
Addition of two floating-point numbers.
Subtraction of two floating-point numbers.
Multiplication of two floating-point numbers.
Division of two floating-point numbers.

- 3.Write a Go program that declares an array of 5 integer values, prints the array on the screen, changes the values at specific indices, and then prints the updated array on the screen:
Declare an array of 5 integers.
Initialize the array with some values.
Print the original array on the screen.
Change the values at specific indices as desired.
Print the updated array on the screen.


## **SLICES**

What is a sclice? 
So , in go a slice is like an array whit undecleared index, like this 

        slice := []string{"i'm" "an array" "too"}

The difference between array and slices is that you can change the dimension of a slice.
An array can contain a limited number of values, determined by its index, while a slice, not having a declared index, can be changed; we can add more values or remove them.

### **how to operate with slices?**

You can append values to a slice :

        slice := []int{1,2,3}
        slice = append(slice,23,24)

You can appen a slice to your slice , but slices have to be of contains the same data type.

        slice1 := []int{55,33}
        slice := append(slice, slice1)

Now, the question is: how the f**k can I remove an index from a slice?
Well, let's assume you want to remove the last index of the slice. In that case, you need to create another slice, possibly the same one, with a modified length.
Example:
                slice := []int{1,2,3}
                slice2 = slice[:2]

With this notation, slice2 takes the indexes of slice from index 0 up to index 2, but index 2 is not included. Therefore, we can say that it's as if slice2 takes from index 0 to index 1. 
