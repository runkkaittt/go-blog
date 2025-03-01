package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func main() {
	for range 3 {
		newUser()
	}
	// updateUser()
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

// func updateUser() {
// 	data := []byte(`{"name": "Dima"}`)
// 	body := bytes.NewReader(data)

// 	req, err := http.NewRequest("PUT", "http://localhost:8080/user", body)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	req.Header.Add("Content-Type", "application/json")

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer resp.Body.Close()

// 	fmt.Println("Запрос отправлен")

// 	d, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	}
// 	fmt.Println(string(d))
// }
