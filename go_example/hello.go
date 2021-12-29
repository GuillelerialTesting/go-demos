package main  // Name of the package. Starting point to run the program

 /*
 This is a preprocessor command which tells the Go compiler to include the 
 files lying in the package fmt
 */
import "fmt"

/*
 Go program consists of various tokens. 
 A token is either a keyword, an identifier, a constant, a string literal, or a symbol
*/
import "rsc.io/quote"

func main() {
    /*
    Notice the capital P of Println method.
     In Go language, a name is exported if it starts with capital letter.
     Exported means the function or variable/constant is accessible to the importer of the respective package.
    */
    fmt.Println(quote.Go())

}