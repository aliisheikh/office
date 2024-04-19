package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	write, err := os.Create("File1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer write.Close()

	read := strings.NewReader("Muhammad Ali")
	io.Copy(write, read)

}
