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
**Matematical operators:**
- \+ -> addition
- \- -> subtraction
- \* -> moltiplication
- / -> division
- % -> modulus , returns the remainder of a division operation 
- ++ -> increment 
- \-- -> decrement

**Logical operators**

- && -> Logical And -> (5 > x && x > 3) 
- || -> Logical Or  -> (5 > x || x > 3)
- !  -> Logical Not -> !(5 < x && x < 3)


### **FMT MODULE**
In Go, fmt is a package that provides functionality for formatting and printing output to the console. You can use fmt.Println() to print a new line to the console, fmt.Printf() to print formatted output.

                package main
                import "fmt"
                func main() {
                    fmt.Println("Hello, world!")
                    fmt.Printf("Hello, world!\n")
                }
#### **FORMAT VERBS**

With Fmt you can print formatted lines with Printf.
When we use printf, we can combine text to be printed with variable values, and we can choose how to format the variable value and consequently how to display it.
Here an example:

                fmt.Printf("this is your variable: %v", variable)

As you can see, we used the notation %v\n. The character \n only allows to line break, while %v is the notation to print the value of a variable in normal format.

**Other general formatting verbs are:**

- %#v	a Go-syntax representation of the value
- %T	a Go-syntax representation of the type of the value
- %%	a literal percent sign; consumes no value

**Format for integer values:**

- %b	base 2 
- %c	the character represented by the corresponding Unicode code point
- %d	base 10  
- %o	base 8  

**Format for Float values:**

- %f     default width, default precision
- %9f    width 9, default precision
- %.2f   default width, precision 2
- %9.2f  width 9, precision 2
- %9.f   width 9, precision 0

These are just a few of the most commonly used format verbs. You can find others on the official documentation page of the fmt package.
[-> fmt Official Documentation Page](https://pkg.go.dev/fmt)
Write a function that takes a slice of integers and two indices as input and returns a slice containing all the elements from the input slice between the specified indices (inclusive).
## **EXERCISES**

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

Now, the question is: how can I remove an index from a slice?
Well, let's assume you want to remove the last index of the slice. In that case, you need to create another time the same slice,with a modified length.
Example:
               
                slice := []int{1,2,3}
                slice2 = slice[:2]

With this notation, slice2 takes the indexes of slice from index 0 up to index 2, but index 2 is not included. Therefore, we can say that it's as if slice2 takes from index 0 to index 1. 

## **EXERCISES 2**

- 1.Write a function that takes an array as input and returns a slice containing the same values as the array.
- 2.Write a function that takes two slices of float64 as input and returns a slice containing all the elements from both slices, in the order they appear.
- 3.Write a function that takes as input two slices and an integer value "x", and creates a third slice corresponding to the first slice plus a number of elements from the second slice equal to "x".

## **IF CONDITIONS**

In Go, like in any other programming language, the conditional if statement is used to execute blocks of code when a certain condition is met.

**SYNTAX**

The syntax for if statements is very simple.
Example:

        x := 7
        if x > 5 {
                fmt.Println("x is greater than 5")
        } else if x = 5 {
                fmt.Println("x is equal to 5")
        }else {
                fmt.Println("x is less than 5" )
        }

## **SWITCH STATEMENT**
Switch statements are specifically designed for comparing a single value against a set of possible values. This makes them a better choice when you have many conditions based on the same variable.

**SYNTAX**

The syntax for switch statement is the following

        x := 7
        switch x {
        case x > 5: 
                fmt.Println("x is greater than 5")
        case x < 5:
                fmt.Println("x is equal to 5")               
        case x = 5:
                fmt.Println("x is less than 5")
        default:
                fmt.Println("x doesn't exist")
        }

As you can see , in switch statement we can use the **default**.
The "default" in a switch statement in Go is a special case that is executed when none of the cases match the value of the control variable. It is similar to an else statement in an if-else statement. The default case is optional and can be included at the end of a switch block to handle all values that do not match any of the specified cases.

## FOR LOOPS
A for loop in Go is used to repeatedly execute a block of code until a specified condition is met. Here's a simple explanation with an example:

In a for loop, you typically have three components:

Initialization: This is where you initialize any variables that will be used in the loop.
Condition: This is the condition that is evaluated before each iteration of the loop. If the condition evaluates to true, the loop continues; if it evaluates to false, the loop terminates.
Post statement: This is where you update any variables or perform any actions after each iteration of the loop.

Here's an example of a simple for loop that prints the numbers from 1 to 5:

        for i := 1; i <= 5; i++ {
                fmt.Println(i) // Print the current value of i
    }

In this example:

- The loop initializes i to 1.
- It then checks if i is less than or equal to 5.
- If the condition is true, it executes the block of code inside the loop (in this case, printing the value of i).
- After each iteration, it increments i by 1.
- It continues this process until i is no longer less than or equal to 5.
