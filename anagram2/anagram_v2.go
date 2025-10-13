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
See lib/read_cipher.go for a description of the data file format.
*/
const DICTFILE = "/Users/mlg/pyprj/caesar/data/dictionary.txt"
const SHORTDICTFILE = "/Users/mlg/pyprj/caesar/data/short_words.txt"

const WORDLIMIT = 51 // max word size is 50
const SHOWTIMES = false
const SHOWCOUNTS = false // number of words in each dictionary slot

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func read_dict(dict []map[string]string, dictfilename string) (int, int) {
	maxlen := 0
	nwords := 0

	dictfile, err := os.Open(dictfilename)
	if err != nil {
		panic(fmt.Sprint("Cannot open ", dictfilename))
	}
	scanner := bufio.NewScanner(dictfile)
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
		dict[lenword][line] = SortString(line)
	}
	dictfile.Close()
	return maxlen, nwords
}

func main() {

	if len(os.Args) < 2 {
		panic("\nOne parameter required, the cipher number.")
	}
	totstart := time.Now()

	/******************************
		Create the empty dictionary
	*******************************/

	wordDict := make([]map[string]string, WORDLIMIT)
	for w := 1; w < WORDLIMIT; w++ {
		wordDict[w] = make(map[string]string)
	}

	/*************************
		Read the dictionary
	**************************/

	m, n := read_dict(wordDict, DICTFILE)
	maxlen := m
	nwords := n
	m, n = read_dict(wordDict, SHORTDICTFILE)
	if m > maxlen {
		maxlen = m
	}
	nwords += n

	fmt.Printf("Dictionary contains %d words.\n", nwords)
	fmt.Printf("Maximum word length: %d.\n", maxlen)

	if SHOWCOUNTS {
		counts := make([]int, WORDLIMIT)
		for i := 1; i < WORDLIMIT; i++ {
			counts[i] = len(wordDict[i])
		}
		for i := 1; i < WORDLIMIT; i++ {
			if counts[i] > 0 {
				fmt.Println(i, counts[i])
			}
		}
	}

	/*************************
		Get the cipher
	**************************/

	p, _ := strconv.Atoi(os.Args[1]) // get the puzzle number
	c := 1
	if len(os.Args) > 2 {
		c, _ = strconv.Atoi(os.Args[2]) // get the cipher number
	}

	ciphertext, err := lib.ReadCipher2(lib.DATAFILE, p, c)
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
		guesses := make(map[string]struct{}) // effectively a "set"
		start := time.Now()
		lenword := len(word)
		if lenword >= WORDLIMIT {
			panic(fmt.Sprintf("\"%v\" in cipher exceeds %d characters. Length = %d\n", word, WORDLIMIT, lenword))
		}
		sword := SortString(word)
		for dword := range wordDict[lenword] {
			// if the sorted cipher word equals the sorted dictionary word, it is an anagram.
			if sword == wordDict[lenword][dword] {
				guesses[dword] = struct{}{}
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
