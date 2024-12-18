package main
import ("fmt")
func main(){
	defer fmt.Println("world")
	//dfer stop execting that line and excute just befor 
	//end curly bracket
	fmt.Println("hello")
}