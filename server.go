package main

import (
	"log"
	"net/http"
	"encoding/json"
	"encoding/base64"
	"encoding/csv"
    "os"
)

type JSONResponse struct {
	message string
}

type GoogleLocation struct {
	Criteria_ID string
	Name string
	Canonical_Name string
	Parent_ID string
	Country_Code string
	Target_Type string
	Status string
}

func calculateSecretValue( lengthOfString int ) (string, bool) {
	var char string 

	switch lengthOfString {
		case 4: 
		char = "E"
		case 5: 
		char = "F"
		case 6: 
		char = "G"
		case 7: 
		char = "H"
		case 8: 
		char = "I"
		case 9: 
		char = "J"
		case 10: 
		char = "K"
		case 11: 
		char = "L"
		case 12: 
		char = "M"
		case 13: 
		char = "N"
		case 14: 
		char = "O"
		case 15: 
		char = "P"
		case 16: 
		char = "Q"
		case 17: 
		char = "R"
		case 18: 
		char = "S"
		case 19: 
		char = "T"
		case 20: 
		char = "U"
		case 21: 
		char = "V"
		case 22: 
		char = "W"
		case 23: 
		char = "X"
		case 24: 
		char = "Y"
		case 25: 
		char = "Z"
		case 26: 
		char = "a"
		case 27: 
		char = "b"
		case 28: 
		char = "c"
		case 29: 
		char = "d"
		case 30: 
		char = "e"
		case 31: 
		char = "f" 
		 case 32: 
		char = "g" 
		 case 33: 
		char = "h" 
		 case 34: 
		char = "i" 
		 case 35: 
		char = "j" 
		 case 36: 
		char = "k" 
		 case 37: 
		char = "l" 
		 case 38: 
		char = "m" 
		 case 39: 
		char = "n" 
		 case 40: 
		char = "o" 
		 case 41: 
		char = "p" 
		 case 42: 
		char = "q" 
		 case 43: 
		char = "r" 
		 case 44: 
		char = "s" 
		 case 45: 
		char = "t" 
		 case 46: 
		char = "u" 
		 case 47: 
		char = "v" 
		 case 48: 
		char = "w" 
		 case 49: 
		char = "x" 
		 case 50: 
		char = "y" 
		 case 51: 
		char = "z" 
		 case 52: 
		char = "0" 
		 case 53: 
		char = "1" 
		 case 54: 
		char = "2" 
		 case 55: 
		char = "3" 
		 case 56: 
		char = "4" 
		 case 57: 
		char = "5" 
		 case 58: 
		char = "6" 
		 case 59: 
		char = "7" 
		 case 60: 
		char = "8" 
		 case 61: 
		char = "9" 
		 case 62: 
		char = "-" 
		 case 63: 
		char = " " 
		 case 64: 
		char = "A" 
		 case 65: 
		char = "B" 
		 case 66: 
		char = "C" 
		 case 67: 
		char = "D" 
		 case 68: 
		char = "E" 
		 case 69: 
		char = "F" 
		 case 70: 
		char = "G" 
		 case 71: 
		char = "H" 
		 case 72: 
		char = "I" 
		 case 73: 
		char = "J" 
		 case 74: 
		char = "K" 
		 case 75: 
		char = "L" 
		 case 76: 
		char = "M" 
		 case 77: 
		char = "N" 
		 case 78: 
		char = "O" 
		 case 79: 
		char = "P" 
		 case 80: 
		char = "Q" 
		 case 81: 
		char = "R" 
		 case 82: 
		char = "S" 
		 case 83: 
		char = "T" 
		 case 84: 
		char = "U" 
		 case 85: 
		char = "V" 
		default:
			return "not found", false
	}
	return char, true

}

func handleAPICall( w http.ResponseWriter, r *http.Request ) {
	
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

	location := locationKeys[0]
	query := queryKeys[0]


	log.Println("Url Param 'location' is: " + string(location))
	log.Println("Url Param 'q' is: " + string(query))


	uuleValue := "w+CAIQICI"


	parsedLocationCode, parseLocationOk := calculateSecretValue(len(location))

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

	log.Println( jsonOutput )

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonOutput)

}

func ReadCsv(filename string) ([][]string, error) {

    // Open CSV file
    f, err := os.Open(filename)
    if err != nil {
        return [][]string{}, err
    }
    defer f.Close()

    // Read File into a Variable
    lines, err := csv.NewReader(f).ReadAll()
    if err != nil {
        return [][]string{}, err
    }

    return lines, nil
}

func main() {
	log.Println("Server Starting...")
	googleLocationSlice := make([]GoogleLocation, 102030)

	lines, err := ReadCsv("google_locations.csv")
	if err != nil {
		panic(err)
	}
	for _, line := range lines {
        data := GoogleLocation{
			Criteria_ID: line[0],
			Name: line[1],
			Canonical_Name: line[2],
			Parent_ID: line[3],
			Country_Code: line[4],
			Target_Type: line[5],
			Status: line[6],
        }
		// fmt.Println(data.Name + " " + data.Canonical_Name)
		googleLocationSlice = append(googleLocationSlice, data)
	}

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/api", handleAPICall)
	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}

// TODO:
//  ---- Parse csv into array of structs x
//  ---- Make autocomplete work with array of structs
//  ---- Make endpoint for autocomplete
//  ---- Rate limiting for API
//  ---- Typecheck API
//  ---- Style FE code
//  ---- Tidy up