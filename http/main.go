package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

// exploring net package
func main() {
	res, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// fmt.Println(res)
	// bs:=make([]byte,99999) //used to create a byte slice which has 99999 empty elements
	// res.Body.Read(bs) //from the interface read the byteslice and store it in bs
	// fmt.Println(string(bs))

	// io.Copy(os.Stdout, res.Body) //io.Copy(writer interface,reader interface) works same as the above 3 lines

	lw:=logWriter{}
	io.Copy(lw,res.Body)

}

func (logWriter) Write(bs []byte)(int,error){
		fmt.Println(string(bs))
		return len(bs),nil
}