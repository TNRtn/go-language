package main
import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter,r *http.Request){
fmt.Fprintln(w)
fmt.Fprintln(w,"hello")
}
func main(){
	fmt.Println("server not yet started....")

	http.HandleFunc("/",home)
	err:=http.ListenAndServe(":8080",nil)
	if err!=nil{
		fmt.Println("server failed to start",err)
	}

}