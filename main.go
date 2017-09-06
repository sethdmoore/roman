package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var Order = []int{
	1000,
	900,
	500,
	400,
	100,
	90,
	50,
	40,
	10,
	9,
	5,
	4,
	1,
}

var Roman = map[int]string{
	1000: "M",
	900:  "CM",
	500:  "D",
	400:  "CD",
	100:  "C",
	90:   "XC",
	50:   "L",
	40:   "XL",
	10:   "X",
	9:    "IX",
	5:    "V",
	4:    "IV",
	1:    "I",
}

func printHelp() {
	fmt.Printf("Usage: \n")
	fmt.Printf("  roman ARG\n")
	fmt.Printf("    $ roman 32\n")
	fmt.Printf("      XXXII\n")
	fmt.Printf("    $ echo 22 | roman\n")
	fmt.Printf("      XXII\n")
	os.Exit(0)
}

func intToRoman(n int) (output string) {
	for _, i := range Order {
		for n >= i {
			n -= i
			output += Roman[i]
		}
	}
	return output
}

func toInt(s string) (int, error) {
	s = strings.TrimSpace(s)
	inp, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return inp, err
}

func parseUserInput() (int, error) {
	var inp int
	var err error
	if len(os.Args) == 2 {
		// If we have args, use them
		inp, err = toInt(os.Args[1])
		if err != nil {
			return 0, err
		}
	} else if len(os.Args) == 1 {
		// read data from stdin
		var stat os.FileInfo

		stat, err = os.Stdin.Stat()
		if err != nil {
			return 0, err
		}

		// ensure stdin is from an application
		if (stat.Mode() & os.ModeCharDevice) == 0 {
			var b []byte
			b, err = ioutil.ReadAll(os.Stdin)
			if err != nil {
				return 0, err
			}
			s := string(b)
			inp, err = toInt(s)
			if err != nil {
				return 0, err
			}
		} else {
			return 0, errors.New("No valid option")
		}
	} else {
		return 0, err
	}
	// err should be nil
	return inp, err
}

func main() {
	inp, err := parseUserInput()
	if err != nil || inp <= 0 {
		printHelp()
	}

	fmt.Println(intToRoman(inp))
}
