package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	const myurl = "https://dsa101.onrender.com/get-chapters"

	res, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	fmt.Println("Status code", res.StatusCode)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println(string(body))
}
