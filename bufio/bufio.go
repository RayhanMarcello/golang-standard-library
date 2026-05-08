package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// io hanya berisi kontrak (interface) jadi haru implement sendiri
	// io (Reader & Writer)

	// bufio(buffered io) -> methode yang dapat digunakan
	// bufio (NewReader & NewWriter)

	input := strings.NewReader("rayhan marcello \n ananda p ")
	reader := bufio.NewReader(input)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		fmt.Println(string(line))
	}

	// write
	write := bufio.NewWriter(os.Stdout)
	write.WriteString("rayhan asdasdasd")
	write.Flush()

}
