package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func jump_forward(source []uint8, start int, source_length int) int {
	var target = 0
	var cnt = 1
	for tip := start; tip < source_length; tip++ {
		cc := string(source[tip])
		if cc == "[" {
			cnt += 1
		} else if cc == "]" {
			cnt -= 1
		}
		if cnt == 0 {
			target = tip
			break
		}
	}
	return target + 1
}

func jump_backward(source []uint8, start int) int {
	var target = 0
	var cnt = 1
	for tip := start; tip > 0; tip-- {
		cc := string(source[tip])
		if cc == "]" {
			cnt += 1
		} else if cc == "[" {
			cnt -= 1
		}
		if cnt == 0 {
			target = tip
			break
		}
	}
	return target + 1
}

func main() {
	pathPtr := flag.String("path", "", "path to Brainfuck source")
	sizePtr := flag.Int("size", 30000, "size of the data array")

	flag.Parse()

	path := *pathPtr
	if path == "" {
		fmt.Println("You must specify a path to Brainfuck source: -path=...")
		os.Exit(1)
	}

	source, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(os.Stdin)

	var data = make([]int, *sizePtr)
	var source_length = len(source)
	var dp = 0
	for ip := 0; ip < source_length; {
		var c = string(source[ip])
		switch {
		case c == ">":
			dp += 1
			ip += 1
		case c == "<":
			dp -= 1
			ip += 1
		case c == "+":
			data[dp] += 1
			ip += 1
		case c == "-":
			data[dp] -= 1
			ip += 1
		case c == ".":
			fmt.Print(string(data[dp]))
			ip += 1
		case c == ",":
			s, err := reader.ReadString('\n')
			if err != nil {
				panic(err)
			}
			if len(s) != 2 {
				panic("You must provide a single character")
			}
			data[dp] = int(s[0])
			ip += 1
		case c == "[":
			if data[dp] == 0 {
				ip = jump_forward(source, ip+1, source_length)
			} else {
				ip += 1
			}
		case c == "]":
			if data[dp] != 0 {
				ip = jump_backward(source, ip-1)
			} else {
				ip += 1
			}
		default:
			ip += 1
		}
	}
}
