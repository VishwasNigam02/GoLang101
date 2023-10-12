package main
import "fmt"
func negate(x *int){
	neg := -*x
	*x = neg
}
func main(){
	x := 3
	fmt.Println(x) //3
	negate(&x)
	fmt.Println(x) //-3
}