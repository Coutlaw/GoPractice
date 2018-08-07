package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Client : Struct to represent the clientg
type Client struct {
	ClientSecret string `json:"clientSecret"`
	RedirectUrls string `json:"redirectUrls"`
	Name         string `json:"name"`
	ClientId     string `json:"clientId"`
	Description  string `json:"description"`
}

// Function to calculate a clients useable Key
func (c Client) generateKey() string {
	// like PrintF but returns the string without printing
	comb := fmt.Sprintf("%v%v", c.ClientId, c.ClientSecret)
	myBytes := sha256.Sum256([]byte(comb))
	return hex.EncodeToString(myBytes[:])
}

// function to take JSON file and Unmarshall to the Client Struct (returns bytes)
func getJSON() Client {
	raw, err := ioutil.ReadFile("./testFile.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var c Client
	json.Unmarshal(raw, &c)
	return c
}

func main() {

	var testClient Client = getJSON()
	fmt.Printf("ClientId: %v\nClientSecret: %v\n", testClient.ClientId, testClient.ClientSecret)

	var APIKey = testClient.generateKey()
	fmt.Println(APIKey)
}
