package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"
)

func main() {
	host := fmt.Sprintf("127.0.0.1:%s", os.Getenv("PORT"))
	safeURL := &url.URL{
		Scheme: "http",
		Host:   host,
		Path:   os.Getenv("HEALTHCHECK"),
	}

	finalURL := safeURL.String()
	fmt.Println("Checking", finalURL)

	client := &http.Client{Timeout: 5 * time.Second}
	res, err := client.Get(finalURL)
	if err != nil {
		fmt.Println("Healthcheck request failed:", err)
		os.Exit(1)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Healthcheck failed, status:", res.StatusCode)
		os.Exit(1)
	}
	fmt.Println("Healthcheck success, status:", res.StatusCode)
}
