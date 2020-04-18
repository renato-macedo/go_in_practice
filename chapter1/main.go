package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func readStatus() {
	// Dial function connects using the type and endpoint specified
	conn, _ := net.Dial("tcp", "golang.org:80")
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")

	status, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(status)
}

func httpGET() {
	resp, err := http.Get("http://example.com")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func count() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 1)
	}
}

func printConcurrently() {
	go count()
	time.Sleep(time.Millisecond * 2)
	fmt.Println("Hello")
	time.Sleep(time.Millisecond * 5)
}

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Hello, my name is Inigo Montoya")
}
func inigo() {
	http.HandleFunc("/", hello)
	http.ListenAndServe("localhost:4000", nil)
}

func main() {
	// readStatus()
	// httpGET()
	// printConcurrently()
	inigo()
}
