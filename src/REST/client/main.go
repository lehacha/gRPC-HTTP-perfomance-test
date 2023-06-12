package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type ProcessRequest struct {
	Text     string `json:"text"`
	Username string `json:"username"`
}

type ProcessResponse struct {
	Words   []string `json:"words"`
	Message string   `json:"message"`
}

func main() {

	var timeToMake1000req []float64

	for j := 0; j < 100; j++ {
		// experiment with 1000 requests

		// experiment start time
		start := time.Now()
		for i := 0; i < 1000; i++ {

			requestBody, err := json.Marshal(ProcessRequest{
				Text:     "test text",
				Username: "test user name",
			})
			if err != nil {
				fmt.Printf("error marshaling request body %s\n", err)

				return
			}

			bodyReader := bytes.NewReader(requestBody)
			req, err := http.NewRequest(http.MethodPost, "http://localhost:8090/process", bodyReader)
			if err != nil {
				fmt.Printf("could not create request: %s\n", err)

				break
			}

			res, err := http.DefaultClient.Do(req)
			if err != nil {
				fmt.Printf("error executing http request: %s\n", err)

				break
			}

			resBody, err := ioutil.ReadAll(res.Body)
			if err != nil {
				fmt.Printf("server: could not read response body: %s\n", err)
			}

			// Declare a new ProcessResponse struct.
			var processResponse ProcessResponse

			// Try to decode the request body into the struct
			err = json.Unmarshal(resBody, &processResponse)
			if err != nil {
				fmt.Printf("error decodding response body: %s\n", err)

				break
			}

		}

		// difference: now - start
		dif := time.Now().Sub(start)
		fmt.Printf("REST: %v \n", dif)

		timeToMake1000req = append(timeToMake1000req, float64(dif.Milliseconds()))

	}

	var sum float64
	for _, t := range timeToMake1000req {
		sum += t
	}
	fmt.Printf("[REST] Result: %vms\n", sum/float64(len(timeToMake1000req)))
}
