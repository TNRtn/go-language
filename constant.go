package main
import ("fmt")
const a int=1//typed constant
const b=2// untyped constant
//constant declared anywhare
func main(){
	const c int=3
	const d=4
	fmt.Println(a,b,c,d)
}