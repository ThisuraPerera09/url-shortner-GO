package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const baseURL = "http://localhost:8080"

type ShortenRequest struct {
	URL        string `json:"url"`
	CustomCode string `json:"custom_code,omitempty"`
}

type ShortenResponse struct {
	ShortCode   string `json:"short_code"`
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

type StatsResponse struct {
	ShortCode    string     `json:"short_code"`
	OriginalURL  string     `json:"original_url"`
	Clicks       int64      `json:"clicks"`
	CreatedAt    time.Time  `json:"created_at"`
	LastAccessed *time.Time `json:"last_accessed,omitempty"`
}

func main() {
	fmt.Println("🧪 Testing URL Shortener API")
	fmt.Println("============================\n")

	// Test 1: Health Check
	fmt.Println("1️⃣  Testing Health Check...")
	testHealthCheck()

	// Test 2: Shorten URL
	fmt.Println("\n2️⃣  Testing URL Shortening...")
	shortCode := testShortenURL("https://www.github.com/golang/go", "")

	// Test 3: Custom Short Code
	fmt.Println("\n3️⃣  Testing Custom Short Code...")
	customCode := testShortenURL("https://www.golang.org", "golang")

	// Test 4: Get Stats
	fmt.Println("\n4️⃣  Testing Get Stats...")
	testGetStats(shortCode)

	// Test 5: Redirect (test click tracking)
	fmt.Println("\n5️⃣  Testing Redirect and Click Tracking...")
	testRedirect(shortCode)
	time.Sleep(100 * time.Millisecond) // Wait for async update
	testGetStats(shortCode)            // Should show clicks incremented

	// Test 6: List URLs
	fmt.Println("\n6️⃣  Testing List URLs...")
	testListURLs()

	// Test 7: Delete URL
	fmt.Println("\n7️⃣  Testing Delete URL...")
	testDeleteURL(customCode)

	fmt.Println("\n✅ All tests completed!")
}

func testHealthCheck() {
	resp, err := http.Get(baseURL + "/api/health")
	if err != nil {
		fmt.Printf("❌ Failed: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode == 200 {
		fmt.Printf("✅ Health check passed: %s\n", string(body))
	} else {
		fmt.Printf("❌ Health check failed: %d\n", resp.StatusCode)
	}
}

func testShortenURL(url, customCode string) string {
	req := ShortenRequest{URL: url, CustomCode: customCode}
	jsonData, _ := json.Marshal(req)

	resp, err := http.Post(baseURL+"/api/shorten", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("❌ Failed: %v\n", err)
		return ""
	}
	defer resp.Body.Close()

	var result ShortenResponse
	json.NewDecoder(resp.Body).Decode(&result)

	if resp.StatusCode == 201 {
		fmt.Printf("✅ URL shortened successfully\n")
		fmt.Printf("   Original: %s\n", result.OriginalURL)
		fmt.Printf("   Short Code: %s\n", result.ShortCode)
		fmt.Printf("   Short URL: %s\n", result.ShortURL)
		return result.ShortCode
	} else {
		fmt.Printf("❌ Failed with status: %d\n", resp.StatusCode)
		return ""
	}
}

func testGetStats(shortCode string) {
	resp, err := http.Get(baseURL + "/api/stats/" + shortCode)
	if err != nil {
		fmt.Printf("❌ Failed: %v\n", err)
		return
	}
	defer resp.Body.Close()

	var stats StatsResponse
	json.NewDecoder(resp.Body).Decode(&stats)

	if resp.StatusCode == 200 {
		fmt.Printf("✅ Stats retrieved successfully\n")
		fmt.Printf("   Short Code: %s\n", stats.ShortCode)
		fmt.Printf("   Original URL: %s\n", stats.OriginalURL)
		fmt.Printf("   Clicks: %d\n", stats.Clicks)
		fmt.Printf("   Created: %s\n", stats.CreatedAt.Format(time.RFC3339))
	} else {
		fmt.Printf("❌ Failed with status: %d\n", resp.StatusCode)
	}
}

func testRedirect(shortCode string) {
	// Create client that doesn't follow redirects
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	resp, err := client.Get(baseURL + "/" + shortCode)
	if err != nil {
		fmt.Printf("❌ Failed: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 301 {
		location := resp.Header.Get("Location")
		fmt.Printf("✅ Redirect working\n")
		fmt.Printf("   Redirects to: %s\n", location)
	} else {
		fmt.Printf("❌ Expected 301 redirect, got: %d\n", resp.StatusCode)
	}
}

func testListURLs() {
	resp, err := http.Get(baseURL + "/api/urls?limit=10&offset=0")
	if err != nil {
		fmt.Printf("❌ Failed: %v\n", err)
		return
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	if resp.StatusCode == 200 {
		urls := result["urls"].([]interface{})
		fmt.Printf("✅ Retrieved %d URLs\n", len(urls))
		fmt.Printf("   Total count: %.0f\n", result["count"].(float64))
	} else {
		fmt.Printf("❌ Failed with status: %d\n", resp.StatusCode)
	}
}

func testDeleteURL(shortCode string) {
	client := &http.Client{}
	req, _ := http.NewRequest("DELETE", baseURL+"/api/urls/"+shortCode, nil)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("❌ Failed: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		fmt.Printf("✅ URL deleted successfully: %s\n", shortCode)
	} else {
		fmt.Printf("❌ Failed with status: %d\n", resp.StatusCode)
	}
}
