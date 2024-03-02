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

## **SLICES**

What is a sclice? 
So , in go a slice is like an array whit undecleared index, like this 

        slice := []string{"i'm" "an array" "too"}

The difference between array and slices is that you can change the dimension of a slice.
An array can contain a limited number of values, determined by its index, while a slice, not having a declared index, can be changed; we can add more values or remove them.

### **how to canghe arrays and slices values ?**

You can modify a value in the array/slice simply declaring a new value for that index , like this: 

        array := [2]int{1,2}
        array[0] = 44
If you print the value of your array on the screen, it will be [44 2] and not [1 2] because you changed the index 0 from its previous value to 44.

## OPERATORS 
Matematichs operators:
- \+ -> addition
- \- -> subtraction
- \* -> moltiplication
- / -> division
- % -> modulus , returns the remainder of a division operation 
- ++ -> increment 
- \-- -> decrement



