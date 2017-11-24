//James Kinsella - g00282261@gmit.ie
//GoLang chatbot - November 2017

//Sources: https://golang.org/pkg/regexp/syntax/
// https://golang.org/pkg/regexp/
// https://gobyexample.com/regular-expressions
// https://github.com/StefanSchroeder/Golang-Regex-Tutorial
//

package main

import (
	"regexp"
	//"time"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
)

type text struct {
	input string
}

func ElizaResponse(input string) string {
	//guess := input
	//array of responses the eliza program will respond with
	responses := []string{

		"I don't follow?",
		"You'll have to speak up, I'm wearing a towel",
		"...Sorry, what? I wasn't listening",
	}

	iam := regexp.MustCompile("(?i)i(?:'| a|)?m(.*)")
	if iam.MatchString(input) {
		return iam.ReplaceAllString(input, "How do I know you are $1")
	}

	//Parses input for "father"
	father, _ := regexp.MatchString(`(?i)\\bfather\\b`, input) //\b and ?i help parse "grandfather" and "Father"
	mother, _ := regexp.MatchString(`(?i)\\bmother\\b`, input)
	brother, _ := regexp.MatchString(`(?i)\\bbrother\\b`, input)
	sister, _ := regexp.MatchString(`(?i)\\bsister\\b`, input)
	//Catches input, response:
	if father {
		return ("Why don’t you tell me more about your father?")
	}
	if mother {
		return ("Why don’t you tell me more about your mother?")
	}
	if brother {
		return ("Why don’t you tell me more about your brother?")
	}
	if sister {
		return ("Why don’t you tell me more about your sister?")
	}
	//returns the responses to the main function
	return responses[rand.Intn(len(responses))]
}

func reflection(input string) string {

	//Swi
	pronouns := [][]string{
		{`am`, `are`},
		{`I`, `you`},
		{`you`, `I`},
		{`me`, `you`},
		{`your`, `my`},
		{`my`, `your`},
	}

	// Split input into values
	boundaries := regexp.MustCompile(`\b`)

	values := boundaries.Split(input, -1)

	//Loop through the range of values and reflect the pronoun if it finds a match
	for i, token := range values {
		for _, reflection := range pronouns {
			if matched, _ := regexp.MatchString(reflection[0], token); matched {

				values[i] = reflection[1]
				break
			}
		}
	}

	//Join the string of values back together
	answer := strings.Join(values, ``)

	counterResp := []string{
		"Why do ",
		"How do you know that ",
		"I find it fasinating that ",
		"Are you certain that ",
		"I can't say I would ",
	}

	return (counterResp[rand.Intn(len(counterResp))] + answer)
}

func elizaHandler(w http.ResponseWriter, r *http.Request) {

	userInput := r.URL.Query().Get("input")
	elizaResponse := ElizaResponse(userInput)

	fmt.Fprintf(w, elizaResponse)
}

func main() {
	// Webpage Directory
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/chatbot", elizaHandler)
	http.ListenAndServe(":8080", nil)
}
