package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2/log"
)

type chainlinkDataStructure struct {
	JobRunID   int   `json:"id"`
	Data       Data  `json:"data"`
	Result     int64 `json:"result"`
	StatusCode int   `json:"statusCode"`
}

// Define a struct that matches the JSON response structure
type Data struct {
	Address string `json:"address"`
	Balance int64  `json:"balance"`
}

func main() {

	http.HandleFunc("/getMax", handler)

	port := "8080" // Change this to the desired port
	fmt.Printf("Listening on :%s...\n", port)
	err := http.ListenAndServe("0.0.0.0:"+port, nil)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

}

func handler(w http.ResponseWriter, r *http.Request) {

	data := getMax()

	responseData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseData)
}

func getMax() chainlinkDataStructure {
	values := fetchData()

	var maxData Data
	maxValue := values[0].Balance

	for _, value := range values {
		if value.Balance > maxValue {
			maxValue = value.Balance
			maxData = Data{
				Address: value.Address,
				Balance: maxValue,
			}
		}
	}

	return chainlinkDataStructure{
		JobRunID:   1,
		Data:       maxData,
		Result:     maxData.Balance,
		StatusCode: 200,
	}
}

// Get Endpoint
func fetchData() []Data {
	url := "https://gist.githubusercontent.com/thodges-gh/3bd03660676504478de60c3a17800556/raw/0013f560b97eb1b2481fd4d57f02507c96f0d88f/balances.json"

	// Send a GET request to the API
	resp, err := http.Get(url)
	if err != nil {
		log.Error("Error occoured while getting")
		return nil
	}

	// ensure that the response body of the HTTP request is closed properly when you're done processing the response.
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return nil
	}

	var data []Data
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Printf("Error parsing JSON response: %v\n", err)
		return nil
	}

	return data
}
