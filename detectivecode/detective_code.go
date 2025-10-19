// detective_code.go

package main

import (
	"fmt"
	lib "github.com/mikegleen/murdle-lib"
	"os"
	"strconv"
	"unicode"
)

func main() {

	if len(os.Args) < 2 {
		panic("\nOne parameter required, the cipher number.")
	}

	// Translate table for the Murdle code "A" where the cipher alphabet is in reverse order.
	lokup := make([]rune, 26)
	for r := 'A'; r <= 'Z'; r++ {
		ix := 'Z' - r
		lokup[ix] = r
	}
	// fmt.Println(lokup)

	p, _ := strconv.Atoi(os.Args[1]) // get the page number

		c := 1
	if len(os.Args) > 2 {
		c, _ = strconv.Atoi(os.Args[2]) // get the cipher number
	}

	ciphertext, ciphertype, err := lib.ReadCipher2(lib.DATAFILE, p, c)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Cipher type: %s\n", ciphertype)
	ciphertextr := []rune(ciphertext)
	plaintext := make([]rune, len(ciphertextr))
	for i, r := range ciphertextr {
		r = unicode.ToUpper(r)
		if r < 'A' || r > 'Z' {
			plaintext[i] = r
		} else {
			plaintext[i] = lokup[r-'A']
		}
	}
	fmt.Println(string(plaintext))
}
