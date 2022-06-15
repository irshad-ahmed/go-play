package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Wrote ", len(bs), " bytes")
	return len(bs), nil
}
func main() {
	resp, err := http.Get("https://www.google.com")
	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(1)
	}
	/*
		bs := make([]byte, 99999)
		resp.Body.Read(bs)
		fmt.Println(string(bs), len(bs)) */

	//io.Copy(os.Stdout, resp.Body)

	lw := logWriter{}
	io.Copy(lw, resp.Body)

}
