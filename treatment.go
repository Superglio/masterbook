package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type PubmedResponse struct {
	ResultList PubmedResultList `json:"resultList"`
}

type PubmedResultList struct {
	Result []PubmedResult `json:"result"`
}

type PubmedResult struct {
	Title       string `json:"title"`
	Abstract    string `json:"abstract"`
	PubDate     string `json:"pubDate"`
	StudyType   string `json:"studyType"`
	Treatment   string `json:"treatment"`
}

func main() {
	// Make a query to PubMed API
	query := "your search query" // Replace with your search query
	resultCount := 5             // Number of results to retrieve

	// Build the PubMed API URL
	baseURL := "https://eutils.ncbi.nlm.nih.gov/entrez/eutils/esearch.fcgi"
	queryParams := url.Values{}
	queryParams.Set("db", "pubmed")
	queryParams.Set("retmode", "json")
	queryParams.Set("retmax", fmt.Sprintf("%d", resultCount))
	queryParams.Set("term", query)

	fullURL := fmt.Sprintf("%s?%s", baseURL, queryParams.Encode())

	// Send a GET request to PubMed API
	response, err := http.Get(fullURL)
	if err != nil {
		fmt.Println("Error making PubMed API request:", err)
		return
	}
	defer response.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading PubMed API response:", err)
		return
	}

	// Parse the JSON response
	var pubmedResponse PubmedResponse
	err = json.Unmarshal(body, &pubmedResponse)
	if err != nil {
		fmt.Println("Error parsing PubMed API response:", err)
		return
	}

	// Process the results
	for _, result := range pubmedResponse.ResultList.Result {
		fmt.Println("Title:", result.Title)
		fmt.Println("Abstract:", result.Abstract)
		fmt.Println("Publication Date:", result.PubDate)
		fmt.Println("Study Type:", result.StudyType)
		fmt.Println("Treatment:", result.Treatment)
		fmt.Println("----------------------------------")
	}
}
