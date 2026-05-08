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
	_, err := write.WriteString("asdasdasd \n")
	write.Flush()
	if err != nil {
		fmt.Println(err.Error())
	}

	file, err := os.ReadFile("file.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(file))
}
