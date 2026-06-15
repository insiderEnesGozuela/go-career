package main

import (
	"fmt"
	"io"
	"os"
)

// io.Writer interface aldık, böylece nereye yazdığımızı kontrol edebiliyoruz
// hem test (bytes.Buffer) hem de gerçek uygulama (os.Stdout) ile çalışır
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func main() {
	Greet(os.Stdout, "Elodie")
}
