package main // Name of the package. Starting point to run the program

/*
 This is a preprocessor command which tells the Go compiler to include the
 files lying in the package fmt
*/
import (
	"fmt"
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
            Interface:
            Map:
            Channel:
        5- An integer literal can be a decimal, octal, or hexadecimal constant
            0x or 0X for hexadecimal, 0 for octal, and nothing for decimal
    */

    //Static Type Declaring
    var  i int
    // var  c, ch byte
    // var  f, salary float32
    // var d =  42
    fmt.Printf("x is of type %T\n", i)  // println fails here

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
    */

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

    // select {
    //     case c = <- a :
    //         fmt.Println("select clause")
    // }


}