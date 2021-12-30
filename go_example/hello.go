package main // Name of the package. Starting point to run the program

/*
 This is a preprocessor command which tells the Go compiler to include the
 files lying in the package fmt
*/
import (
	"fmt"
	"math"
	"strings"

	"rsc.io/quote"
)

/*
 Go program consists of various tokens.
 A token is either a keyword, an identifier, a constant, a string literal, or a symbol
*/

func main() {
    /* Reserved keywords
    break	default	func	interface	select
    case	defer	Go	map	Struct
    chan	else	Goto	package	Switch
    const	fallthrough	if	range	Type
    continue	for	import	return	Var
    */
    /*
    Notice the capital P of Println method.
     In Go language, a name is exported if it starts with capital letter.
     Exported means the function or variable/constant is accessible to the importer of the respective package.
    */
    fmt.Println(quote.Go())
    /*
    A Go identifier is a name used to identify a variable, function, or any other user-defined item. 
    An identifier starts with a letter A to Z or a to z or an underscore _ followed by zero or more letters, underscores, and digits (0 to 9).
    Go is a case-sensitive programming language.
    */

    /*
    The type of a variable determines how much space it occupies in storage and how the bit pattern stored is
     interpreted:
        1- Boolean : true or false
        2- Numeric: integer or floats
        3- string: immutable sequence of bytes.
            When certain characters are preceded by a backslash, they will have a special meaning in Go.
            \?	? character
            \a	Alert or bell
            \f	Form feed
            \v	Vertical tab
            \ooo	Octal number of one to three digits
            \xhh . . .	Hexadecimal number of one or more digits
        4- Devired: 
            Pointer:
            Array (aggregate type):
            Structure (aggregate type):
            Union:
            Function:
            Slice: 
                strings are slices
                Libraries to manipulate slices:
                    unicode
                    regexp
                    strings
            Interface:
            Map:
            Channel:
        5- An integer literal can be a decimal, octal, or hexadecimal constant
            0x or 0X for hexadecimal, 0 for octal, and nothing for decimal

    Variables can be defined:
        Inside a function or a block (local variables)
        Outside of all functions (global variables)
        In the definition of function parameters (formal parameters)
    
    Local and global variables are initialized to their default value, which is 0; while pointers are initialized to nil.
    */

    //Static Type Declaring
    numbers := [6]int{1, 2, 3, 4}  // position 4 (5th) and 5  (6th) filled with 0s 
    two_dimension_array := [3][4]int{  
        {0, 1, 2, 3},   /*  initializers for row indexed by 0 */
        {4, 5, 6, 7},   /*  initializers for row indexed by 1 */
        {8, 9, 10, 11}, /*  initializers for row indexed by 2 */
     }
        /* output each array element's value */
    for  i, inner_array := range two_dimension_array {
        for j, element := range inner_array {
        fmt.Printf("two_dimension_array[%d][%d] = %d\n", i,j, element )
        }
    }
    
    // var  c, ch byte
    // var  f, salary float32
    // var d =  42
    fmt.Printf("x is of type %T\n", numbers)  // println fails here
    fmt.Printf("Access to 3er element in array %d\n", numbers[3])
    fmt.Printf("Print elements on list %d\n", numbers)
    // Dynamic Type Declaring
    y := 42.0 
    fmt.Printf("x is of type %T\n", y)

    // Variables of different types can be declared in one go using type inference.
    var a, b, c = 3, 4, "foo"
    fmt.Printf("a is of type %T\n", a)
    fmt.Printf("b is of type %T\n", b)
    fmt.Printf("c is of type %T\n", c)
    
    // Constants
    const FIXED_TEXT string = "Something"
    fmt.Printf("const: %s", FIXED_TEXT)

    /*Prints Println fails to insert values
    "const: %S" --> const: %!S(string=Something)
    "const: %S" --> const: Something
    q flag escapes unprintable characters, with + flag it escapses non-ascii 
    characters as well to make output unambigous
    */

    for i := 0; i < len(FIXED_TEXT); i++ {
        fmt.Printf("%x ", FIXED_TEXT[i])
    }
    // TODO: Find out how to iterate over a list without index? is it possible -> for ch in some_string
    fmt.Printf("\n")
    const sampleText = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98" 
    fmt.Printf("quoted string: ")
    fmt.Printf("%+q", sampleText)
    fmt.Printf("\n") 

    // String libraries

    greetings :=  []string{"Hello","world!"}   
    fmt.Println(strings.Join(greetings, " "))


    /*Operators

    All teh same as in Python
    New ones (Miscellaneous):
        &	Returns the address of a variable.	&a; provides actual address of the variable.
        *	Pointer to a variable.	*a; provides pointer to a variable.
    
    Category	Operator	Associativity
    Postfix	() [] -> . ++ - -	Left to right
    Unary	+ - ! ~ ++ - - (type)* & sizeof	Right to left
    Multiplicative	* / %	Left to right
    Additive	+ -	Left to right
    Shift	<< >>	Left to right
    Relational	< <= > >=	Left to right
    Equality	== !=	Left to right
    Bitwise AND	&	Left to right
    Bitwise XOR	^	Left to right
    Bitwise OR	|	Left to right
    Logical AND	&&	Left to right
    Logical OR	||	Left to right
    Assignment	= += -= *= /= %=>>= <<= &= ^= |=	Right to left
    Comma	,	Left to right

    */

    if FIXED_TEXT == "something" {
        fmt.Println("Nested if parent")
        if c == "foo" {
            fmt.Println("Nested if chilf")
        }
    } else if FIXED_TEXT == "Something" {
        fmt.Println("Nested if parent")
        if c == "foo" {
            fmt.Println("Nested if chilf")
        }
    }

    switch FIXED_TEXT == "Something" {
    case true:
        fmt.Println("Switch case  true")
        
    case false:
        fmt.Println("Switch case false")
        
    }

    /*
    A select statement is similar to switch 
    statement with difference that case statements refers to channel communications

    TODO: What is a channel communication in Go??
    */

    /* for loop execution 
    You can nest them

    break, continue available

    goto transfers control to the labaled statement

    */
    for i, x:= range numbers {
        fmt.Printf("value of x = %d at %d\n", x,i)
     }  

    b = 15
    a = 0
    for a := 0; a < 10; a++ {
        fmt.Printf("value of a: %d\n", a)
        }
    for a < b {
        a++
        fmt.Printf("value of a: %d\n", a)
    }
    // Using labels and goto keyword
    a = 10
    for a < 20 {
        if a == 15 {
           /* skip the iteration */
           a += 1
           goto LOOP
        }
        fmt.Printf("value of a: %d\n", a)
        a++
             
    } 
    fmt.Println("This is never executed")
    LOOP: fmt.Println("Loop exits for good")
    fmt.Println("Second message")

    /*infinite loop
     
    for true  {
       fmt.Printf("This loop will run forever.\n");
    }
    */

    /* Functions
    The formal parameters behave like other local variables
    inside the function and are created upon entry into the function and destroyed upon exit.

        The call by value method of passing arguments to a function copies the actual value of 
    an argument into the formal parameter of the function. In this case, changes made to the
     parameter inside the function have no effect on the argument.
    
    */
    var res = cast_to_float(15)
    fmt.Printf("value of a: %f\n", res)

    var1, var2 := swap_by_value("first", "second")
    fmt.Println(var1, var2)

    /* 
        The call by reference method of passing arguments to a function copies the address of
     an argument into the formal parameter. Inside the function, the address is used to 
     access the actual argument used in the call. This means that changes made to the 
     parameter affect the passed argument.
    
     */
    var str_by_ref_1 string = "100"
    var str_by_ref_2 string = "200"
    swap_by_ref(&str_by_ref_1, &str_by_ref_2)

    fmt.Printf("ref1 " + str_by_ref_1 + " ref2 " + str_by_ref_2 + "\n")


    /*
    Go programming language provides the flexibility to create functions on the fly and use them as values
    */
    /* declare a function variable */
    getSquareRoot := func(x float64) float64 {
        return math.Sqrt(x)
    }
    /* use the function */
    fmt.Println(getSquareRoot(9))

    /*
    Go programming language supports anonymous functions which can acts as function closures. 
    Anonymous functions are used in dynamic programming when we want to define a function 
    inline without passing any name to it.
    */
    /* nextNumber is now a function with i as 0 */
    nextNumber := getSequence()  

    /* invoke nextNumber to increase i by 1 and return the same */
    fmt.Println(nextNumber())
    fmt.Println(nextNumber())
    /* create a new sequence and see the result, i is 0 again*/
    nextNumber1 := getSequence()  
    fmt.Println(nextNumber1())
    fmt.Println(nextNumber()) // should be 3, not 2

    /*
    Go programming language supports special types of functions called methods.
     In method declaration syntax, a "receiver" is present to represent the 
     container of the function. This receiver can be used to call a function using "." operator.
    */

    circle := Circle{x:0, y:0, radius:5} // Positional arguments also valid (not mixed options available)
    fmt.Printf("Circle area: %f\n", circle.area())
    // To print structs better https://gosamples.dev/print-struct-variables/
    fmt.Printf("%+v\n%#v\n", circle, circle)
    //cannot use numbers (variable of type [6]int) as []int value in argument to getAverage
    var  balance = []int {1000, 2, 3, 17, 50}
    fmt.Printf("Average of array %d = %f\n", balance, getAverage(balance))



}

/*Casting to float64 the hard way*/
func cast_to_float( a int) float64 {
    return float64(1.0 * a)
}

func swap_by_value(x, y string) (string, string) {
    return y, x
 }

 func swap_by_ref(x *string, y *string) {
     //No return required
    temp := *x    /* save the value at address x */
    *x = *y      /* put y into x */
    *y = temp    /* put temp into y */
 }

 func getSequence() func() int {
     /*
     we created a function getSequence() which returns another function. 
     The purpose of this function is to close over a variable i of upper function to form a closure.
     */
    i:=0
    return func() int {
       i+=1
       return i  
    }
 }

 /* define a circle For a method demo */
type Circle struct {
    x,y,radius float64
 }
 
 /* define a method for circle */
 func(circle Circle) area() float64 {
    return math.Pi * circle.radius * circle.radius
 }

 func getAverage(my_list []int) float32 {
    var sum int  
 
    for _, item := range my_list {
       sum += item
    }
 
    return float32(sum / len(my_list))
 }