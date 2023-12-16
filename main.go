package main

import (
	"bufio"
	"fmt"
	hash "hashAlgorithm/HashPassword"
	decode "hashAlgorithm/ReadPassword"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome to hash algorithm")
	mainProgress()
}

func mainProgress() {
	fmt.Print("CHOOSE WHAT TO DO options: hash or decode: ")
	reader := bufio.NewReader(os.Stdin)
	choose, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occurred while reading input. Please try again", err)
		return
	}

	choose = strings.TrimSpace(choose)
	switch strings.ToLower(choose) {
	case "hash":
		hashProgress()
	case "decode":
		decodeProgress()
	default:
		fmt.Println("Invalid option")
		mainProgress()
	}

}

func hashProgress() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter password to hash: ")
	password, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occurred while reading input. Please try again", err)
		return
	}
	password = strings.TrimSpace(password)
	fmt.Println(hash.HashPassword(password))
}

func decodeProgress() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter password to decode: ")
	password, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occurred while reading input. Please try again", err)
		return
	}
	password = strings.TrimSpace(password)
	fmt.Println(decode.DecodePassword(password))
}
