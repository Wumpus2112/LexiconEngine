// Reading and writing files are basic tasks needed for
// many Go programs. First we'll look at some examples of
// reading files.

package main


import (
	"bufio"
	"fmt"
	"os"
	"io/ioutil"

	"net/http"
	"log"
	"encoding/json"
	"net/url"

	"labix.org/v2/mgo"
)


func check(e error) {
	if e != nil {
		panic(e)
	}
}


func main() {
	fmt.Println("BEGIN PROGRAM")

	//filename := "\\Data\\google-10000-english-master\\google-10000-english.txt"
	//readLines(filename)
	//writeflie()

	getentry ("ace")


}


// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Print("No file found")
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}





func writeflie() {

	// To start, here's how to dump a string (or just
	// bytes) into a file.
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("/dat1", d1, 0644)
	check(err)

	// For more granular writes, open a file for writing.
	f, err := os.Create("/dat2")
	check(err)

	// It's idiomatic to defer a `Close` immediately
	// after opening a file.
	defer f.Close()

	// You can `Write` byte slices as you'd expect.
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	// A `WriteString` is also available.
	n3, err := f.WriteString("writes\n")
	fmt.Printf("wrote %d bytes\n", n3)

	// Issue a `Sync` to flush writes to stable storage.
	f.Sync()

	// `bufio` provides buffered writers in addition
	// to the buffered readers we saw earlier.
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n4)

	// Use `Flush` to ensure all buffered operations have
	// been applied to the underlying writer.
	w.Flush()

}



func getentry(word string) {
	//phone := "14158586273"
	// QueryEscape escapes the phone string so
	// it can be safely placed inside a URL query
	safeWord := url.QueryEscape(word)
	//url := fmt.Sprintf("http://apilayer.net/api/validate?access_key=YOUR_ACCESS_KEY&number=%s", safePhone)

	//word := "ace"

	fmt.Println(safeWord)
	urlText := "https://od-api.oxforddictionaries.com:443/api/v1/entries/en/" + safeWord
	url := fmt.Sprintf(urlText)

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
	var record EntriesSchema

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}

	for i:=0;i< len(record.Results); i++  {
		fmt.Println("Word =",record.Results[i].Word)
		fmt.Println("Type =",record.Results[i].Type)

		for j:=0;j<len(record.Results[i].LexicalEntries); j++{
			fmt.Println("-> Definition",j)
			fmt.Println("Category=",record.Results[i].LexicalEntries[j].LexicalCategory)

			for k:=0;k<len(record.Results[i].LexicalEntries[j].Entries); k++ {

				for l:=0;l<len(record.Results[i].LexicalEntries[j].Entries[k].GrammaticalFeatures); l++ {
					fmt.Println("Feature Type=",record.Results[i].LexicalEntries[j].Entries[k].GrammaticalFeatures[l].Type)
					fmt.Println("Feature Text=",record.Results[i].LexicalEntries[j].Entries[k].GrammaticalFeatures[l].Text)
				}

				for l:=0;l<len(record.Results[i].LexicalEntries[j].Entries[k].Senses); l++ {
					fmt.Println("Domain=",record.Results[i].LexicalEntries[j].Entries[k].Senses[l].Domains)

					for m:=0;m<len(record.Results[i].LexicalEntries[j].Entries[k].Senses[l].Subsenses); m++ {
						fmt.Println("Domain=",record.Results[i].LexicalEntries[j].Entries[k].Senses[l].Subsenses[m].Domains)
					}

				}

			}



			}


	}





	fmt.Println("Word = ", record.Results[0].Word)
	fmt.Println("Type = ", record.Results[0].Type)

}


type EntriesSchema struct {
	Metadata struct {
		Provider string `json:"provider"`
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
				Senses []struct {
					Definitions []string `json:"definitions"`
					Domains []string `json:"domains"`
					Examples []struct {
						Registers []string `json:"registers,omitempty"`
						Text string `json:"text"`
					} `json:"examples"`
					ID string `json:"id"`
					Registers []string `json:"registers,omitempty"`
					Subsenses []struct {
						Definitions []string `json:"definitions"`
						Domains []string `json:"domains"`
						Examples []struct {
							Text string `json:"text"`
						} `json:"examples"`
						ID string `json:"id"`
						Registers []string `json:"registers"`
					} `json:"subsenses,omitempty"`
				} `json:"senses"`
			} `json:"entries"`
			Language string `json:"language"`
			LexicalCategory string `json:"lexicalCategory"`
			Pronunciations []struct {
				AudioFile string `json:"audioFile"`
				Dialects []string `json:"dialects"`
				PhoneticNotation string `json:"phoneticNotation"`
				PhoneticSpelling string `json:"phoneticSpelling"`
			} `json:"pronunciations"`
			Text string `json:"text"`
		} `json:"lexicalEntries"`
		Type string `json:"type"`
		Word string `json:"word"`
	} `json:"results"`
}

