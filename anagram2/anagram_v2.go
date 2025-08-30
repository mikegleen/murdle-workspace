// anagram_2.go

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	lib "github.com/mikegleen/murdle-lib"
)

/*
The data file contains one row for each cipher.
Blank rows and rows with '#' in column 1 are ignored.

Columns
1-2     Puzzle number in decimal
3-n 	The cipher
*/
const DATAFILE = "/Users/mlg/goprj/murdle_workspace/data.txt"
const DICTFILE = "/Users/mlg/pyprj/caesar/data/dictionary.txt"

const WORDLIMIT = 51  // max word size is 50
const SHOWTIMES = false

// func nextline(scanner bufio.Scanner) (string, error) {

// }
func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func main() {

	if len(os.Args) < 2 {
		panic("\nOne parameter required, the cipher number.")
	}
	totstart := time.Now()

	/*************************
		Load the dictionary
	**************************/

	// wordDict := make(map[string]struct{})
	wordDict := make([]map[string]string, WORDLIMIT)
	for w := 1; w < WORDLIMIT; w++ {
		wordDict[w] = make(map[string]string)
	}
	datafile, err := os.Open(DICTFILE)
	if err != nil {
		panic(fmt.Sprint("Cannot open ", DATAFILE))
	}

	/*************************
		Read the cipher
	**************************/

	maxlen := 0
	nwords := 0
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
		if len(line) > maxlen {
			maxlen = len(line)
		}
		nwords += 1
		lenword := len(line)
		if lenword >= WORDLIMIT {
			panic(fmt.Sprintf("\"%v\" in dictionary exceeds %d characters. Length = %d\n", line, WORDLIMIT, lenword))
		}
		wordDict[lenword][line] = SortString(line)
	}
	datafile.Close()
	fmt.Printf("Dictionary contains %d words.\n", nwords)
	fmt.Printf("Maximum word length: %d.\n", maxlen)
	c, _ := strconv.Atoi(os.Args[1]) // get the cipher number
	ciphertext, err := lib.ReadCipher(DATAFILE, c)
	if err != nil {
		panic(err)
	}

	reg, _ := regexp.Compile("[^A-Z]+") // remove everything except letters
	words := strings.Fields(ciphertext)

	/*****************************
		Search for the anagrams
	******************************/

	for _, word := range words {
		rword := reg.ReplaceAllString(word, "")
		fmt.Println("word: ", rword)
		guesses := make(map[string]struct{})
		start := time.Now()
		lenword := len(word)
		if lenword >= WORDLIMIT {
			panic(fmt.Sprintf("\"%v\" in cipher exceeds %d characters. Length = %d\n", word, WORDLIMIT, lenword))
		}
		sword := SortString(word)
		for dword := range wordDict[lenword] {
			if sword == wordDict[lenword][dword] {
				guesses[dword] = struct {}{}
				fmt.Printf("%16s %s\n", "", dword)
			}
		}

		end := time.Now()
		if SHOWTIMES {
			fmt.Printf("Calculation finished in %s \n", end.Sub(start))
		}
		if len(guesses) == 0 {
			fmt.Printf("%16s %s\n", "", "?")
		}
	}
	totend := time.Now()
	fmt.Printf("Program finished in %s \n", totend.Sub(totstart))
}
