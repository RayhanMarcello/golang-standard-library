package main

import (
	"encoding/base64"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

type Users struct {
	Name string
	Age  int
}

func main() {
	// Encode = mengubah data dari bentuk asli (struct/object) menjadi format tertentu.
	// Decode = mengubah kembali data dari format tertentu menjadi bentuk asli.

	// contoh encode ke json
	x := Users{
		Name: "rayhan",
		Age:  20,
	}
	jsonEncode, err := json.Marshal(x)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(jsonEncode))

	var userDecode Users
	err = json.Unmarshal(jsonEncode, &userDecode)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(userDecode.Name, userDecode.Age)

	// contoh encode ke base64
	var encodeBase64 = base64.StdEncoding.EncodeToString([]byte("rayhan marcello"))
	fmt.Println(encodeBase64)

	decodeBase64, err := base64.StdEncoding.DecodeString(encodeBase64)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(decodeBase64))

	// contoh csv reader
	csvString := "rayhan,marcello,ananda\n" + "rayhan,marcello,ananda\n" + "rayhan,marcello,ananda\n"
	reader := csv.NewReader((strings.NewReader(csvString)))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(record)
	}
}
