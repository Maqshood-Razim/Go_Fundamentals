package main

import (
	"fmt"
	"net/url"
	"time"
)

const Myurl = "https://lco.dev:3000/learning?coursename=Go"

func main() {
	fmt.Println("url")
	fmt.Println(Myurl)

	//parse the url

	result, err := url.Parse(Myurl)

	if err != nil {
		panic(err)
	}

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	// fmt.Println(result.Port())
	qparams := result.Query()

	fmt.Printf("type of query %T \n", qparams)

	fmt.Println(qparams["coursename"])

	for _, val := range qparams {
		fmt.Println("param is :", val)
	}

	partsurl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}

	newurl := partsurl.String()
	fmt.Println(newurl)
  
	go greeter("hello")
	greeter("world")

  

}


func greeter(s string)  {
	for i := 0; i < 5; i++ {
		time.Sleep(7* time.Millisecond)
	   fmt.Println(s)
	}
 }
