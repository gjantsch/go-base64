package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"
)

var CODES = []byte{
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J',
	'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T',
	'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd',
	'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n',
	'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x',
	'y', 'z', '0', '1', '2', '3', '4', '5', '6', '7',
	'8', '9', '+', '/', '='}

func Base64Encode(input []byte) string {
	var i int
	var enc [4]byte
	var encoded string
	var input_length int = len(input)

	for i = 0; i < input_length; i += 3 {
		enc[0] = (input[i] & 0b11111100) >> 2
		enc[1] = (input[i] & 0b00000011) << 4

		// by default last two chars point to filler '='
		enc[2] = 0b01000000
		enc[3] = 0b01000000

		if input_length > i+1 {
			enc[1] = enc[1] | ((input[i+1] & 0b11110000) >> 4)
			enc[2] = (input[i+1] & 0b00001111) << 2

			if input_length > i+2 {
				enc[2] = enc[2] | ((input[i+2] & 0b11000000) >> 6)
				enc[3] = input[i+2] & 0b00111111
			}
		}

		encoded += string([]byte{CODES[enc[0]], CODES[enc[1]], CODES[enc[2]], CODES[enc[3]]})
	}

	return encoded
}

func getIndex(input byte) byte {

	if input == '/' {
		return 63
	}

	if input == '+' {
		return 62
	}

	// 0-9
	if input >= 48 && input <= 57 {
		return input + 4
	}

	// A-Z
	if input >= 65 && input <= 90 {
		return input - 65
	}

	// a-z
	if input >= 97 && input <= 122 {
		return input - 71
	}

	return 0
}

func Base64Decode(input []byte) string {

	var i int
	var decoded string
	var a, b, c, d byte
	var input_length = len(input)

	for i = 0; i < input_length; i += 4 {
		a = getIndex(input[i])
		b = 0
		c = 0
		d = 0

		if input_length > i+1 {
			b = getIndex(input[i+1])
		}

		if input_length > i+2 {
			c = getIndex(input[i+2])
		}

		if input_length > i+3 {
			d = getIndex(input[i+3])
		}

		a = (a << 2) | (b >> 4)
		b = (b << 4) | (c >> 2)
		c = (c << 6) | d
		decoded += string([]byte{a, b, c})
	}

	return decoded
}

func printHelpMessage() {
	fmt.Println("BASE64 Encoder")
	fmt.Println("Usage:")
	fmt.Println(" base64enc [-d|-h] [filename|--]")
	fmt.Println(" -d       : decode")
	fmt.Println(" -h       : print this help message")
	fmt.Println(" --       : read content from stdin")
	fmt.Println(" filename : literal to be encoded or decoded")
}

func main() {

	args := os.Args
	if len(args) < 2 {
		printHelpMessage()
		return
	}

	var input = []byte{}

	var opts string = "encode"
	for i, opt := range args {
		if opt == "-d" {
			if len(args) < 3 {
				fmt.Println("Invalid arguments.")
				printHelpMessage()
				return
			}
			opts = "decode"
			args = slices.Delete(args, i, i+1)

		} else if opt == "-h" || opt == "--help" {
			printHelpMessage()
			return
		}
	}

	if args[1] == "--" {
		scanner := bufio.NewReader(os.Stdin)
		input, _ = scanner.ReadBytes(0x0)
	} else {
		re := regexp.MustCompile(`\s+`)
		var cleaned string = re.ReplaceAllString(args[1], "")
		cleaned = strings.TrimRight(cleaned, string(byte(0b00001010)))
		input = []byte(cleaned)
	}

	if opts == "decode" {
		fmt.Print(Base64Decode(input))
	} else {
		fmt.Print(Base64Encode(input))
	}

	fmt.Println()
}
