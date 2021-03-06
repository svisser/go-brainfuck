package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	charIncrementPointer = '>'
	charDecrementPointer = '<'
	charIncrementData    = '+'
	charDecrementData    = '-'
	charOutput           = '.'
	charInput            = ','
	charJumpForward      = '['
	charJumpBackward     = ']'
)

func jump_forward(source []uint8, start int, source_length int) int {
	var count = 1
	for tip := start; tip < source_length; tip++ {
		if source[tip] == charJumpForward {
			count += 1
		} else if source[tip] == charJumpBackward {
			count -= 1
		}
		if count == 0 {
			return tip + 1
		}
	}
	return -1
}

func jump_backward(source []uint8, start int) int {
	var count = 1
	for tip := start; tip >= 0; tip-- {
		if source[tip] == charJumpBackward {
			count += 1
		} else if source[tip] == charJumpForward {
			count -= 1
		}
		if count == 0 {
			return tip + 1
		}
	}
	return -1
}

func main() {
	pathPtr := flag.String("path", "", "path to Brainfuck source")
	sizePtr := flag.Int("size", 30000, "size of the data array")

	flag.Parse()

	if *pathPtr == "" {
		fmt.Println("You must specify a path to Brainfuck source: -path=...")
		os.Exit(1)
	}

	source, err := ioutil.ReadFile(*pathPtr)
	if err != nil {
		panic(err)
	}

	var data = make([]int64, *sizePtr)
	var source_length = len(source)
	var dp = 0
	for ip := 0; ip < source_length; {
		var c = source[ip]
		switch {
		case c == charIncrementPointer:
			dp += 1
			ip += 1
		case c == charDecrementPointer:
			dp -= 1
			ip += 1
		case c == charIncrementData:
			data[dp] += 1
			ip += 1
		case c == charDecrementData:
			data[dp] -= 1
			ip += 1
		case c == charOutput:
			fmt.Printf("%c", data[dp])
			ip += 1
		case c == charInput:
			fmt.Scanf("%c", &data[dp])
			ip += 1
		case c == charJumpForward:
			if data[dp] == 0 {
				ip = jump_forward(source, ip+1, source_length)
			} else {
				ip += 1
			}
		case c == charJumpBackward:
			if data[dp] != 0 {
				ip = jump_backward(source, ip-1)
			} else {
				ip += 1
			}
		default:
			ip += 1
		}
		// execution ends as expected when ip == source_length
		if ip < 0 || ip > source_length {
			panic(fmt.Sprintf("Aborting execution as instruction pointer is invalid: %d.", ip))
		}
		if dp < 0 || dp >= *sizePtr {
			panic(fmt.Sprintf("Aborting execution as data pointer is invalid: %d.", dp))
		}
	}
}
