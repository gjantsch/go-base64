package main

import (
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

const paddingChar = '='

var CODES = []byte{
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J',
	'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T',
	'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd',
	'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n',
	'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x',
	'y', 'z', '0', '1', '2', '3', '4', '5', '6', '7',
	'8', '9', '+', '/'}

func Base64Encode(input []byte) string {
	var i int
	var enc [4]byte
	var sb strings.Builder
	var inputLength int = len(input)

	sb.Grow(((inputLength + 2) / 3) * 4)

	for i = 0; i < inputLength; i += 3 {
		enc[0] = (input[i] & 0b11111100) >> 2
		enc[1] = (input[i] & 0b00000011) << 4

		sb.WriteByte(CODES[enc[0]])

		if inputLength > i+1 {
			enc[1] = enc[1] | ((input[i+1] & 0b11110000) >> 4)
			enc[2] = (input[i+1] & 0b00001111) << 2

			sb.WriteByte(CODES[enc[1]])

			if inputLength > i+2 {
				enc[2] = enc[2] | ((input[i+2] & 0b11000000) >> 6)
				enc[3] = input[i+2] & 0b00111111
				sb.WriteByte(CODES[enc[2]])
				sb.WriteByte(CODES[enc[3]])
			} else {
				sb.WriteByte(CODES[enc[2]])
				sb.WriteByte(paddingChar)
			}
		} else {
			sb.WriteByte(CODES[enc[1]])
			sb.WriteByte(paddingChar)
			sb.WriteByte(paddingChar)
		}
	}

	return sb.String()
}

func getIndex(input byte) byte {

	if input == '/' {
		return 63
	}

	if input == '+' {
		return 62
	}

	if input >= '0' && input <= '9' {
		return input - '0' + 52
	}
	if input >= 'A' && input <= 'Z' {
		return input - 'A'
	}
	if input >= 'a' && input <= 'z' {
		return input - 'a' + 26
	}

	return 0
}

func Base64Decode(input []byte) string {

	var i int
	var sb strings.Builder
	var a, b, c, d byte
	var inputLength = len(input)

	sb.Grow((inputLength / 4) * 3)

	for i = 0; i < inputLength; i += 4 {
		a = getIndex(input[i])
		b = 0
		c = 0
		d = 0

		if inputLength > i+1 {
			b = getIndex(input[i+1])
		}

		if inputLength > i+2 {
			c = getIndex(input[i+2])
		}

		if inputLength > i+3 {
			d = getIndex(input[i+3])
		}

		a = (a << 2) | (b >> 4)
		b = (b << 4) | (c >> 2)
		c = (c << 6) | d
		sb.WriteByte(a)

		if input[i+2] != '=' {
			sb.WriteByte(b)
		}
		if input[i+3] != '=' {
			sb.WriteByte(c)
		}
	}

	return sb.String()
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
		switch opt {
		case "-d":
			if len(args) < 3 {
				fmt.Println("Invalid arguments.")
				printHelpMessage()
				return
			}
			opts = "decode"
			args = slices.Delete(args, i, i+1)
		case "-h", "--help":
			printHelpMessage()
			return
		}
	}

	if args[1] == "--" {
		input, _ = io.ReadAll(os.Stdin)
	} else if data, err := os.ReadFile(args[1]); err == nil {
		input = data
	} else {
		cleaned := strings.Join(strings.Fields(args[1]), "")
		cleaned = strings.TrimRight(cleaned, "\r\n")
		input = []byte(cleaned)
	}

	if opts == "decode" {
		fmt.Print(Base64Decode(input))
	} else {
		fmt.Print(Base64Encode(input))
	}

	fmt.Println()
}
