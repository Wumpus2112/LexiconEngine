package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Numverify struct {
	Valid               bool   `json:"valid"`
	Number              string `json:"number"`
	LocalFormat         string `json:"local_format"`
	InternationalFormat string `json:"international_format"`
	CountryPrefix       string `json:"country_prefix"`
	CountryCode         string `json:"country_code"`
	CountryName         string `json:"country_name"`
	Location            string `json:"location"`
	Carrier             string `json:"carrier"`
	LineType            string `json:"line_type"`
}

func entries(word string) {
	//phone := "14158586273"
	// QueryEscape escapes the phone string so
	// it can be safely placed inside a URL query
	//safePhone := url.QueryEscape(phone)
	//url := fmt.Sprintf("http://apilayer.net/api/validate?access_key=YOUR_ACCESS_KEY&number=%s", safePhone)

	//word := "ace"
	url := fmt.Sprintf("https://od-api.oxforddictionaries.com:443/api/v1/entries/en/",word)

	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	// For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	// A Client is an HTTP client
	client := &http.Client{}


	// headers
	req.Header.Set("Accept", "application/json")
	req.Header.Set("app_id", "c305db46")
	req.Header.Set("app_key", "539ae75d43042248f92ce6f6a07a8d8d")

	//"Accept": "application/json",
	//"app_id": "c305db46",
	//"app_key": "539ae75d43042248f92ce6f6a07a8d8d"


	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()

	// Fill the record with the data from the JSON
	var record DictionarySchema

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}

	fmt.Println("Word = ", record.Results[0].Word)
	fmt.Println("Type = ", record.Results[0].Type)

}


type DictionarySchema struct {
	Metadata struct {
	} `json:"metadata"`
	Results []struct {
		ID string `json:"id"`
		Language string `json:"language"`
		LexicalEntries []struct {
			Entries []struct {
				Etymologies []string `json:"etymologies"`
				GrammaticalFeatures []struct {
					Text string `json:"text"`
					Type string `json:"type"`
				} `json:"grammaticalFeatures"`
				HomographNumber string `json:"homographNumber"`
				Pronunciations []struct {
					AudioFile string `json:"audioFile"`
					Dialects []string `json:"dialects"`
					PhoneticNotation string `json:"phoneticNotation"`
					PhoneticSpelling string `json:"phoneticSpelling"`
					Regions []string `json:"regions"`
				} `json:"pronunciations"`
				Senses []struct {
					CrossReferenceMarkers []string `json:"crossReferenceMarkers"`
					CrossReferences []struct {
						ID string `json:"id"`
						Text string `json:"text"`
						Type string `json:"type"`
					} `json:"crossReferences"`
					Definitions []string `json:"definitions"`
					Domains []string `json:"domains"`
					Examples []struct {
						Definitions []string `json:"definitions"`
						Domains []string `json:"domains"`
						Regions []string `json:"regions"`
						Registers []string `json:"registers"`
						SenseIds []string `json:"senseIds"`
						Text string `json:"text"`
						Translations []struct {
							Domains []string `json:"domains"`
							GrammaticalFeatures []struct {
								Text string `json:"text"`
								Type string `json:"type"`
							} `json:"grammaticalFeatures"`
							Language string `json:"language"`
							Regions []string `json:"regions"`
							Registers []string `json:"registers"`
							Text string `json:"text"`
						} `json:"translations"`
					} `json:"examples"`
					ID string `json:"id"`
					Pronunciations []struct {
						AudioFile string `json:"audioFile"`
						Dialects []string `json:"dialects"`
						PhoneticNotation string `json:"phoneticNotation"`
						PhoneticSpelling string `json:"phoneticSpelling"`
						Regions []string `json:"regions"`
					} `json:"pronunciations"`
					Regions []string `json:"regions"`
					Registers []string `json:"registers"`
					Subsenses []struct {
					} `json:"subsenses"`
					Translations []struct {
						Domains []string `json:"domains"`
						GrammaticalFeatures []struct {
							Text string `json:"text"`
							Type string `json:"type"`
						} `json:"grammaticalFeatures"`
						Language string `json:"language"`
						Regions []string `json:"regions"`
						Registers []string `json:"registers"`
						Text string `json:"text"`
					} `json:"translations"`
					VariantForms []struct {
						Regions []string `json:"regions"`
						Text string `json:"text"`
					} `json:"variantForms"`
				} `json:"senses"`
				VariantForms []struct {
					Regions []string `json:"regions"`
					Text string `json:"text"`
				} `json:"variantForms"`
			} `json:"entries"`
			GrammaticalFeatures []struct {
				Text string `json:"text"`
				Type string `json:"type"`
			} `json:"grammaticalFeatures"`
			Language string `json:"language"`
			LexicalCategory string `json:"lexicalCategory"`
			Pronunciations []struct {
				AudioFile string `json:"audioFile"`
				Dialects []string `json:"dialects"`
				PhoneticNotation string `json:"phoneticNotation"`
				PhoneticSpelling string `json:"phoneticSpelling"`
				Regions []string `json:"regions"`
			} `json:"pronunciations"`
			Text string `json:"text"`
			VariantForms []struct {
				Regions []string `json:"regions"`
				Text string `json:"text"`
			} `json:"variantForms"`
		} `json:"lexicalEntries"`
		Pronunciations []struct {
			AudioFile string `json:"audioFile"`
			Dialects []string `json:"dialects"`
			PhoneticNotation string `json:"phoneticNotation"`
			PhoneticSpelling string `json:"phoneticSpelling"`
			Regions []string `json:"regions"`
		} `json:"pronunciations"`
		Type string `json:"type"`
		Word string `json:"word"`
	} `json:"results"`
}
