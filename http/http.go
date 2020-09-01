package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

func main() {
	doGet()
	doPost()
	doFormPost()
	doCustomPost()
}

func doGet() {
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatalln(err)
	}

	//Dont forget to close resources
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))
}

func doPost() {
	requestBody, err := json.Marshal(map[string]string{
		"name":  "Daniel",
		"email": "test@email.com",
	})
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalln(err)
	}

	//Dont forget to close resources
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Request:", string(requestBody))
	log.Println("Response:", string(body))
}

func doFormPost() {
	formData := url.Values{
		"name": {"daniel"},
	}

	resp, err := http.PostForm("https://httpbin.org/post", formData)
	if err != nil {
		log.Fatalln(err)
	}

	//Dont forget to close resources
	defer resp.Body.Close()

	var result map[string]interface{}

	//One way
	//json.NewDecoder(resp.Body).Decode(&result)

	//Another way
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	json.Unmarshal(body, &result)

	//Common way
	log.Println(result["form"])
}

func doCustomPost() {
	requestBody, err := json.Marshal(map[string]string{
		"name":  "Daniel",
		"email": "test@email.com",
	})
	if err != nil {
		log.Fatalln(err)
	}

	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	request, err := http.NewRequest("POST", "https://httpbin.org/post", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-type", "application/json")

	if err != nil {
		log.Fatalln(err)
	}

	resp, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}

	//Dont forget to close resources
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))

}
