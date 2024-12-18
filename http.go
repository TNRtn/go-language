package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://6603e4ab14491e840e279106--incandescent-zuccutto-cc62bf.netlify.app/"

func main() {
	fmt.Println("Sending HTTP request...")
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close() 
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
