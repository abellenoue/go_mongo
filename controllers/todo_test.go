package controllers

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func TestCreateTodo(t *testing.T) {
	log.Printf("test")
	// newTodo := gin.H{
	// 	"Title":     "stp",
	// 	"Body":      "fonctionne",
	// 	"Completed": "test",
	// }

	// data := url.Values{}
	// data.Set("Title", "stp")
	// data.Set("surname", "bar")
	// u, error := url.ParseRequestURI("http://localhost:4747")
	// fmt.Println(error)
	// resource := "/todo/"
	// u.Path = resource
	// urlStr := u.String()

	// client := &http.Client{}
	// r, _ := http.NewRequest(http.MethodPost, urlStr, strings.NewReader(data.Encode()))
	// r.Header.Set("Content-Type", "application/json")
	// resp, _ := client.Do(r)
	// fmt.Println(resp.StatusCode)

	url := "http://localhost:493/todo"
	fmt.Println("URL:>", url)

	var jsonStr = []byte(`{"Title":"Buy cheese and bread for breakfast.", "Body":"ALLER", "Completed":"true"}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	// req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.StatusCode)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	if resp.StatusCode != 201 {
		t.Errorf("Insertion problem")
	}
}
