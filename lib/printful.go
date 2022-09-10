package lib

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func PrintfulSendRequest(method string, endpoint string) (*interface{}, error) {
	req, err := http.NewRequest(method, os.Getenv("PRINTFUL_API_URL") + endpoint, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Basic " + os.Getenv("PRINTFUL_API_KEY"))
	req.Header.Add("X-PF-Store-Id", "8726494")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)

	if resp.StatusCode != http.StatusOK {
		defer resp.Body.Close()
		_, _ = io.ReadAll(resp.Body)

		return nil, err
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	var data interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}


