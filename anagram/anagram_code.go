// detective_code.go

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/mikegleen/murdle-lib"

	"github.com/etnz/permute"
)

/*
	See read_cipher.go for a description of the input data.
*/

const DATAFILE = "/Users/mlg/goprj/murdle_workspace/data.txt"
const DICTFILE = "/Users/mlg/pyprj/caesar/data/dictionary.txt"
const VERBOS = 1
// func nextline(scanner bufio.Scanner) (string, error) {

// }

func main() {

	if len(os.Args) < 2 {
		panic("\nOne parameter required, the cipher number.")
	}
	totstart := time.Now()

	// Translate table for the Murdle code "A" where the cipher alphabet is in reverse order.
	lokup := make([]rune, 26)
	for r := 'A'; r <= 'Z'; r++ {
		ix := 'Z' - r
		lokup[ix] = r
	}
	// fmt.Println(lokup)

	/*
		Load the dictionary
	*/
	wordDict := make(map[string]struct{})
	datafile, err := os.Open(DICTFILE)
	if err != nil {
		panic(fmt.Sprint("Cannot open ", DATAFILE))
	}
	scanner := bufio.NewScanner(datafile)
	for scanner.Scan() {
		line := scanner.Text()
		err = scanner.Err()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error reading dictionary.")
			log.Fatal(err)
		}
		line = strings.TrimSpace(line)
		if len(line) == 0 || line[0] == '#' {
			continue
		}
		wordDict[line] = struct{}{}
	}
	datafile.Close()
	fmt.Printf("Dictionary contains %d words.\n", len(wordDict))

	c, _ := strconv.Atoi(os.Args[1]) // get the cipher number
	ciphertext, err := lib.ReadCipher(DATAFILE, c)
	if err != nil {
		panic(err)
	}

	reg, _ := regexp.Compile("[^A-Z]+") // remove everything except letters
	words := strings.Fields(ciphertext)
	fmt.Println("\nwords: ", words)
	for _, word := range words {
		rword := reg.ReplaceAllString(word, "")
		fmt.Println("word: ", rword)
		guesses := make(map[string]struct{})
		w := []rune(rword)
		start := time.Now()
		for _, try := range permute.Permutations(w) {
			stry := string(try)
			// If there are repeated letters, the same guess will be printed multiple times.
			_, found := guesses[stry]
			if found {
				continue
			}
			_, found = wordDict[stry]
			if found {
				guesses[stry] = struct{}{}
				fmt.Printf("%16s %s\n", "", stry)
			}
		}
		end := time.Now()
		if VERBOS > 1 {
			fmt.Printf("Calculation finished in %s \n", end.Sub(start))
		}
		if len(guesses) == 0 {
			fmt.Printf("%16s %s\n", "", "?")
		}
	}
	totend := time.Now()
	duration := float64(totend.Sub(totstart)) // nanoseconds
	duration = duration / 1000000000.
	fmt.Printf("Program finished in %.3f seconds.\n", duration)
}