//type EntriesSchema struct {
//	Metadata struct {
//	} `json:"metadata"`
//	Results []struct {
//		ID string `json:"id"`
//		Language string `json:"language"`
//		LexicalEntries []struct {
//			Entries []struct {
//				Etymologies []string `json:"etymologies"`
//				GrammaticalFeatures []struct {
//					Text string `json:"text"`
//					Type string `json:"type"`
//				} `json:"grammaticalFeatures"`
//				HomographNumber string `json:"homographNumber"`
//				Pronunciations []struct {
//					AudioFile string `json:"audioFile"`
//					Dialects []string `json:"dialects"`
//					PhoneticNotation string `json:"phoneticNotation"`
//					PhoneticSpelling string `json:"phoneticSpelling"`
//					Regions []string `json:"regions"`
//				} `json:"pronunciations"`
//				Senses []struct {
//					CrossReferenceMarkers []string `json:"crossReferenceMarkers"`
//					CrossReferences []struct {
//						ID string `json:"id"`
//						Text string `json:"text"`
//						Type string `json:"type"`
//					} `json:"crossReferences"`
//					Definitions []string `json:"definitions"`
//					Domains []string `json:"domains"`
//					Examples []struct {
//						Definitions []string `json:"definitions"`
//						Domains []string `json:"domains"`
//						Regions []string `json:"regions"`
//						Registers []string `json:"registers"`
//						SenseIds []string `json:"senseIds"`
//						Text string `json:"text"`
//						Translations []struct {
//							Domains []string `json:"domains"`
//							GrammaticalFeatures []struct {
//								Text string `json:"text"`
//								Type string `json:"type"`
//							} `json:"grammaticalFeatures"`
//							Language string `json:"language"`
//							Regions []string `json:"regions"`
//							Registers []string `json:"registers"`
//							Text string `json:"text"`
//						} `json:"translations"`
//					} `json:"examples"`
//					ID string `json:"id"`
//					Pronunciations []struct {
//						AudioFile string `json:"audioFile"`
//						Dialects []string `json:"dialects"`
//						PhoneticNotation string `json:"phoneticNotation"`
//						PhoneticSpelling string `json:"phoneticSpelling"`
//						Regions []string `json:"regions"`
//					} `json:"pronunciations"`
//					Regions []string `json:"regions"`
//					Registers []string `json:"registers"`
//					Subsenses []struct {
//						Domains []string `json:"domains"`
//					} `json:"subsenses"`
//					Translations []struct {
//						Domains []string `json:"domains"`
//						GrammaticalFeatures []struct {
//							Text string `json:"text"`
//							Type string `json:"type"`
//						} `json:"grammaticalFeatures"`
//						Language string `json:"language"`
//						Regions []string `json:"regions"`
//						Registers []string `json:"registers"`
//						Text string `json:"text"`
//					} `json:"translations"`
//					VariantForms []struct {
//						Regions []string `json:"regions"`
//						Text string `json:"text"`
//					} `json:"variantForms"`
//				} `json:"senses"`
//				VariantForms []struct {
//					Regions []string `json:"regions"`
//					Text string `json:"text"`
//				} `json:"variantForms"`
//			} `json:"entries"`
//			GrammaticalFeatures []struct {
//				Text string `json:"text"`
//				Type string `json:"type"`
//			} `json:"grammaticalFeatures"`
//			Language string `json:"language"`
//			LexicalCategory string `json:"lexicalCategory"`
//			Pronunciations []struct {
//				AudioFile string `json:"audioFile"`
//				Dialects []string `json:"dialects"`
//				PhoneticNotation string `json:"phoneticNotation"`
//				PhoneticSpelling string `json:"phoneticSpelling"`
//				Regions []string `json:"regions"`
//			} `json:"pronunciations"`
//			Text string `json:"text"`
//			VariantForms []struct {
//				Regions []string `json:"regions"`
//				Text string `json:"text"`
//			} `json:"variantForms"`
//		} `json:"lexicalEntries"`
//		Pronunciations []struct {
//			AudioFile string `json:"audioFile"`
//			Dialects []string `json:"dialects"`
//			PhoneticNotation string `json:"phoneticNotation"`
//			PhoneticSpelling string `json:"phoneticSpelling"`
//			Regions []string `json:"regions"`
//		} `json:"pronunciations"`
//		Type string `json:"type"`
//		Word string `json:"word"`
//	} `json:"results"`
//}
