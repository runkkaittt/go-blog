package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func main() {
	for range 3 {
		newUser()
	}
}

func newUser() {
	data := []byte(`{"name": "Eric"}`)
	r := bytes.NewReader(data)

	resp, err := http.Post("http://localhost:8080/user", "application/json", r)
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer resp.Body.Close()

	fmt.Println("Запрос отправлен")

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(string(body))
}

func updateUser() {
	data := []byte(`{"name": "Eric"}`)
	r := bytes.NewReader(data)

	resp, err := http.Post("http://localhost:8080/user/"+strconv.Itoa(0), "application/json", r)
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer resp.Body.Close()

	fmt.Println("Запрос отправлен")

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(string(body))
}
