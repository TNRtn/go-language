package main
import ("fmt"
"os"
"io")
func main(){
	fmt.Println("files")
	content:="tnr from srm"
	file,err:=os.Create("./tnr.txt")
	if err!=nil {
		panic(err)//stop the execution and show the error
	}
	length,err:=io.WriteString(file,content)
	if err!=nil{
		panic(err)
	}
	fmt.Println("length is :",length)
	file.Close()
}