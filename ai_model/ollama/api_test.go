package ollama

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestOlama(t *testing.T) {

	// Create JSON request body
	requestBody := map[string]string{
		"model":    "llama2",
		"messages": "Hi",
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		panic(err)
	}

	// Make HTTP request
	url := "http://localhost:11434/api/generate"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Read response line-by-line
	var response strings.Builder
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		fmt.Println(scanner.Text()) // print each line
		response.WriteString(scanner.Text())
	}
	ts := response.String()
	if ts != "" {
		// Print accumulated response
		t.Fatalf(response.String())
	}

}
