package main

import (
	crand "crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Person struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func obfuscateEmail(email string) string {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return email
	}

	local := parts[0]
	if len(local) <= 2 {
		return email
	}

	return local[:2] + strings.Repeat("*", len(local)-2) + "@" + parts[1]
}

func obfuscateName(name string) string {
	runes := []rune(name)
	length := len(runes)
	if length <= 2 {
		return name
	}

	return string(runes[0]) + strings.Repeat("*", length-2) + string(runes[length-1])
}

func randomFileName() string {
	bytes := make([]byte, 16)
	if _, err := crand.Read(bytes); err != nil {
		log.Fatal(err)
	}
	return hex.EncodeToString(bytes) + ".txt"
}

func main() {
	// Read the input.json file
	jsonFile, err := ioutil.ReadFile("input.json")
	if err != nil {
		log.Fatal(err)
	}

	// Parse the JSON data
	var people []Person
	err = json.Unmarshal(jsonFile, &people)
	if err != nil {
		log.Fatal(err)
	}

	// Set the random seed
	rand.Seed(time.Now().UnixNano())

	// Randomly select a person
	if len(people) > 0 {
		randomIndex := rand.Intn(len(people))
		randomName := obfuscateName(people[randomIndex].Name)
		randomEmail := obfuscateEmail(people[randomIndex].Email)
		fmt.Println("Randomly selected name and email:", randomName, randomEmail)

		// Save the actual name and email to a file
		fileName := randomFileName()
		file, err := os.Create(fileName)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		_, err = file.WriteString("Name: " + people[randomIndex].Name + "\nEmail: " + people[randomIndex].Email + "\n")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Actual name and email saved to:", fileName)
	} else {
		fmt.Println("No people data found.")
	}
}
