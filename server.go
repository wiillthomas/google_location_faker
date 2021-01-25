package main

import (
	"encoding/base64"
	"encoding/csv"
	"encoding/json"
	secretvalue "google_location_faker/lib"
	"log"
	"net/http"
	"os"

	"github.com/didip/tollbooth"
	"github.com/falmar/goradix"
)

var radix = goradix.New(false)

type JSONResponse struct {
	message string
}

type GoogleLocation struct {
	Criteria_ID    string
	Name           string
	Canonical_Name string
	Parent_ID      string
	Country_Code   string
	Target_Type    string
	Status         string
}

func handleAPICall(w http.ResponseWriter, r *http.Request) {

	locationKeys, locationOk := r.URL.Query()["location"]
	queryKeys, queryOk := r.URL.Query()["q"]

	if !locationOk || len(locationKeys[0]) < 1 {
		log.Println("Url Param 'location' is missing")
		return
	}

	if !queryOk || len(queryKeys[0]) < 1 {
		log.Println("Url Parameter 'q' is missing ")
		return
	}

	if len(locationKeys[0]) > 100 {
		log.Println("Url parameter 'location' is too long")
	}

	if len(queryKeys[0]) > 100 {
		log.Println("Url parameter 'q' is too long")
	}

	location := locationKeys[0]
	query := queryKeys[0]

	uuleValue := "w+CAIQICI"

	parsedLocationCode, parseLocationOk := secretvalue.CalculateSecretValue(len(location))

	log.Println(parseLocationOk)

	if !parseLocationOk {
		return
	}

	var base64EncodedLocation string = base64.StdEncoding.EncodeToString([]byte(location))

	response := "q=" + query + "&uule=" + uuleValue + parsedLocationCode + base64EncodedLocation

	jsonOutput, jsonErr := json.Marshal(response)

	log.Println("hello")

	if jsonErr == nil {
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(jsonOutput)

}

func ReadCsv(filename string) ([][]string, error) {

	f, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}

func handleAutoCompleteInput(w http.ResponseWriter, r *http.Request) {
	inputKeys, inputOK := r.URL.Query()["input"]

	if !inputOK || len(inputKeys[0]) < 1 {
		log.Println("Url parameter 'input' is missing ")
		return
	}

	if len(inputKeys[0]) > 100 {
		log.Println("Url parameter 'input' is too long")
		return
	}

	log.Println(inputKeys[0])

	words, err := radix.AutoComplete(inputKeys[0], false)

	if err != nil {
		return
	}

	jsonOutput, jsonErr := json.Marshal(words)

	if jsonErr != nil {
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(jsonOutput)

}

func main() {
	log.Println("Server Starting...")

	lines, err := ReadCsv("google_locations.csv")
	if err != nil {
		panic(err)
	}
	for _, line := range lines {
		radix.Insert(line[2])
	}

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.Handle("/api", tollbooth.LimitFuncHandler(tollbooth.NewLimiter(1, nil), handleAPICall))
	http.HandleFunc("/autocomplete", handleAutoCompleteInput)
	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}
