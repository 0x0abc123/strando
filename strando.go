package main

import (
	"fmt"
	"flag"
	"strando/worddata"
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
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

func generateRandomBytes(length int) ([]byte, error) {
	randomBytes := make([]byte, length)

	_, err := rand.Read(randomBytes)

	return randomBytes, err
}

func generateGuid() (string, error) {

	randomBytes, err := generateRandomBytes(16)

	// Set the first two bits of the 5th byte to 0b10
	if len(randomBytes) >= 5 {
		randomBytes[6] &= 0b00001111 // Clear the first four bits
		randomBytes[6] |= 0b01000000 // Set the second bit to 1 to indicate UUID version 4
		randomBytes[8] &= 0b00111111 // Clear the first two bits
		randomBytes[8] |= 0b10000000 // Set the first bit to 1
	}

	hexString := hex.EncodeToString(randomBytes)

	// 8-4-4-4-12 format
	guidString := hexString[0:8] + "-" + hexString[8:12] + "-" + hexString[12:16] + "-" + hexString[16:20] + "-" + hexString[20:32]

	return guidString, err
}

func generateRandomInt() (int, error) {

	randomBytes, err := generateRandomBytes(8)
	randomInt := int(binary.BigEndian.Uint64(randomBytes))
	return randomInt, err
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

	if complexity == 4 {
		randomString, _ = generateGuid()

	} else if complexity > 0 {
	
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