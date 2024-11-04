package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Seed random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate a random number between 1 and 100
	target := rand.Intn(100) + 1

	// Create a reader for input
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Tebak angka dari 1 hingga 100")

	for {
		fmt.Print("Masukkan tebakan Anda: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Masukkan angka yang valid.")
			continue
		}

		if guess < target {
			fmt.Println("Terlalu rendah!")
		} else if guess > target {
			fmt.Println("Terlalu tinggi!")
		} else {
			fmt.Println("Selamat! Anda menebak dengan benar.")
			break
		}
	}
}
