package main

import (
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// HomeVariables :: Template Values For Homepage
type HomeVariables struct {
	Date string
	Time string
}

// ScaleVariables :: Template Values For Scale Buttons
type ScaleVariables struct {
	Date      string
	Time      string
	Scale     string
	Type      string
	Direction string
	Interval  string
	GenScale  bool
	GenInt    bool
	Clear     bool
}

// HomePage :: Update HTML
func HomePage(w http.ResponseWriter, r *http.Request) {

	now := time.Now()
	DateTime := HomeVariables{
		Date: now.Format("02-01-2006"),
		Time: now.Format("15:04:05"),
	}

	// Parse Template File, Handle Error
	t, err := template.ParseFiles("./templates/homepage.html")
	if err != nil {
		log.Print("template parsing error: ", err)
	}
	err = t.Execute(w, DateTime)
	if err != nil {
		log.Print("template executing error: ", err)
	}
}

// ScaleSelect :: Update HTML With Scale
func ScaleSelect(w http.ResponseWriter, r *http.Request) {
	// Initialize Random "Tables"
	scales := [12]string{"C", "C# / Db", "D", "D# / Eb", "E", "F", "F# / Gb", "G", "G# / Ab", "A", "A# / Bb", "B"}
	types := [4]string{"Major", "Minor", "Melodic Minor", "Harmonic Minor"}
	directions := [2]string{"Ascending", "Descending"}
	intervals := [7]string{"Steps", "3rds", "4ths", "5ths", "6ths", "7ths", "Octaves"}

	// Get Random Element
	randScale := RandNum(len(scales))
	randTypes := RandNum(len(types))
	randDirections := RandNum(len(directions))
	randIntervals := RandNum(len(intervals))

	// Init Time Value
	now := time.Now()

	// Instantiate Struct ScaleData
	ScaleData := ScaleVariables{
		Date:      now.Format("02-01-2006"),
		Time:      now.Format("15:04:05"),
		Scale:     scales[randScale],
		Type:      types[randTypes],
		Direction: directions[randDirections],
		Interval:  intervals[randIntervals],
		GenScale:  false,
		GenInt:    false,
		Clear:     true,
	}

	// Parse Template File, Handle Error
	t, err := template.ParseFiles("./templates/scale.html")
	if err != nil {
		log.Print("template parsing error: ", err)
	}
	err = t.Execute(w, ScaleData)
	if err != nil {
		log.Print("template executing error: ", err)
	}
}

// RandNum :: Generate Random Number Based On Array Length
func RandNum(lenNum int) (randInt int) {
	randInt = rand.Intn(lenNum)
	return randInt
}

// Serve on 8080
func main() {
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("templates"))))
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/scale", ScaleSelect)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
