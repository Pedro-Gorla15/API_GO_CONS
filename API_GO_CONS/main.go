package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Dog struct {
	Message string `json: "message"`
	Status  string `json: "status"`
}

func main() {
	printf := fmt.Printf
	logf := log.Printf
	printf("Consuming API using GO\n")

	url := "https://dog.ceo/api/breeds/image/random"

	request, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		logf("Could not prepare a new request - %v", err)
	}

	request.Header.Add("content-type", "application/json")
	request.Header.Add("Authorization", "api-key afgl6752390nasdjghtuyvbn786")

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		logf("Could not perform a new request - %v", err)
	}

	responseBytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		logf("Could not read response body - %v", err)
	}

	var dog Dog

	if err := json.Unmarshal(responseBytes, &dog); err != nil {
		logf("Cloud not unmarshal json response byte - %v", err)
	}

	fileUrl := string(dog.Message)

	resp, err := http.Get(fileUrl)

	if err != nil {
		logf("Cloud not perfom a new request to file URL - %v", err)
	}

	defer resp.Body.Close()

	out, err := os.Create("dog.jpg")

	defer out.Close()

	_, err = io.Copy(out, resp.Body)

	if err != nil {
		logf("Cloud not copy bytes to the output file - %v", err)
	}

}
