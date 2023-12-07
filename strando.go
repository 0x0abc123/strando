package main

import (
	"fmt"
	"flag"
	"strando/worddata"
	"crypto/rand"
	"encoding/binary"
)

const defaultLength = 32
const defaultComplexity = 3

// Character sets corresponding to the complexity level
const charset_1 = "abcdefghijklmnopqrstuvwxyz0123456789"
const charset_2 = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const charset_3 = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-=_+:;{}[]|,./?><~"

var nounStrings = worddata.NounStrings
var adjStrings = worddata.AdjStrings


func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func generateRandomInt() (int, error) {

	randomBytes := make([]byte, 8)

	_, err := rand.Read(randomBytes)
	if err != nil {
		return 0, err
	}

	randomInt := int(binary.BigEndian.Uint64(randomBytes))
	return randomInt, nil
}

func generateRandomString(length int, charset string) string {

	// Generate the random string
	result := make([]byte, length)
	for i := range result {
		randInt, _ := generateRandomInt()
		result[i] = charset[abs(randInt % len(charset))]
	}

	return string(result)
}

func generateRandomPhrase() string {

	// Generate the random phrase
	randInt, _ := generateRandomInt()
	randAdjective := adjStrings[abs(randInt % len(adjStrings))]

	randInt, _ = generateRandomInt()
	randNoun := nounStrings[abs(randInt % len(nounStrings))]

	return randAdjective + "-" + randNoun
}


func main() {
	// Parse command-line options
	var length int
	flag.IntVar(&length, "l", defaultLength, "Length of the random string")

	var complexity int
	flag.IntVar(&complexity, "c", defaultComplexity, "Complexity of the random string")

	flag.Parse()

	randomString := ""

	if complexity > 0 {
	
		charset := charset_3

		switch complexity {
		case 1:
			charset = charset_1
		case 2:
			charset = charset_2
		default:
			charset = charset_3
		}

		// Generate and print the random string
		randomString = generateRandomString(length, charset)
	} else {
		// Generate and print the random phrase
		randomString = generateRandomPhrase()
	}

	fmt.Printf("%s\n", randomString)
}

// to build: 
// GOOS=linux GOARCH=amd64 go build -o strando