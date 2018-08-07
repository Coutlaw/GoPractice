package main

import (
	"testing"
)

var (
	clientSecret = "4b45-beed-17fa839"
	redirectUrls = "www.google.com"
	name         = "its a name"
	clientId     = "05de34eeeecd0ea8"
	description  = "Company: google | Project Description: stuff"
	apiKey       = "788dcc277dc7e426efeb47314746b6fa2841083ddaf306b5cea4b8af3afce944"
)

// TestGenerateKey : is testing JSON reading, and api key calculation
func TestGenerateKey(t *testing.T) {
	var testClient = GetJSON("./testFile.json")
	if testClient.ClientId != clientId {
		t.Errorf("Expected: %v, Got: %v", clientId, testClient.ClientId)
	}
	if testClient.ClientSecret != clientSecret {
		t.Errorf("Expected: %v, Got: %v", clientSecret, testClient.ClientSecret)
	}
	if testClient.generateKey() != apiKey {
		t.Errorf("API Key was not calculated correctly, got %v", testClient.generateKey())
	}
}
