package main
import "fmt"
func main(){
	// declaration of the pointer variable 
	var ptr *string

	//initialise variable
	greet := "Hello"

	// assign greet address to pointer
	ptr = &greet
	//printing
	fmt.Println("greet:",greet)
	fmt.Println("address of greet;",ptr)
	fmt.Println("Value stored in ptr:",*ptr)
}