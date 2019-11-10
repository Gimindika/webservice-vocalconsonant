package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

//VocalConsonantCount is a Json format of the result to display
type VocalConsonantCount struct {
	StringInput string `json:"stringInput"`
	Result      string `json:"result"`
	Vocal       int    `json:"vocal"`
	Consonant   int    `json:"consonant"`
}

func landingPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `Request to url with params :
		the string input for the vocal consonant counting function
		
		Example: localhost:8080/vocalconsonant/osama`)
}

func vocalConsonant(str string) (string, int, int) {
	vocalList := "aeiou"
	consonantList := "bcdfghjklmnpqrstvwxyz"

	vocalCount := make(map[string]int)
	consonantCount := make(map[string]int)

	for _, char := range str {
		tmpChar := string(char)
		tmpChar = strings.ToLower(tmpChar)

		if strings.Contains(vocalList, tmpChar) {
			if vocalCount[tmpChar] == 1 {
				vocalCount[tmpChar]++
			} else {
				vocalCount[tmpChar] = 1
			}
		} else if strings.Contains(consonantList, tmpChar) { //only consonants and vocals will be counted, special characters will be ignored
			if consonantCount[tmpChar] == 1 {
				consonantCount[tmpChar]++
			} else {
				consonantCount[tmpChar] = 1
			}
		}
	}

	var result string = "huruf mati: " + strconv.Itoa(len(consonantCount)) + ", huruf hidup: " + strconv.Itoa(len(vocalCount))

	return result, len(vocalCount), len(consonantCount)
}

func vocalConsonantHandler(w http.ResponseWriter, r *http.Request) {
	str := strings.Replace(r.URL.Path, "/vocalconsonant/", "", 1)

	if str != "" {
		result, vocal, consonant := vocalConsonant(str)
		w.Header().Set("Content-Type", "application/json")
		var resultToShow VocalConsonantCount
		resultToShow.StringInput = str
		resultToShow.Result = result
		resultToShow.Vocal = vocal
		resultToShow.Consonant = consonant
		json.NewEncoder(w).Encode(resultToShow)
	} else {
		landingPage(w, r)
	}

	// with query
	// keys, ok := r.URL.Query()["str"]
	// if ok {
	// 	str := keys[0]
	//  result, vocal, consonant := vocalConsonant(str)
	// 	w.Header().Set("Content-Type", "application/json")
	// 	var resultToShow VocalConsonantCount
	// 	resultToShow.StringInput = str
	// 	resultToShow.Result = result
	// 	resultToShow.Vocal = vocal
	// 	resultToShow.Consonant = consonant
	// 	json.NewEncoder(w).Encode(resultToShow)
	// } else {
	// 	landingPage(w, r)
	// }
}

func handleRequests() {
	http.HandleFunc("/vocalconsonant/", vocalConsonantHandler)
	http.HandleFunc("/", landingPage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
