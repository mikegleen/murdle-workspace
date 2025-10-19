// factorial.go

package main

import (

	"fmt"
	"os"
	"strconv"
	"unicode"
	lib "github.com/mikegleen/murdle-lib"
)

func main() {
	p, _ := strconv.Atoi(os.Args[1]) // get the puzzle number
	c := 1
	if len(os.Args) > 2 {
		c, _ = strconv.Atoi(os.Args[2]) // get the cipher number
	}

	ciphertext, ciphertype, err := lib.ReadCipher2(lib.DATAFILE, p, c)
	if err != nil {
		panic(err)
	}
	var tally [26]int
	inta := int('A')
	fmt.Printf("Cipher type: %s\n", ciphertype)
	for _, ch := range ciphertext {
		ch = unicode.ToUpper(ch)
		if ch < 'A' || ch > 'Z' {
			continue
		}
		tally[int(ch) - inta] += 1
	}
	for n := range 26 {
		fmt.Printf("  %c", rune(inta + n))
	}
	fmt.Println()
	for n := range 26 {
		fmt.Printf("%3d", tally[n])
	}
	fmt.Println()
}
