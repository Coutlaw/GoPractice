package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

// Error : Create the error logger
var Error *log.Logger

// Init : Initilize the logs
func Init(
	errorHandle io.Writer) {
	Error = log.New(errorHandle, "ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

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

// GetJSON : function to take JSON file and Unmarshall to the Client Struct (returns bytes)
func GetJSON(file string) Client {
	raw, err := ioutil.ReadFile(file)
	if err != nil {
		Error.Println("Unable to properly read file: ", err)
		os.Exit(1)
	}

	var c Client
	json.Unmarshal(raw, &c)
	return c
}

func main() {
	// for the error logg
	Init(os.Stderr)

	var testClient = GetJSON("./testFile.json")
	fmt.Printf("ClientId: %v\nClientSecret: %v\n", testClient.ClientId, testClient.ClientSecret)

	var APIKey = testClient.generateKey()
	fmt.Println(APIKey)
}
