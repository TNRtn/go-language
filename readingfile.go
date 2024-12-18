package main
import (
	"fmt"
	"os"
	"io"
)
func main(){
	file,err:=os.Open("./tnr.txt")
	if err!=nil{
		panic(err)
	}
	defer file.Close()
	content,err:=io.ReadAll(file)
	if err!=nil{
		panic(err)
	}
	fmt.Println(content)
	fmt.Println(string(content))

}