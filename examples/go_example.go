// Unlocoder API examples - Go
// Get your free API key at: https://rapidapi.com/contactliamnoonan/api/unlocoder

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	apiKey  = "YOUR_RAPIDAPI_KEY"
	baseURL = "https://unlocoder.p.rapidapi.com"
	host    = "unlocoder.p.rapidapi.com"
)

type ConvertRequest struct {
	Input     string `json:"input"`
	Precision int    `json:"precision,omitempty"`
}

type ConvertResponse struct {
	Latitude       float64           `json:"latitude"`
	Longitude      float64           `json:"longitude"`
	DetectedFormat string            `json:"detectedFormat"`
	Precision      int               `json:"precision"`
	Outputs        map[string]string `json:"outputs"`
	Location       *Location         `json:"location,omitempty"`
	NearbyUnLocodes []NearbyResult   `json:"nearbyUnLocodes,omitempty"`
}

type Location struct {
	Code       string  `json:"code"`
	Name       string  `json:"name"`
	TimezoneID string  `json:"timezoneId"`
	UTCOffset  string  `json:"utcOffset"`
	LocalTime  string  `json:"localTime"`
}

type NearbyResult struct {
	Country    string  `json:"country"`
	Location   string  `json:"location"`
	DistanceKm float64 `json:"distanceKm"`
}

func convertCoordinate(input string, precision int) (*ConvertResponse, error) {
	body, _ := json.Marshal(ConvertRequest{Input: input, Precision: precision})
	req, _ := http.NewRequest("POST", baseURL+"/api/convert", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-rapidapi-key", apiKey)
	req.Header.Set("x-rapidapi-host", host)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP %d: %s", resp.StatusCode, data)
	}

	var result ConvertResponse
	return &result, json.Unmarshal(data, &result)
}

func findNearby(lat, lng float64) ([]NearbyResult, error) {
	u := fmt.Sprintf("%s/unlocodes/nearby?latitude=%f&longitude=%f", baseURL, lat, lng)
	req, _ := http.NewRequest("GET", u, nil)
	req.Header.Set("x-rapidapi-key", apiKey)
	req.Header.Set("x-rapidapi-host", host)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP %d: %s", resp.StatusCode, data)
	}

	var results []NearbyResult
	return results, json.Unmarshal(data, &results)
}

func lookupUnlocode(code string, referenceTime string) (map[string]interface{}, error) {
	u := baseURL + "/unlocodes/" + url.PathEscape(code)
	if referenceTime != "" {
		u += "?referenceTime=" + url.QueryEscape(referenceTime)
	}
	req, _ := http.NewRequest("GET", u, nil)
	req.Header.Set("x-rapidapi-key", apiKey)
	req.Header.Set("x-rapidapi-host", host)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP %d: %s", resp.StatusCode, data)
	}

	var result map[string]interface{}
	return result, json.Unmarshal(data, &result)
}

func main() {
	// 1. Convert decimal degrees
	result, err := convertCoordinate("51.5074, -0.1278", 6)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Detected: %s\n", result.DetectedFormat)
	for fmt_name, value := range result.Outputs {
		fmt.Printf("  %s: %s\n", fmt_name, value)
	}

	// 2. Resolve UN/LOCODE
	locode, err := convertCoordinate("GBLON", 6)
	if err == nil && locode.Location != nil {
		fmt.Printf("\nGBLON: %s, TZ: %s\n", locode.Location.Name, locode.Location.TimezoneID)
	}

	// 3. Find nearby
	nearby, err := findNearby(40.7128, -74.0060)
	if err == nil {
		fmt.Println("\nNearby UN/LOCODEs to New York:")
		for _, n := range nearby {
			fmt.Printf("  %s%s - %.1f km\n", n.Country, n.Location, n.DistanceKm)
		}
	}
}
